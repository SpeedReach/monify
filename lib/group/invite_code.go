package group

import (
	"github.com/google/uuid"
	"time"
)

const (
	ExpiresInterval = time.Minute * 10
)

type InviteCode struct {
	GroupId   uuid.UUID
	Code      string
	CreatedAt time.Time
}

func (i InviteCode) IsExpired() bool {
	return i.CreatedAt.Add(ExpiresInterval).Before(time.Now())
}

func (i InviteCode) ExpiresAfter() time.Duration {
	return i.CreatedAt.Add(ExpiresInterval).Sub(time.Now())
}
