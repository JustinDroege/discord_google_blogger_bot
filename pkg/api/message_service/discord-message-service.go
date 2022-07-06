package message_service

import (
	"fmt"
	"github.com/JustinDroege/BloggerBot/pkg/utils"
	"github.com/bwmarrin/discordgo"
)

type DiscordMessageService struct {
	DiscordSession       *discordgo.Session
	ChannelIDs           *[]string
	MaximumMessageLength int
}

func (d *DiscordMessageService) SendMessage(items *[]utils.Item) error {
	messages, err := d.convertToMessages(items)
	if err != nil {
		return err
	}

	summedUpMessages := d.sumUpMessages(messages)
	for _, message := range *summedUpMessages {
		for _, channelID := range *d.ChannelIDs {
			_, err := d.DiscordSession.ChannelMessageSend(channelID, message)
			if err != nil {
				return err
			}
		}
	}

	return nil
}

func (d *DiscordMessageService) sumUpMessages(messages *[]string) *[]string {
	var summedUpMessages []string

	currentSize := 0
	for _, message := range *messages {
		if currentSize+len(message) > d.MaximumMessageLength-2 || len(summedUpMessages) == 0 {
			summedUpMessages = append(summedUpMessages, message)
			currentSize = len(message)
		} else {
			summedUpMessages[len(summedUpMessages)-1] += "\n" + message
			currentSize += len(summedUpMessages[len(summedUpMessages)-1])
		}
	}

	return &summedUpMessages
}

func (d *DiscordMessageService) convertToMessages(items *[]utils.Item) (*[]string, error) {
	var messages []string

	for _, item := range *items {
		switch item.Type {
		case utils.HeadLine:
			messages = append(messages, "***"+item.Value+"***")
		case utils.Text:
			messages = append(messages, item.Value)
		default:
			return nil, fmt.Errorf("unknown item type")
		}
	}

	return &messages, nil
}
