package test

import (
	"context"
	monify "monify/protobuf/gen/go"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestListJoined(t *testing.T) {
	client := GetTestClient(t)
	userId1 := client.CreateTestUser()
	group, err := client.CreateGroup(context.TODO(), &monify.CreateGroupRequest{Name: "test"})
	client.SetTestUser(userId1)
	groups, err := client.ListJoinedGroups(context.TODO(), &monify.Empty{})
	assert.NoError(t, err)
	assert.Equal(t, groups.GetGroups()[0].GetGroupId(), group.GroupId)
	assert.Equal(t, groups.GetGroups()[0].GetName(), "test")
}
