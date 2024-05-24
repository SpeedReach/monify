package test

import (
	"context"
	"github.com/stretchr/testify/assert"
	monify "monify/protobuf"
	"testing"
)

func TestCreateGroup(t *testing.T) {
	client := GetTestClient(t)
	_ = client.CreateTestUser()
	group, err := client.CreateGroup(context.TODO(), &monify.CreateGroupRequest{Name: "test"})
	assert.NoError(t, err)
	assert.NotEmpty(t, group)
}
