package message_service

import (
	"fmt"
	"github.com/JustinDroege/BloggerBot/pkg/utils"
)

type MessageBroker struct {
	messageServices *[]MessageService
}

func New(messageServices *[]MessageService) *MessageBroker {
	return &MessageBroker{
		messageServices: messageServices,
	}
}

func (m *MessageBroker) AppendMessageService(messageService MessageService) {
	for _, service := range *m.messageServices {
		if service == messageService {
			return
		}
	}

	*m.messageServices = append(*m.messageServices, messageService)
}

func (m *MessageBroker) GetMessageServices() []MessageService {
	return *m.messageServices
}

func (m *MessageBroker) RemoveMessageService(messageService MessageService) {
	for i, service := range *m.messageServices {
		if service == messageService {
			*m.messageServices = append((*m.messageServices)[:i], (*m.messageServices)[i+1:]...)
			return
		}
	}
}

func (m *MessageBroker) SendMessage(items *[]utils.Item) {
	for _, messageService := range *m.messageServices {
		err := messageService.SendMessage(items)
		if err != nil {
			fmt.Print(err)
		}
	}
}
