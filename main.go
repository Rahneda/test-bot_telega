package main

import (
	"log"
	"os"
	"strings"

	"telegram-bot/botHandlers"

	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
	"github.com/PaulSonOfLars/gotgbot/v2/ext/handlers"
	"github.com/PaulSonOfLars/gotgbot/v2/ext/handlers/filters/message"
	"github.com/joho/godotenv"
)

func main() {
	// Load .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Get token from environment variable
	token := os.Getenv("TELEGRAM_TOKEN")
	if token == "" {
		log.Fatal("TELEGRAM_TOKEN environment variable is required")
	}

	// Create bot
	bot, err := gotgbot.NewBot(token, nil)
	if err != nil {
		log.Fatal("Failed to create new bot: ", err)
	}

	// Create dispatcher
	dispatcher := ext.NewDispatcher(&ext.DispatcherOpts{
		Error: func(b *gotgbot.Bot, ctx *ext.Context, err error) ext.DispatcherAction {
			log.Println("an error occurred while handling update:", err.Error())
			return ext.DispatcherActionNoop
		},
		MaxRoutines: ext.DefaultMaxRoutines,
	})

	// Create updater
	updater := ext.NewUpdater(dispatcher, nil)

	// Add handler for all messages
	dispatcher.AddHandler(handlers.NewMessage(message.Text, func(b *gotgbot.Bot, ctx *ext.Context) error {
		text := strings.ToLower(ctx.EffectiveMessage.Text)
		if text == "hi" {
			return botHandlers.WelcomeMessage(b, ctx)
		}
		// If not "hi", reverse the message
		return botHandlers.ReverseMessage(b, ctx)
	}))

	// Start receiving updates
	err = updater.StartPolling(bot, &ext.PollingOpts{
		DropPendingUpdates: true,
		GetUpdatesOpts: &gotgbot.GetUpdatesOpts{
			Timeout: 9,
		},
	})
	if err != nil {
		log.Fatal("Failed to start polling: ", err)
	}
	log.Printf("Bot @%s started successfully!", bot.User.Username)

	// Keep the program running
	updater.Idle()
}
