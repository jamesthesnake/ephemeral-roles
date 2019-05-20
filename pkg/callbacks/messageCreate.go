package callbacks

import (
	"os"
	"strings"

	"github.com/ewohltman/ephemeral-roles/pkg/storage"

	"github.com/bwmarrin/discordgo"
	"github.com/ewohltman/ephemeral-roles/pkg/logging"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/sirupsen/logrus"
)

var (
	prometheusMessageCreateCounter = prometheus.NewCounter(
		prometheus.CounterOpts{
			Namespace: "ephemeral_roles",
			Name:      "message_create_events",
			Help:      "Total MessageCreate events",
		},
	)
	infoMessage = &discordgo.MessageEmbed{
		URL:    "https://github.com/ewohltman/ephemeral-roles",
		Title:  "Ephemeral Roles",
		Color:  0xffa500,
		Footer: &discordgo.MessageEmbedFooter{Text: "Made using the discordgo library"},
		Image:  &discordgo.MessageEmbedImage{URL: "https://raw.githubusercontent.com/ewohltman/ephemeral-roles/master/web/static/logo_Testa_anatomica_(1854)_-_Filippo_Balbi.jpg"},
		Fields: []*discordgo.MessageEmbedField{
			{
				Name:   "About",
				Value:  "Ephemeral Roles is a discord bot designed to assign roles based upon voice channel member presence",
				Inline: false,
			},
			{
				Name:   "Author",
				Value:  "Ephemeral Roles is created by ewohltman",
				Inline: false,
			},
			{
				Name:   "Library",
				Value:  "Ephemeral Roles uses the discordgo library by bwmarrin",
				Inline: false,
			},
		},
	}
)

type incomingMessage struct {
	s      *discordgo.Session
	m      *discordgo.MessageCreate
	c      *discordgo.Channel
	g      *discordgo.Guild
	tokens []string
}

// MessageCreate is the callback function for the MessageCreate event from Discord
func MessageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	// Increment the total number of MessageCreate events
	prometheusMessageCreateCounter.Inc()

	// Ignore all messages from bots
	if m.Author.Bot {
		return
	}

	// Check if the message starts with our keyword
	if !strings.HasPrefix(m.Content, BOTKEYWORD) {
		return
	}

	// [BOT_KEYWORD] [command] [options] :: "!eph" "log_level" "debug"
	tokens := strings.Split(strings.TrimSpace(m.Content), " ")
	if len(tokens) < 2 {
		return
	}

	// Find the channel
	c, err := s.State.Channel(m.ChannelID)
	if err != nil {
		log.WithError(err).Debugf("Unable to find channel")
		return
	}

	// Find the guild for that channel
	g, err := s.State.Guild(c.GuildID)
	if err != nil {
		log.WithError(err).Debugf("Unable to find guild")
		return
	}

	parseMessage(
		&incomingMessage{
			s:      s,
			m:      m,
			c:      c,
			g:      g,
			tokens: tokens,
		},
	)
}

func parseMessage(message *incomingMessage) {
	logFields := logrus.Fields{
		"author":  message.m.Author.Username,
		"content": message.m.Content,
		"channel": message.c.Name,
		"guild":   message.g.Name,
		"tokens":  message.tokens,
	}

	log.WithFields(logFields).Debugf("New message")

	switch strings.ToLower(message.tokens[1]) {
	case "info":
		_, err := message.s.ChannelMessageSendEmbed(message.m.ChannelID, infoMessage)
		if err != nil {
			log.WithError(err).Debugf("Unable to send message")
			return
		}
	case "config":
		parseServerConfig(message.g.Name, message.tokens[2:])
	case "log_level":
		if len(message.tokens) >= 3 {
			levelOpt := strings.ToLower(message.tokens[2])

			logFields["log_level"] = levelOpt

			switch levelOpt {
			case "debug":
				updateLogLevel(levelOpt)
				log.WithFields(logFields).Debugf("Logging level changed")
			case "info":
				updateLogLevel(levelOpt)
				log.WithFields(logFields).Infof("Logging level changed")
			case "warn":
				updateLogLevel(levelOpt)
				log.WithFields(logFields).Warnf("Logging level changed")
			case "error":
				updateLogLevel(levelOpt)
				log.WithFields(logFields).Errorf("Logging level changed")
			case "fatal":
				updateLogLevel(levelOpt)
			case "panic":
				updateLogLevel(levelOpt)
			}
		}
	default:
		// Silently fail for unrecognized command
	}
}

func parseServerConfig(name string, options []string) error {
	// TODO: Implement custom server config parsing
	config := &storage.ServerConfig{
		Name: name,
	}

	storageClient.Store(config)

	return nil
}

func updateLogLevel(levelOpt string) {
	err := os.Setenv("LOG_LEVEL", levelOpt)
	if err != nil {
		log.WithError(err).Warn("Unable to set LOG_LEVEL environment variable")
		return
	}

	logging.Reinitialize()
}
