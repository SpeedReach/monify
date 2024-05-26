package auth

import (
	monify "monify/protobuf/gen/go"
)

type Service struct {
	Secret string
	monify.UnimplementedAuthServiceServer
}
