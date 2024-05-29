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
	createdGroups := map[string]string{}
	group1, err := client.CreateGroup(context.TODO(), &monify.CreateGroupRequest{Name: "test1"})
	assert.NoError(t, err)
	createdGroups["test1"] = group1.GroupId
	group2, err := client.CreateGroup(context.TODO(), &monify.CreateGroupRequest{Name: "test2"})
	createdGroups["test2"] = group2.GroupId

	_ = client.CreateTestUser()
	group3, err := client.CreateGroup(context.TODO(), &monify.CreateGroupRequest{Name: "test3"})
	assert.NoError(t, err)
	group4, err := client.CreateGroup(context.TODO(), &monify.CreateGroupRequest{Name: "test4"})
	assert.NoError(t, err)
	createdGroups["test3"] = group3.GroupId
	createdGroups["test4"] = group4.GroupId

	groups, err := client.ListJoinedGroups(context.TODO(), &monify.Empty{})
	assert.NoError(t, err)
	assert.Equal(t, 2, len(groups.GetGroups()))
	for _, group := range groups.GetGroups() {
		assert.Equal(t, createdGroups[group.GetName()], group.GetGroupId())
	}

	client.SetTestUser(user1)
	groups, err = client.ListJoinedGroups(context.TODO(), &monify.Empty{})
	assert.NoError(t, err)
	assert.Equal(t, 2, len(groups.GetGroups()))
	for _, group := range groups.GetGroups() {
		assert.Equal(t, createdGroups[group.GetName()], group.GetGroupId())
	}
}
