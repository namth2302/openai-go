package openai_test

import (
	"fmt"
	"testing"

	openai "github.com/namth2302/openai-go"
)

// TestListModels Tests the list models endpoint of the API using the mocked server.
func TestListModels(t *testing.T) {
	config := openai.DefaultConfig("whatever")
	client := openai.NewClientWithConfig(config)

	resp, _ := client.ListModels()
	fmt.Println(resp)
}
