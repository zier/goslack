package goslack

import (
	"bytes"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
)

type SlackService struct {
	WebhookURL string
	Username   string
	Icon       string
	Channel    string
	Message    string
}

func New(url, username, icon, channel string) (*SlackService, error) {
	if url == "" || username == "" || icon == "" || channel == "" {
		return nil, errors.New("invalid data")
	}

	return &SlackService{
		WebhookURL: url,
		Username:   username,
		Icon:       icon,
		Channel:    channel,
	}, nil
}

func (slack *SlackService) SetIcon(icon string) *SlackService {
	slack.Icon = icon
	return slack
}

func (slack *SlackService) SetChannel(channel string) *SlackService {
	slack.Channel = channel
	return slack
}

func (slack *SlackService) SetUsername(username string) *SlackService {
	slack.Username = username
	return slack
}

func (slack *SlackService) SetMessage(message string) *SlackService {
	slack.Message = message
	return slack
}

func (slack *SlackService) Send() error {
	urlRequest := slack.WebhookURL
	data := url.Values{}
	data.Set("payload", slack.buildTextJSON())

	client := &http.Client{}
	r, err := http.NewRequest("POST", urlRequest, bytes.NewBufferString(data.Encode()))
	r.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	r.Header.Add("Content-Length", strconv.Itoa(len(data.Encode())))

	if err != nil {
		return err
	}

	resp, err := client.Do(r)
	if err != nil {
		return err
	}

	if resp.StatusCode != 200 {
		respText, _ := ioutil.ReadAll(resp.Body)
		return errors.New(string(respText))
	}

	return nil
}

func (slack *SlackService) buildTextJSON() string {
	return fmt.Sprintf(`{"text": "%s", "username":"%s", "icon_emoji":"%s", "channel":"%s"}`, slack.Message, slack.Username, slack.Icon, slack.Channel)
}
