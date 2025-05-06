package clients

import (
	"context"
	"fmt"
	"log"
	"os"
	"strings"

	openai "github.com/sashabaranov/go-openai"
)

type OpenAIClient struct {
	client *openai.Client
}

func NewOpenAIClient() (*OpenAIClient, error) {
	apiKey := os.Getenv("OPENAI_API_KEY")
	apiKey = strings.TrimSpace(apiKey) // Remove any whitespace

	if apiKey == "" {
		return nil, fmt.Errorf("OPENAI_API_KEY environment variable is required")
	}

	if !strings.HasPrefix(apiKey, "sk-") {
		return nil, fmt.Errorf("invalid API key format: should start with 'sk-'")
	}

	log.Printf("Initializing OpenAI client with API key starting with: %s", apiKey[:6])
	client := openai.NewClient(apiKey)
	return &OpenAIClient{
		client: client,
	}, nil
}

func (c *OpenAIClient) GenerateSciFiStory() (string, error) {
	ctx := context.Background()

	log.Printf("Attempting to generate story with model: gpt-3.5-turbo")
	resp, err := c.client.CreateChatCompletion(
		ctx,
		openai.ChatCompletionRequest{
			Model: "gpt-3.5-turbo",
			Messages: []openai.ChatCompletionMessage{
				{
					Role:    openai.ChatMessageRoleSystem,
					Content: "You are a creative sci-fi story writer. Write very concise but engaging stories.",
				},
				{
					Role:    openai.ChatMessageRoleUser,
					Content: "Generate a creative and engaging sci-fi story in less than 400 characters. Make it concise but impactful.",
				},
			},
			MaxTokens: 150,
		},
	)

	if err != nil {
		log.Printf("❌ OpenAI API error details: %v", err)
		return "", fmt.Errorf("failed to generate story: %w", err)
	}

	log.Printf("✅ Successfully generated story")
	return resp.Choices[0].Message.Content, nil
}
