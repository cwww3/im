package service

import "github.com/cwww3/im/internal/model"

type IMessageService interface {
	SendNewMsg(senderUid, recipientUid uint64, content string , msgType int) (model.Message,error)
	QueryConversationMsg(ownerUid, otherUid uint64) ([]model.Message,error)
	QueryNewerMsgFrom(ownerUid, otherUid, lastMid uint64) ([]model.Message,error)
	QueryContacts(ownerUid uint64) ([]model.Message,error)
}

