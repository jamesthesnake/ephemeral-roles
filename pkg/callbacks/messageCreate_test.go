package callbacks

import (
	"testing"

	"github.com/bwmarrin/discordgo"
)

func TestMessageCreate(t *testing.T) {
	_, err := dgTestBotSession.ChannelMessageSendComplex(
		devTextChannelID,
		&discordgo.MessageSend{
			Content: "AUTOMATED TESTING",
		},
	)
	if err != nil {
		t.Error(err)
	}

	// message from a bot
	sendBotMessage()

	// non keyphrase message
	sendMessage("ixnay")

	// keyphrase message, unrecognized command
	sendMessage(BOTKEYWORD + "ixnay")

	// keyphrase message, unrecognized command
	sendMessage(BOTKEYWORD + "AUTOMATED TEST")

	// keyphrase info
	sendMessage(BOTKEYWORD + "info")

	// log_level debug
	sendMessage(BOTKEYWORD + "log_level debug")

	// log_level info
	sendMessage(BOTKEYWORD + "log_level info")

	// log_level warn
	sendMessage(BOTKEYWORD + "log_level warn")

	// log_level error
	sendMessage(BOTKEYWORD + "log_level error")

	// log_level fatal
	sendMessage(BOTKEYWORD + "log_level fatal")

	// log_level panic
	sendMessage(BOTKEYWORD + "log_level panic")

	// log_level info
	sendMessage(BOTKEYWORD + "log_level info")
}

func sendBotMessage() {
	botMsg := &discordgo.MessageCreate{
		Message: &discordgo.Message{
			Author: &discordgo.User{
				Username: "AUTOMATED TEST BOT USER",
				Bot:      true,
			},
			GuildID:   devGuildID,
			ChannelID: devTextChannelID,
			Content:   "AUTOMATED TEST BOT USER",
		},
	}

	MessageCreate(dgTestBotSession, botMsg)
}

func sendMessage(message string) {
	msg := &discordgo.MessageCreate{
		Message: &discordgo.Message{
			Author: &discordgo.User{
				Username: "AUTOMATED TEST USER",
				Bot:      false,
			},
			GuildID:   devGuildID,
			ChannelID: devTextChannelID,
			Content:   message,
		},
	}

	MessageCreate(dgTestBotSession, msg)
}
