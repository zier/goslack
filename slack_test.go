package goslack

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func newTestSlackService() *SlackService {
	slack, _ := New(os.Getenv("SLACKURLHOOK"), "Bot", ":smile_cat:", "#test-channel")
	return slack
}

func TestSendMessage(t *testing.T) {
	slack := newTestSlackService()
	err := slack.SetIcon(":ghost:").SetMessage("Hello World").Send()

	assert.NoError(t, err)
}

func TestSetIcon(t *testing.T) {
	slack := newTestSlackService()
	slack.SetIcon(":smile_cate:")

	assert.Equal(t, ":smile_cate:", slack.Icon)
}

func TestSetChannel(t *testing.T) {
	slack := newTestSlackService()
	slack.SetChannel("#football")

	assert.Equal(t, "#football", slack.Channel)
}

func TestUsername(t *testing.T) {
	slack := newTestSlackService()
	slack.SetUsername("admin")

	assert.Equal(t, "admin", slack.Username)
}
