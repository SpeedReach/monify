package auth

import (
	monify "monify/protobuf"
)

type Service struct {
	Secret string
	monify.UnimplementedAuthServiceServer
}
