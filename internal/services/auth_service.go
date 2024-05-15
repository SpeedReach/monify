package services

import (
	"context"
	monify "monify/protobuf"
)

type AuthService struct {
	monify.UnimplementedAuthServiceServer
}

func (s AuthService) EmailLogin(context.Context, *monify.EmailLoginRequest) (*monify.EmailLoginResponse, error) {
	panic("todo")
}
