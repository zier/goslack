# goslack

golang api overide data slack web hook for send message
 
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

you can overide data by method set before send message

`
slack.SetChannel("@john").SetMessage("Hello World").Send()
`