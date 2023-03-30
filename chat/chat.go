package chat

import (
	"context"
	"github.com/sashabaranov/go-openai"
)

type Chat interface {
	SendMessageAndGetAnswer(message string) (string, error)
}

type chatData struct {
	client         *openai.Client
	model          string
	messageHistory []openai.ChatCompletionMessage
}

func (c *chatData) SendMessageAndGetAnswer(message string) (string, error) {
	userMessage := openai.ChatCompletionMessage{
		Role:    openai.ChatMessageRoleUser,
		Content: message,
	}

	c.messageHistory = append(c.messageHistory, userMessage)
	resp, err := c.client.CreateChatCompletion(
		context.Background(),
		openai.ChatCompletionRequest{
			Model:    c.model,
			Messages: c.messageHistory,
		},
	)

	if err != nil {
		return "", err
	}

	responseMessageText := resp.Choices[0].Message.Content
	responseMessage := openai.ChatCompletionMessage{
		Role:    openai.ChatMessageRoleAssistant,
		Content: responseMessageText,
	}
	c.messageHistory = append(c.messageHistory, responseMessage)

	return responseMessageText, nil
}

func NewChat(model string, apiKey string) Chat {
	client := openai.NewClient(apiKey)
	return &chatData{
		model:  model,
		client: client,
	}
}
