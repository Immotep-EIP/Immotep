package utils

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-chi/oauth"
)

func WriteJSON(w http.ResponseWriter, status int, data interface{}) error {
	out, err := json.MarshalIndent(data, "", "\t")
	if err != nil {
		fmt.Println("Error Creating JSON")
		return err
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_, err = w.Write(out)
	if err != nil {
		fmt.Println("Error Writing to response")
		return err
	}
	return nil
}

func WriteError(w http.ResponseWriter, status int, text string, err error) {
	if err != nil {
		fmt.Println(text + ":", err)
	} else {
		fmt.Println(text)
	}
	WriteJSON(w, status, map[string]string{"error": text})
}

func GetClaims(r *http.Request) map[string]string {
	claims, ok := r.Context().Value(oauth.ClaimsContext).(map[string]string)
	if !ok {
		return nil
	}
	return claims
}
