package media

import (
	"github.com/google/uuid"
	monify "monify/protobuf/gen/go"
	"time"
)

type TmpFile struct {
	Id            uuid.UUID
	ExpectedUsage monify.Usage
	Uploader      uuid.UUID
	UploadedAt    time.Time
	Path          string
}

type ConfirmedFile struct {
	Id          uuid.UUID
	Usage       monify.Usage
	Uploader    uuid.UUID
	UploadedAt  time.Time
	Path        string
	ConfirmedAt time.Time
}
