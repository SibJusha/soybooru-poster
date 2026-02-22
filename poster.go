package main

import (
	"context"
	"fmt"
	"log"
	"strings"

	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
)

func postToTgChannel(ctx context.Context, b *bot.Bot, tgChannel string, url string, post Post) error {
	var caption strings.Builder
	for tag := range strings.SplitSeq(post.Tags, " ") {
		if !strings.Contains(tag, ":") {
			caption.WriteString("#" + tag + " ")
		}
	}
	caption.WriteString("\n" + fmt.Sprint(post.ID))

	params := bot.SendPhotoParams{
		ChatID:  tgChannel,
		Photo:   &models.InputFileString{Data: url + post.FileURL},
		Caption: caption.String(),
	}

	_, err := b.SendPhoto(ctx, &params)
	if err != nil {
		log.Fatalf("SendPhoto failed: %v", err)
	}

	return err
}
