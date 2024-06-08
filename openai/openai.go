package openai

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

func MakeOpenAICall(userMessage string) string {

	openAIData := TOpenAIData{
		Model: "gpt-4o",
		Messages: []TOpenAIMessage{
			{Role: "system", Content: "You are a helpful, general purpose chatbot. Be as concise as possible."},
			{Role: "user", Content: userMessage},
		},
	}

	jsonData, err := json.Marshal(openAIData)
	if err != nil {
		log.Fatalln("Failed to marshal json:", err)
	}

	newRequest, err := http.NewRequest(http.MethodPost, "https://api.openai.com/v1/chat/completions", bytes.NewBuffer(jsonData))
	if err != nil {
		log.Fatalf("Failed to create request: %v", err)
	}

	OPENAI_API_KEY := os.Getenv("OPENAI_API_KEY")

	newRequest.Header.Set("Content-Type", "application/json")
	newRequest.Header.Set("Authorization", fmt.Sprintf("Bearer %s", OPENAI_API_KEY))

	client := &http.Client{}
	response, err := client.Do(newRequest)

	if err != nil {
		log.Fatalf("Failed to make the request: %v", err)
	}

	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		log.Fatalf("Failed to read the response body: %v", err)
	}

	var openAIResponse TOpenAIResponse
	err = json.Unmarshal(body, &openAIResponse)

	if err != nil {
		log.Fatal(err)
	}

	return openAIResponse.Choices[0].Message.Content
}
