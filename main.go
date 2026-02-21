package main

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/go-telegram/bot"
	"github.com/joho/godotenv"
)

const mainURL = "https://soybooru.com/"

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Printf("Warning: couldn't load .env file: %v, loading from env var", err)
	}

	token := os.Getenv("TELEGRAM_BOT_APITOKEN")
	if token == "" {
		log.Fatal("TELEGRAM_BOT_APITOKEN env var is required")
	}

	channel := os.Getenv("TELEGRAM_CHANNEL")
	if channel == "" {
		log.Fatal("TELEGRAM_CHANNEL env var is required (e.g. @mychannel or -100xxxxxxxxxx)")
	}

	bot, err := bot.New(token)
	if err != nil {
		log.Fatal(err)
	}

	ticker := time.NewTicker(15 * time.Minute)
	defer ticker.Stop()

	for {
		state, err := loadState()
		log.Printf("Polling... last_max_id = %d", state.LastMaxID)

		newPosts, newMaxID, err := fetchNewPosts(state.LastMaxID)
		if err != nil {
			log.Printf("Fetch error: %v", err)
			time.Sleep(5 * time.Minute) // backoff
			continue
		}

		if len(newPosts) == 0 {
			log.Println("No new posts")
		} else {
			log.Printf("Found %d new posts", len(newPosts))

			ctx := context.Background()
			successCount := 0

			for _, p := range newPosts {
				if err := postToTgChannel(ctx, bot, channel, mainURL, p); err != nil {
					log.Printf("Post #%d failed: %v", p.ID, err)
				} else {
					successCount++
					log.Printf("Posted #%d", p.ID)
				}
				time.Sleep(10 * time.Second)
			}

			if successCount > 0 {
				state.LastMaxID = newMaxID
				saveState(state)
				log.Printf("Updated state to %d", newMaxID)
			}
		}

		<-ticker.C
	}
}
