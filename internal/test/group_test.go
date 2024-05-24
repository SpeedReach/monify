package test

import (
	"context"
	monify "monify/protobuf"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateGroup(t *testing.T) {
	client := GetTestClient(t)
	userId1 := client.CreateTestUser()
	group, err := client.CreateGroup(context.TODO(), &monify.CreateGroupRequest{Name: "test"})
	assert.NoError(t, err)
	assert.NotEmpty(t, group)
	//userId2 := client.CreateTestUser()
	client.SetTestUser(userId1)
}
