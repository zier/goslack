package goslack

import (
	"fmt"
	"log"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func newTestSlackService(testClient *http.Client) *SlackService {
	slack, _ := New("http://hooks.slack.com", "TEST_TOKEN", testClient)
	return slack
}

func testServer(code int, body string) (*httptest.Server, *http.Client) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(code)
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintln(w, body)
	}))

	u, err := url.Parse(server.URL)
	if err != nil {
		log.Fatalln("failed to parse httptest.Server URL:", err)
	}

	return server, &http.Client{Transport: RewriteTransport{URL: u}}
}

func TestSendMessage(t *testing.T) {
	server, client := testServer(200, ``)
	defer server.Close()

	slack := newTestSlackService(client)
	err := slack.Send("Hello World", "")

	assert.NoError(t, err)
}

func TestGetHistoryFromChannel(t *testing.T) {
	server, client := testServer(200,
		`{
		    "ok": true,
		    "messages": [
		        {
		            "text": "hello",
		            "username": "incoming-webhook",
		            "bot_id": "B09065XXX",
		            "type": "message",
		            "subtype": "bot_message",
		            "ts": "1439495572.194790"
		        },
		        {
		            "text": "i'm test",
		            "username": "incoming-webhook",
		            "bot_id": "B09065XXX",
		            "type": "message",
		            "ts": "1439495672.000002"
		        }
		    ],
		    "has_more": false
		}`)
	defer server.Close()

	slack := newTestSlackService(client)
	hr, err := slack.GetHistoryFromChannel("CHANNELID", time.Now().AddDate(0, 0, -1), time.Now(), 100)
	assert.NoError(t, err)
	assert.True(t, hr.Ok)
	assert.Equal(t, len(hr.Messages), 2)
	assert.Equal(t, hr.Messages[0]["text"], "hello")
}
