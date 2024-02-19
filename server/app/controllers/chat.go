package controllers

import (
	"context"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/m-butterfield/heatmap-creator/server/app/data"
	"github.com/m-butterfield/heatmap-creator/server/app/lib"
	"github.com/sashabaranov/go-openai"
	"io"
	"os"
	"time"
)

func chat(c *gin.Context) {
	prompt := c.Query("p")
	if prompt == "" {
		c.AbortWithStatus(400)
		return
	}
	token := c.Query("t")
	if token == "" {
		c.AbortWithStatus(403)
		return
	}

	accessToken, err := ds.GetAccessTokenByQueryToken(token)
	if err != nil {
		lib.InternalError(err, c)
		return
	}

	yesterday := time.Now().AddDate(0, 0, -1)
	queryCount, err := ds.GetQueryCountForUser(accessToken.UserID, &yesterday)
	if err != nil {
		lib.InternalError(err, c)
		return
	}
	if queryCount >= accessToken.User.DailyQueries {
		c.Stream(func(w io.Writer) bool {
			c.SSEvent("message", "Daily query limit exceeded.")
			return false
		})
		return
	}

	query := data.Query{
		UserID:    accessToken.UserID,
		QueryText: prompt,
	}
	if err = ds.CreateQuery(&query); err != nil {
		lib.InternalError(err, c)
		return
	}

	streamChat(c, prompt, accessToken.User.Username)
}

func streamChat(c *gin.Context, prompt string, username string) {
	gpt := openai.NewClient(os.Getenv("OPENAI_API_KEY"))
	ctx := context.Background()

	req := openai.ChatCompletionRequest{
		Model:       openai.GPT3Dot5Turbo,
		Temperature: 0.3,
		MaxTokens:   1024,
		Messages: []openai.ChatCompletionMessage{
			{
				Content: prompt,
				Role:    "user",
			},
		},
		Stream: true,
		User:   username,
	}

	stream, err := gpt.CreateChatCompletionStream(ctx, req)
	if err != nil {
		lib.InternalError(err, c)
		return
	}
	defer stream.Close()

	chanStream := make(chan string, 10)
	go func() {
		defer close(chanStream)
		for {
			response, err := stream.Recv()
			if errors.Is(err, io.EOF) {
				return
			}

			if err != nil {
				lib.InternalError(err, c)
				return
			}
			if len(response.Choices) > 0 {
				chanStream <- response.Choices[0].Delta.Content
			}
		}
	}()
	c.Stream(func(w io.Writer) bool {
		if msg, ok := <-chanStream; ok {
			c.SSEvent("message", "\""+msg+"\"") // somehow leading whitespace was being stripped from the msg, wrap message in quotes to avoid this and strip them on the FE
			return true
		}
		return false
	})
}
