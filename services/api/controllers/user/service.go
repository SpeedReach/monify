package user

import (
	monify "monify/protobuf/gen/go"
)

type Service struct {
	monify.UserServiceServer
}
