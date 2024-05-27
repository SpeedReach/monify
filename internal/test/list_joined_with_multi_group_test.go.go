package test

import (
	"context"
	monify "monify/protobuf/gen/go"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestListJoinedWithMultiGroup(t *testing.T) {
	client := GetTestClient(t)
	user1 := client.CreateTestUser()
	group1, err := client.CreateGroup(context.TODO(), &monify.CreateGroupRequest{Name: "test1"})
	group2, err := client.CreateGroup(context.TODO(), &monify.CreateGroupRequest{Name: "test2"})
	user2 := client.CreateTestUser()
	group3, err := client.CreateGroup(context.TODO(), &monify.CreateGroupRequest{Name: "test3"})
	group4, err := client.CreateGroup(context.TODO(), &monify.CreateGroupRequest{Name: "test4"})
	client.SetTestUser(user1)
	groups, err := client.ListJoinedGroups(context.TODO(), &monify.Empty{})
	assert.NoError(t, err)
	assert.Equal(t, groups.GetGroups()[0].GetGroupId(), group1.GroupId)
	assert.Equal(t, groups.GetGroups()[0].GetName(), "test1")
	assert.Equal(t, groups.GetGroups()[1].GetGroupId(), group2.GroupId)
	assert.Equal(t, groups.GetGroups()[1].GetName(), "test2")
	client.SetTestUser(user2)
	groups, err = client.ListJoinedGroups(context.TODO(), &monify.Empty{})
	assert.NoError(t, err)
	assert.Equal(t, groups.GetGroups()[0].GetGroupId(), group3.GroupId)
	assert.Equal(t, groups.GetGroups()[0].GetName(), "test3")
	assert.Equal(t, groups.GetGroups()[1].GetGroupId(), group4.GroupId)
	assert.Equal(t, groups.GetGroups()[1].GetName(), "test4")
}
