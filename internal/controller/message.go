package controller

import (
	"encoding/json"
	"github.com/cwww3/im/internal/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

type MessageController struct {
	messageService service.IMessageService
}

func NewMessageController(messageService service.IMessageService) *MessageController {
	return &MessageController{messageService: messageService}
}

func (ms *MessageController) SendMessage(ctx *gin.Context) {
	params := new(struct {
		SenderUid    uint64 `json:"senderUid"`
		RecipientUid uint64 `json:"recipientUid"`
		Content      string `json:"content"`
		MsgType      int    `json:"msgType"`
	})
	err := ctx.Bind(params)
	msg, err := ms.messageService.SendNewMsg(params.SenderUid, params.RecipientUid, params.Content, params.MsgType)
	if err != nil {
		ctx.String(http.StatusOK, "")
		return
	}
	data, _ := json.Marshal(msg)
	ctx.String(http.StatusOK, string(data))
}

func (ms *MessageController) QueryMsg(ctx *gin.Context) {
	params := new(struct {
		OwnerUid uint64 `json:"ownerUid"`
		OtherUid uint64 `json:"otherUid"`
	})
	err := ctx.Bind(params)
	msgList, err := ms.messageService.QueryConversationMsg(params.OwnerUid, params.OtherUid)
	if err != nil {
		ctx.String(http.StatusOK, "")
		return
	}
	data, _ := json.Marshal(msgList)
	ctx.String(http.StatusOK, string(data))
}

func (ms *MessageController) QueryMsgSinceMid(ctx *gin.Context) {
	params := new(struct {
		OwnerUid uint64 `json:"ownerUid"`
		OtherUid uint64 `json:"otherUid"`
		LastMid  uint64 `json:"lastMid"`
	})
	err := ctx.Bind(params)
	msgList, err := ms.messageService.QueryNewerMsgFrom(params.OwnerUid, params.OtherUid, params.LastMid)
	if err != nil {
		ctx.String(http.StatusOK, "")
		return
	}
	data, _ := json.Marshal(msgList)
	ctx.String(http.StatusOK, string(data))

}

func (ms *MessageController) QueryContacts(ctx *gin.Context) {
	params := new(struct {
		OwnerUid uint64 `json:"ownerUid"`
	})
	err := ctx.Bind(params)
	msgList, err := ms.messageService.QueryContacts(params.OwnerUid)
	if err != nil {
		ctx.String(http.StatusOK, "")
		return
	}
	data, _ := json.Marshal(msgList)
	ctx.String(http.StatusOK, string(data))
}
