package test

import (
	"context"
	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc/metadata"
	monify "monify/protobuf"
	"testing"
)

func TestCreateGroup(t *testing.T) {
	client := GetTestClient(t)
	_, accessToken := client.CreateTestUser()
	md := metadata.Pairs("authorization", "Bearer "+accessToken)
	ctx := metadata.NewOutgoingContext(context.Background(), md)

	group, err := client.CreateGroup(ctx, &monify.CreateGroupRequest{Name: "test"})
	assert.NoError(t, err)
	assert.NotEmpty(t, group.GroupId)
}
