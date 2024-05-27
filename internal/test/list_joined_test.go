package test

import (
	"context"
	monify "monify/protobuf/gen/go"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestListJoined(t *testing.T) {
	client := GetTestClient(t)
	_ = client.CreateTestUser()
	group1, err := client.CreateGroup(context.TODO(), &monify.CreateGroupRequest{Name: "test1"})
	group2, err := client.CreateGroup(context.TODO(), &monify.CreateGroupRequest{Name: "test2"})
	groups, err := client.ListJoinedGroups(context.TODO(), &monify.Empty{})
	assert.NoError(t, err)
	assert.Equal(t, groups.GetGroups()[0].GetGroupId(), group1.GroupId)
	assert.Equal(t, groups.GetGroups()[0].GetName(), "test1")
	assert.Equal(t, groups.GetGroups()[1].GetGroupId(), group2.GroupId)
	assert.Equal(t, groups.GetGroups()[1].GetName(), "test2")
}
