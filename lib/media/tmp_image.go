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
	URL           string
}

type ConfirmedImage struct {
	Id          uuid.UUID
	Usage       Usage
	Uploader    uuid.UUID
	UploadedAt  time.Time
	URL         string
	ConfirmedAt time.Time
}
