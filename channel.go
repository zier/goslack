package goslack

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"
)

// HistoryResponse response from slack
type HistoryResponse struct {
	Ok       bool
	Messages []map[string]string
	HasMore  string
	Error    string
}

// GetHistoryFromChannel get history messages from channel
func (slack *SlackService) GetHistoryFromChannel(channelID string, startDate time.Time, endDate time.Time, limit int64) (*HistoryResponse, error) {
	start := strconv.Itoa(int(startDate.Unix()))
	end := strconv.Itoa(int(endDate.Unix()))

	urlRequest := fmt.Sprintf("https://slack.com/api/channels.history?token=%s&channel=%s&oldest=%s&latest=%s", slack.Token, channelID, start, end)

	r, err := http.NewRequest("GET", urlRequest, nil)

	if err != nil {
		return nil, err
	}

	resp, err := slack.HTTPClient.Do(r)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != 200 {
		respText, _ := ioutil.ReadAll(resp.Body)

		return nil, errors.New(string(respText))
	}

	dataResp, _ := ioutil.ReadAll(resp.Body)

	hr := &HistoryResponse{}
	err = json.Unmarshal(dataResp, &hr)
	if err != nil {
		return nil, err
	}

	return hr, nil
}
