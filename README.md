# goslack

golang api use slack web hook for send message to slack
 
you can read more webhook slack api -> [doc](https://api.slack.com/incoming-webhooks)

###Example New


`
	slack := New("yourSlackUrlHook", "yourName", ":yourIconEmoji:", "#yourChannel")
`


if you want to send to user you can set 

`
slack := New("yourSlackUrlHook", "yourName", ":yourIconEmoji:", "#john")
`

###Example Send Message
`
	slack.SetMessage("Hello World").Send()
`

###Example Set

method set data before send message

`
slack.SetChannel("@john").SetMessage("Hello World").Send()
`