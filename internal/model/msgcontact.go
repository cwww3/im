package model

import "time"

type MessageContact struct {
	OwnerUid uint64
	OtherUid uint64
	Mid uint64
	Type uint64
	CreateTime time.Time
}

