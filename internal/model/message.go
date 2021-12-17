package model

import (
	"time"
)

type Message struct {
	mid         uint64
	Content     string
	SenderId    uint64
	RecipientId uint64
	MsgType     uint64
	CreateTime time.Time
}
