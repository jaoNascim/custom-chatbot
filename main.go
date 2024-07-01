package main

import (
	"bufio"
	"context"
	"fmt"
	"os"

	"github.com/sashabaranov/go-openai"
)

func main() {
	c := openai.NewClient(os.Getenv("OPENAI_API_KEY"))

	req := openai.ChatCompletionRequest{
		Model: openai.GPT3Dot5Turbo,
		Messages: []openai.ChatCompletionMessage{
			{
				Role:    openai.ChatMessageRoleSystem,
				Content: "respond as an exaggerated Mojo Jojo",
			},
		},
		MaxTokens: 1024,
	}

	fmt.Print("> ")

	s := bufio.NewScanner(os.Stdin)

	for s.Scan() {
		req.Messages = append(req.Messages, openai.ChatCompletionMessage{
			Role:    openai.ChatMessageRoleUser,
			Content: s.Text(),
		})

		resp, err := c.CreateChatCompletion(context.Background(), req)
		if err != nil {
			panic(err)
		}

		fmt.Printf("%s\n\n", resp.Choices[0].Message.Content)
		req.Messages = append(req.Messages, resp.Choices[0].Message)
		fmt.Print("> ")
	}
}
