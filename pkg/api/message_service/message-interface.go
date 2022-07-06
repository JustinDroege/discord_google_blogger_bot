package message_service

import "github.com/JustinDroege/BloggerBot/pkg/utils"

type MessageService interface {
	SendMessage(items *[]utils.Item) error
}
