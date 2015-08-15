package goslack

import (
	"bytes"
	"errors"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
)

// Send message to Slack
func (slack *SlackService) Send(message, channel string) error {
	urlRequest := slack.WebhookURL
	data := url.Values{}
	data.Set("payload", slack.buildTextJSON(message, channel))

	r, err := http.NewRequest("POST", urlRequest, bytes.NewBufferString(data.Encode()))
	r.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	r.Header.Add("Content-Length", strconv.Itoa(len(data.Encode())))

	if err != nil {
		return err
	}

	resp, err := slack.HTTPClient.Do(r)
	if err != nil {
		return err
	}

	if resp.StatusCode != 200 {
		respText, _ := ioutil.ReadAll(resp.Body)
		return errors.New(string(respText))
	}

	return nil
}
