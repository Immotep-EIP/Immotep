package chatgpt

import (
	"bytes"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"os"

	"keyz/backend/prisma/db"
	"keyz/backend/utils"
)

type Response struct {
	Choices []Choice `json:"choices"`
}

type Choice struct {
	Message Message `json:"message"`
}

type Message struct {
	Content string `json:"content"`
}

func buildImageContent(pictureUri string) map[string]any {
	return map[string]any{
		"type": "image_url",
		"image_url": map[string]any{
			"url": pictureUri,
		},
	}
}

func buildSummaryPromptMessage(stuffType string, name string) string {
	promptRoom := `
This request is part of an inventory report for a real estate lease. The goal is to document the condition of a room at the beginning of a lease. The provided images show an overview of the room: ` + name + `.

Your task is to analyze these images and generate a structured summary of the room's condition, focusing on key elements such as walls, floor, ceiling, shelves, and any visible signs of damage, wear, or cleanliness.

The response must strictly follow this format:
<state>|<cleanliness>|<notes>
Ex: "good|clean|Example of a note."
DON'T SEND ANYTHING ELSE IN THE RESPONSE.

In case of error, your response must strictly follow this format:
error|<error_message>
Ex: "error|The image does not correspond to the name provided."

Where:
- **State** represents the overall condition of the room and must be one of the following: broken, needsRepair, bad, medium, good, new.
- **Cleanliness** describes the level of cleanliness and must be one of: dirty, medium, clean.
- **Notes** is a short text providing relevant details about any specific observations (e.g., "scuff marks on the wall" or "newly painted").

Please provide an objective and concise assessment based purely on the visual information from the images.
`
	promptFurniture := `
This request is part of an inventory report for a real estate lease. The goal is to document the condition of a specific piece of furniture at the beginning of a lease. The provided images focus on: ` + name + `.

Your task is to analyze these images and generate a structured summary of the furniture's condition. Only assess the furniture itself—ignore the background and any unrelated objects. Identify visible wear, damage, and cleanliness.

The response must strictly follow this format:
<state>|<cleanliness>|<notes>
Ex: "good|clean|Example of a note."
DON'T SEND ANYTHING ELSE IN THE RESPONSE.

In case of error, your response must strictly follow this format:
error|<error_message>
Ex: "error|The image does not correspond to the name provided."

Where:
- **State** represents the overall condition of the furniture and must be one of the following: broken, needsRepair, bad, medium, good, new.
- **Cleanliness** describes the level of cleanliness and must be one of: dirty, medium, clean.
- **Notes** is a short text providing relevant details about any specific observations (e.g., "scratches on the surface" or "like new").

Provide an objective and concise assessment based purely on the visual details in the images.
`
	return utils.Ternary(stuffType == "room", promptRoom, promptFurniture)
}

func buildComparePromptMessage(stuffType string, name string, initialState string, initialCleanliness string, initialNote string) string {
	promptRoom := `
This request is part of a real estate lease inventory comparison process. The goal is to evaluate the changes in the condition of a room: ` + name + ` between the start and end of a lease.

The images for this comparison are sent in two separate messages:
- The first set contains the original images of the room, taken at the start of the lease.
- The second set contains new images of the same room, taken at the end of the lease.

Additionally, here is the room's initial assessment from the first inventory report:
- **Initial State:** ` + initialState + `
- **Initial Cleanliness:** ` + initialCleanliness + `
- **Initial Notes:** ` + initialNote + `

Your task is to analyze and compare the new images with the original ones, focusing on key elements such as walls, floor, ceiling, shelves, and any visible changes in damage, wear, or cleanliness. Consider the initial assessment while determining if the condition has worsened, improved, or remained the same.

The response must strictly follow this format:
<state>|<cleanliness>|<notes>
Ex: "good|clean|Example of a note."
DON'T SEND ANYTHING ELSE IN THE RESPONSE.

In case of error, your response must strictly follow this format:
error|<error_message>
Ex: "error|The image does not correspond to the name provided."

Where:
- **State** represents the overall condition of the room at the end of the lease and must be one of the following: broken, needsRepair, bad, medium, good, new. Compare with the initial state and adjust accordingly if significant deterioration or improvement has occurred.
- **Cleanliness** describes the current level of cleanliness and must be one of: dirty, medium, clean. Compare with the initial cleanliness to determine if the room is cleaner or dirtier.
- **Notes** is an optional short text describing any relevant changes or observations (e.g., "new stains on the floor" or "walls repainted and look new").

Provide an objective and concise assessment based on the visual differences between the two sets of images, incorporating the initial assessment as a reference.
`
	promptFurniture := `
This request is part of a real estate lease inventory comparison process. The goal is to evaluate the changes in the condition of a piece of furniture: ` + name + ` between the start and end of a lease.

The images for this comparison are sent in two separate messages:
- The first set contains the original images of the furniture, taken at the start of the lease.
- The second set contains new images of the same furniture, taken at the end of the lease.

Additionally, here is the furniture's initial assessment from the first inventory report:
- **Initial State:** ` + initialState + `
- **Initial Cleanliness:** ` + initialCleanliness + `
- **Initial Notes:** ` + initialNote + `

Your task is to analyze and compare the new images with the original ones, focusing only on the furniture itself. Ignore the background and any unrelated objects. Identify any visible differences in condition, wear, damage, or cleanliness. Consider the initial assessment while determining if the furniture’s condition has worsened, improved, or remained the same.

The response must strictly follow this format:
<state>|<cleanliness>|<notes>
Ex: "good|clean|Example of a note."
DON'T SEND ANYTHING ELSE IN THE RESPONSE.

In case of error, your response must strictly follow this format:
error|<error_message>
Ex: "error|The image does not correspond to the name provided."

Where:
- **State** represents the overall condition of the furniture at the end of the lease and must be one of the following: broken, needsRepair, bad, medium, good, new. Compare with the initial state and adjust accordingly if significant deterioration or improvement has occurred.
- **Cleanliness** describes the current level of cleanliness and must be one of: dirty, medium, clean. Compare with the initial cleanliness to determine if the furniture is cleaner or dirtier.
- **Notes** is an optional short text describing any relevant changes or observations (e.g., "new scratch on the surface" or "polished and looks better than before").

Provide an objective and concise assessment based on the visual differences between the two sets of images, incorporating the initial assessment as a reference.
`
	return utils.Ternary(stuffType == "room", promptRoom, promptFurniture)
}

func callChatGPT(messages []map[string]any) (string, error) {
	apiURL := "https://api.openai.com/v1/chat/completions"
	apiKey := os.Getenv("OPENAI_API_KEY")

	payload := map[string]any{
		"model":      "gpt-4o",
		"messages":   messages,
		"max_tokens": 300,
	}

	payloadBytes, err := json.Marshal(payload)
	if err != nil {
		panic(err)
	}

	req, err := http.NewRequest(http.MethodPost, apiURL, bytes.NewBuffer(payloadBytes))
	if err != nil {
		panic(err)
	}
	req.Header.Set("Authorization", "Bearer "+apiKey)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	return string(respBody), nil
}

func summarize(stuffType string, name string, picturesUri []string) (string, error) {
	content := []map[string]any{
		{
			"type": "text",
			"text": buildSummaryPromptMessage(stuffType, name),
		},
	}
	for _, picture := range picturesUri {
		content = append(content, buildImageContent(picture))
	}
	messages := []map[string]any{
		{
			"role":    "user",
			"content": content,
		},
	}

	resp, err := callChatGPT(messages)
	if err != nil {
		return "", err
	}
	log.Println(resp)

	var result Response
	if err := json.Unmarshal([]byte(resp), &result); err != nil {
		return "", err
	}

	if len(result.Choices) > 0 {
		return result.Choices[0].Message.Content, nil
	}

	return "", nil
}

func SummarizeRoom(roomName string, picturesUri []string) (string, error) {
	return summarize("room", roomName, picturesUri)
}

func SummarizeFurniture(furnitureName string, picturesUri []string) (string, error) {
	return summarize("furniture", furnitureName, picturesUri)
}

func compare(stuffType string, initialState string, initialCleanliness string, initialNote string, name string, initialPicturesUri []string, currentPicturesUri []string) (string, error) {
	contentInitial := []map[string]any{
		{
			"type": "text",
			"text": "These pictures were taken during the first inventory report.",
		},
	}
	for _, picture := range initialPicturesUri {
		contentInitial = append(contentInitial, buildImageContent(picture))
	}
	contentCurrent := []map[string]any{
		{
			"type": "text",
			"text": "These pictures were taken during the current (last) inventory report.",
		},
	}
	for _, picture := range currentPicturesUri {
		contentCurrent = append(contentCurrent, buildImageContent(picture))
	}
	contentText := []map[string]any{
		{
			"type": "text",
			"text": buildComparePromptMessage(stuffType, name, initialState, initialCleanliness, initialNote),
		},
	}
	messages := []map[string]any{
		{
			"role":    "user",
			"content": contentInitial,
		},
		{
			"role":    "user",
			"content": contentCurrent,
		},
		{
			"role":    "user",
			"content": contentText,
		},
	}

	resp, err := callChatGPT(messages)
	if err != nil {
		return "", err
	}
	log.Println(resp)

	var result Response
	if err := json.Unmarshal([]byte(resp), &result); err != nil {
		return "", err
	}

	if len(result.Choices) > 0 {
		return result.Choices[0].Message.Content, nil
	}
	return "", nil
}

func CompareRoom(roomName string, initialReport db.RoomStateModel, initialPicturesUri []string, currentPicturesUri []string) (string, error) {
	return compare("room", string(initialReport.State), string(initialReport.Cleanliness), initialReport.Note, roomName, initialPicturesUri, currentPicturesUri)
}

func CompareFurniture(furnitureName string, initialReport db.FurnitureStateModel, initialPicturesUri []string, currentPicturesUri []string) (string, error) {
	return compare("furniture", string(initialReport.State), string(initialReport.Cleanliness), initialReport.Note, furnitureName, initialPicturesUri, currentPicturesUri)
}
