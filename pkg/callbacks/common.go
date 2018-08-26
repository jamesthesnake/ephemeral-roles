// Package callbacks is a collection of the callback functions used in response
// to events from Discord's Websocket (WS)API.  Common definitions across the
// package are contained within common.go
package callbacks

import (
	"os"

	"github.com/ewohltman/ephemeral-roles/pkg/logging"
	"github.com/prometheus/client_golang/prometheus"
)

var log = logging.Instance()

// BOTNAME is the name of the bot
var BOTNAME = os.Getenv("BOT_NAME")

// BOTKEYWORD is the keyword message prefix the bot should watch for
var BOTKEYWORD = os.Getenv("BOT_KEYWORD")

// ROLEPREFIX is the prefix to add before ephemeral role names
var ROLEPREFIX = os.Getenv("ROLE_PREFIX") + " "

func init() {
	err := prometheus.Register(prometheusReadyCounter)
	if err != nil {
		log.WithError(err).Error("Unable to register Ready events metric with Prometheus")
	}

	err = prometheus.Register(prometheusVoiceStateUpdateCounter)
	if err != nil {
		log.WithError(err).Error("Unable to register VoiceStateUpdate events metric with Prometheus")
	}

	err = prometheus.Register(prometheusMessageCreateCounter)
	if err != nil {
		log.WithError(err).Error("Unable to register MessageCreate events metric with Prometheus")
	}
}
