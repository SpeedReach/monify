package group

import (
	monify "monify/protobuf/gen/go"
)

type Service struct {
	monify.UnimplementedGroupServiceServer
}
