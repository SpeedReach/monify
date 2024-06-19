package media

import (
	"github.com/google/uuid"
	"time"
)

type TmpImage struct {
	Id            uuid.UUID
	ExpectedUsage Usage
	Uploader      uuid.UUID
	UploadedAt    time.Time
	Path          string
}

type ConfirmedImage struct {
	Id          uuid.UUID
	Usage       Usage
	Uploader    uuid.UUID
	UploadedAt  time.Time
	Path        string
	ConfirmedAt time.Time
}
