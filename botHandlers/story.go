package botHandlers

import (
	"fmt"
	"telegram-bot/clients"

	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
)

func GenerateStory(b *gotgbot.Bot, ctx *ext.Context) error {
	// Send a "thinking" message
	msg, err := ctx.EffectiveMessage.Reply(b, "🤔 Generating a sci-fi story...", nil)
	if err != nil {
		return err
	}

	// Initialize OpenAI client
	openaiClient, err := clients.NewOpenAIClient()
	if err != nil {
		errorMsg := fmt.Sprintf("❌ OpenAI Client Error: %v", err)
		_, _, err = b.EditMessageText(errorMsg, &gotgbot.EditMessageTextOpts{
			ChatId:    ctx.EffectiveChat.Id,
			MessageId: msg.MessageId,
		})
		return err
	}

	// Generate the story
	story, err := openaiClient.GenerateSciFiStory()
	if err != nil {
		errorMsg := fmt.Sprintf("❌ Story Generation Error: %v", err)
		_, _, err = b.EditMessageText(errorMsg, &gotgbot.EditMessageTextOpts{
			ChatId:    ctx.EffectiveChat.Id,
			MessageId: msg.MessageId,
		})
		return err
	}

	// Edit the message with the generated story
	_, _, err = b.EditMessageText("🚀 Here's your sci-fi story:\n\n"+story, &gotgbot.EditMessageTextOpts{
		ChatId:    ctx.EffectiveChat.Id,
		MessageId: msg.MessageId,
	})
	return err
}
