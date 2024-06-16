package auth

import (
	"github.com/google/uuid"
	monify "monify/protobuf/gen/go"
)

type Service struct {
	Secret string
	monify.UnimplementedAuthServiceServer
}

func GenerateRefreshToken(_ uuid.UUID) string {
	return uuid.New().String()
}
