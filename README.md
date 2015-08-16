# goslack

###Features
- Send message to slack ([WebHook](https://api.slack.com/incoming-webhooks))
- Get history messages from channel ([SlackAPI](https://api.slack.com/methods/channels.history))

###Example New
if you want only get message from channel don't want to send to you can empty string to url web hook


`
	slack := New("yourSlackUrlHook", "yourSlackToken", nil)
`



###Example send message (use default channel in webhook setting)
`
	slack.Send("Hello World","")
`


###Example send with channel

`
slack.Send("Hello World","#game")
`

###Example get message from channel
you can read more about response message value in document channels.history

```
//HistoryResponse response from slack
type HistoryResponse struct {
	Ok       bool
	Messages []map[string]string
	HasMore  string
}
```

`
resp, err := slack.GetHistoryFromChannel("CHANNELID", startDate, endDate, limit)
`

* you can find your channel id [here](https://api.slack.com/methods/channels.list)
