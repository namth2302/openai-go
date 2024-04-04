package openai_test

import (
	// "context"
	"errors"
	"testing"

	openai "github.com/namth2302/openai-go"
)

func TestCompletionsWrongModel(t *testing.T) {
	config := openai.DefaultConfig("whatever")
	client := openai.NewClientWithConfig(config)
	err := client.CreateCompletion(
		openai.CompletionRequest{
			MaxTokens: 5,
			Model:     openai.GPT3Dot5Turbo,
			Prompt:    "Lorem ipsum",
		},
	)
	if !errors.Is(err, openai.ErrCompletionUnsupportedModel) {
		t.Fatalf("CreateCompletion should return ErrCompletionUnsupportedModel, but returned: %v", err)
	}
}

func TestCompletionWithStream(t *testing.T) {
	config := openai.DefaultConfig("whatever")
	client := openai.NewClientWithConfig(config)
	err := client.CreateCompletion(
		openai.CompletionRequest{Stream: true},
	)
	if !errors.Is(err, openai.ErrCompletionStreamNotSupported) {
		t.Fatal("CreateCompletion didn't return ErrCompletionStreamNotSupported")
	}
}
