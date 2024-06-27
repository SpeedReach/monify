package friend

import (
	monify "monify/protobuf/gen/go"
)

type Service struct {
	monify.UnimplementedFriendServiceServer
}
