package botHandlers

import (
	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
)

// ReverseMessage reverses the input text and sends it back
func ReverseMessage(b *gotgbot.Bot, ctx *ext.Context) error {
	// Get the message text
	text := ctx.EffectiveMessage.Text

	// Reverse the text
	reversed := reverseString(text)

	// Send the reversed text back
	_, err := ctx.EffectiveMessage.Reply(b, reversed, nil)
	if err != nil {
		return err
	}

	return nil
}

// WelcomeMessage responds to 'hi' with a welcome message
func WelcomeMessage(b *gotgbot.Bot, ctx *ext.Context) error {
	// Send welcome message
	_, err := ctx.EffectiveMessage.Reply(b, "Hello! Welcome to our bot! 👋\nI can reverse any message you send me.", nil)
	if err != nil {
		return err
	}

	return nil
}

// reverseString reverses a string
func reverseString(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}
