package test

import (
	"context"
	monify "monify/protobuf/gen/go"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateGroup(t *testing.T) {
	client := GetTestClient(t)
	_ = client.CreateTestUser()
	group, err := client.CreateGroup(context.Background(), &monify.CreateGroupRequest{Name: "test"})
	assert.NoError(t, err)
	assert.NotEmpty(t, group)
	code, err := client.GenerateInviteCode(context.TODO(), &monify.GenerateInviteCodeRequest{GroupId: group.GroupId})
	assert.NoError(t, err)
	user1 := client.CreateTestUser()

	joinGroup, err := client.JoinGroup(context.TODO(), &monify.JoinGroupRequest{InviteCode: code.InviteCode})
	assert.NoError(t, err)

	groups, err := client.ListJoinedGroups(context.TODO(), &monify.Empty{})
	assert.NoError(t, err)
	assert.Equal(t, joinGroup.GetGroupId(), groups.Groups[0].GroupId)

	client.CreateTestUser()
	_, err = client.DeleteGroup(context.TODO(), &monify.DeleteGroupRequest{GroupId: group.GroupId})
	assert.Error(t, err)
	client.SetTestUser(user1)
	groups, err = client.ListJoinedGroups(context.TODO(), &monify.Empty{})
	assert.NoError(t, err)
	assert.Equal(t, joinGroup.GetGroupId(), groups.Groups[0].GroupId)

	_, err = client.DeleteGroup(context.TODO(), &monify.DeleteGroupRequest{GroupId: group.GroupId})
	assert.NoError(t, err)
	groups, err = client.ListJoinedGroups(context.TODO(), &monify.Empty{})
	assert.NoError(t, err)
	assert.Empty(t, groups.Groups)
}

func TestGroupInviteCode(t *testing.T) {
	client := GetTestClient(t)
	_ = client.CreateTestUser()
	group, err := client.CreateGroup(context.Background(), &monify.CreateGroupRequest{Name: "test"})
	assert.NoError(t, err)
	assert.NotEmpty(t, group)
	code, err := client.GenerateInviteCode(context.Background(), &monify.GenerateInviteCodeRequest{GroupId: group.GroupId})
	assert.NoError(t, err)
	assert.NotEmpty(t, code)
	inviteCode, err := client.GetInviteCode(context.Background(), &monify.GetInviteCodeRequest{GroupId: group.GroupId})
	assert.NoError(t, err)
	assert.Equal(t, code.InviteCode, inviteCode.InviteCode)

	_, err = client.DeleteInviteCode(context.Background(), &monify.DeleteInviteCodeRequest{GroupId: group.GroupId})
	assert.NoError(t, err)
	inviteCode, err = client.GetInviteCode(context.Background(), &monify.GetInviteCodeRequest{GroupId: group.GroupId})
	assert.Error(t, err)
	assert.Nil(t, inviteCode)
}