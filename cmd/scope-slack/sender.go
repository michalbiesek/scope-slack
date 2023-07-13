package main

import (
	"github.com/slack-go/slack"
)

type SlackSender struct {
	client    *slack.Client
	channelId string
}

func getChannelId() string {
	return "C05HBUJTC9W"
}

// NewSlackSender
func NewSlackSender(token string, channelName string) *SlackSender {
	ss := &SlackSender{
		client:    slack.New(token),
		channelId: getChannelId(),
	}

	return ss
}

func (s *SlackSender) SendNotification() {
	s.client.PostMessage(s.channelId,
		slack.MsgOptionText("Hello world from AppScope Donn", false),
		slack.MsgOptionAttachments(),
		slack.MsgOptionAsUser(false))
}
