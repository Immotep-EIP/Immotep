package chatgpt

import (
	"bytes"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"os"

	"immotep/backend/prisma/db"
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

func buildImageContent(picture string) map[string]any {
	return map[string]any{
		"type": "image_url",
		"image_url": map[string]any{
			"url": "data:image/jpeg;base64," + picture,
		},
	}
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

func summarize(stuffType string, name string, pictures []string) (string, error) {
	content := []map[string]any{
		{
			"type": "text",
			"text": "You are a real estate agent doing the first inventory report for a new tenant in a property you own. " +
				"This step is about evaluating a " + stuffType + " named " + name + ". " +
				"Along with these instructions are some pictures of the current state of the " + stuffType + ". " +
				"After analysing these pictures, please give me the " + stuffType + "'s state (broken, needsRepair, bad, medium, good, new), the cleanliness (dirty medium clean), and any additional notes describing the maintenance status of what you see. " +
				"Be concise and clear. Only send the information that I asked for without any other text in the format: state|cleanliness|notes.",
		},
	}
	for _, picture := range pictures {
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

func SummarizeRoom(roomName string, pictures []string) (string, error) {
	return summarize("room", roomName, pictures)
}

func SummarizeFurniture(furnitureName string, pictures []string) (string, error) {
	return summarize("furniture", furnitureName, pictures)
}

func compare(stuffType string, initialState string, initialCleanliness string, initialNote string, name string, initialPictures []string, currentPictures []string) (string, error) {
	contentInitial := []map[string]any{
		{
			"type": "text",
			"text": "These pictures were taken during the first inventory report.",
		},
	}
	for _, picture := range initialPictures {
		contentInitial = append(contentInitial, buildImageContent(picture))
	}
	contentCurrent := []map[string]any{
		{
			"type": "text",
			"text": "These pictures were taken during the current (last) inventory report.",
		},
	}
	for _, picture := range currentPictures {
		contentCurrent = append(contentCurrent, buildImageContent(picture))
	}
	contentText := []map[string]any{
		{
			"type": "text",
			"text": "You are a real estate agent doing the last inventory report for a tenant who has lived in a property you own. " +
				"This step is about highlighting the major differences of a " + stuffType + " named " + name + " in comparison with the first inventory report made with the tennat. " +
				"During the first inventory report, this " + stuffType + " has been evaluated with a state: " + initialState + ", a cleanliness: " + initialCleanliness + ", and a note: " + initialNote + ". " +
				"Before these instructions I sent you two sets of pictures corresponding to the initial and the current state of this " + stuffType + ". " +
				"After analysing these pictures, please give me the " + stuffType + "'s differences between these two sets of pictures as the state (broken, needsRepair, bad, medium, good, new), the cleanliness (dirty medium clean), and any additional notes describing the maintenance status of what you see. " +
				"Be concise and clear. Only send the information that I asked for without any other text in the format: state|cleanliness|notes.",
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

func CompareRoom(roomName string, initialReport db.RoomStateModel, initialPictures []string, currentPictures []string) (string, error) {
	return compare("room", string(initialReport.State), string(initialReport.Cleanliness), initialReport.Note, roomName, initialPictures, currentPictures)
}

func CompareFurniture(furnitureName string, initialReport db.FurnitureStateModel, initialPictures []string, currentPictures []string) (string, error) {
	return compare("furniture", string(initialReport.State), string(initialReport.Cleanliness), initialReport.Note, furnitureName, initialPictures, currentPictures)
}
