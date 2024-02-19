package resolvers

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.44

import (
	"context"
	"os"

	"github.com/sashabaranov/go-openai"
)

// Chat is the resolver for the chat field.
func (r *queryResolver) Chat(ctx context.Context, prompt string) (string, error) {
	gpt := openai.NewClient(os.Getenv("OPENAI_API_KEY"))

	compReq := openai.CompletionRequest{
		Model:     openai.GPT3Dot5Turbo,
		MaxTokens: 20,
		Prompt:    prompt,
	}
	resp, err := gpt.CreateCompletion(ctx, compReq)
	if err != nil {
		return "", err
	}

	return resp.Choices[0].Text, nil
}
