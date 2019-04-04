package slackwebhook

import "github.com/ashwanthkumar/slack-go-webhook"
import (
	"fmt"
	"strconv"
)

func SlackWebHook(webhook string, clientname string, clientdate string, clientpayment string, expirationdate string, expired bool) {
	webhookUrl := webhook

	color := "#2eb886"
	if expired == true {
		color = "#d80000"
	}

	text := string("Client Name: " + clientname + "\nPayment Date: " + clientdate + "\nPayment Amount: $" + clientpayment + "\nExpiration Time: " + expirationdate + "\nExpired: " + strconv.FormatBool(expired))

	attachment1 := slack.Attachment { Color:&color, Text:&text }

	payload := slack.Payload {
		Username: "payment-monitor",
		Channel: "#payment-monitor",
		IconEmoji: ":moneybag:",
		Attachments: []slack.Attachment{attachment1},
	}
	err := slack.Send(webhookUrl, "", payload)
	if len(err) > 0 {
		fmt.Printf("error: %s\n", err)
	}
}
