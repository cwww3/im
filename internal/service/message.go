package service

import (
	"github.com/cwww3/im/internal/model"
	"github.com/cwww3/im/internal/repository"
)

type MessageService struct {
	messageRepository *repository.MessageRepository
}

func (ms *MessageService) SendNewMsg(senderUid, recipientUid uint64, content string, msgType int) (model.Message, error) {
	panic("implement me")
}

func (ms *MessageService) QueryConversationMsg(ownerUid, otherUid uint64) ([]model.Message, error) {
	panic("implement me")
}

func (ms *MessageService) QueryNewerMsgFrom(ownerUid, otherUid, lastMid uint64) ([]model.Message, error) {
	panic("implement me")
}

func (ms *MessageService) QueryContacts(ownerUid uint64) ([]model.Message, error) {
	panic("implement me")
}

func NewMessageService (messageRepository *repository.MessageRepository) *MessageService {
	return &MessageService{
		messageRepository: messageRepository,
	}
}


