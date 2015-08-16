package goslack

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"
)

// SlackService struct
type SlackService struct {
	Token      string
	WebhookURL string
	HTTPClient *http.Client
}

// New SlackService
func New(webhookURL, token string, httpClient *http.Client) (*SlackService, error) {
	if httpClient == nil {
		httpClient = &http.Client{}
	}

	return &SlackService{
		WebhookURL: webhookURL,
		Token:      token,
		HTTPClient: httpClient,
	}, nil
}

// ParseTime parse timestamp from slack
func (slack *SlackService) ParseTime(timeStamp string) time.Time {
	t := strings.Split(timeStamp, ".")
	timestamp, _ := strconv.ParseInt(t[0], 10, 64)

	return time.Unix(timestamp, 0)
}

func (slack *SlackService) buildTextJSON(message, channel string) string {
	if channel != "" {
		return fmt.Sprintf(`{"text": "%s", "channel":"%s"}`, message, channel)
	}
	return fmt.Sprintf(`{"text": "%s"}`, message)
}
