package gpt

import (
	"context"
	"fmt"

	"github.com/sashabaranov/go-openai"
)

type GPTService struct {
	client *openai.Client
}

func NewGPTService(apiKey string) *GPTService {
	return &GPTService{
		client: openai.NewClient(apiKey),
	}
}

func (s *GPTService) GenerateAnalysis(prompt string) (string, error) {
	req := openai.CompletionRequest{
		Model:  openai.GPT3Curie,
		Prompt: prompt,
	}

	resp, err := s.client.CreateCompletion(context.Background(), req)
	if err != nil {
		return "", err
	}

	if len(resp.Choices) > 0 {
		return resp.Choices[0].Text, nil
	}

	return "", fmt.Errorf("no response from GPT-3")
}
