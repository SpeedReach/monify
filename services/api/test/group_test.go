package test

import (
	"context"
	"github.com/stretchr/testify/assert"
	monify "monify/protobuf/gen/go"
	"testing"
)

func TestCreateGroup(t *testing.T) {
	client := GetTestClient(t)
	_ = client.CreateTestUser()
	group, err := client.CreateGroup(context.Background(), &monify.CreateGroupRequest{Name: "test", Description: "cool desc"})
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
	assert.Equal(t, "test", groups.Groups[0].Name)
	assert.Equal(t, "cool desc", groups.Groups[0].Description)

	client.CreateTestUser()
	_, err = client.DeleteGroup(context.TODO(), &monify.DeleteGroupRequest{GroupId: group.GroupId})
	assert.Error(t, err)
	client.SetTestUser(user1)
	groups, err = client.ListJoinedGroups(context.TODO(), &monify.Empty{})
	assert.NoError(t, err)
	assert.Equal(t, joinGroup.GetGroupId(), groups.Groups[0].GroupId)
	assert.Equal(t, "test", groups.Groups[0].Name)
	assert.Equal(t, "cool desc", groups.Groups[0].Description)

	_, err = client.DeleteGroup(context.TODO(), &monify.DeleteGroupRequest{GroupId: group.GroupId})
	assert.NoError(t, err)
	groups, err = client.ListJoinedGroups(context.TODO(), &monify.Empty{})
	assert.NoError(t, err)
	assert.Empty(t, groups.Groups)
}

func TestGroupInviteCode(t *testing.T) {
	client := GetTestClient(t)
	_ = client.CreateTestUser()
	group, err := client.CreateGroup(context.Background(), &monify.CreateGroupRequest{Name: "test", Description: "test123"})
	assert.NoError(t, err)
	assert.NotEmpty(t, group)
	code, err := client.GenerateInviteCode(context.Background(), &monify.GenerateInviteCodeRequest{GroupId: group.GroupId})
	assert.NoError(t, err)
	assert.NotEmpty(t, code)
	inviteCode, err := client.GetInviteCode(context.Background(), &monify.GetInviteCodeRequest{GroupId: group.GroupId})
	assert.NoError(t, err)
	assert.Equal(t, code.InviteCode, inviteCode.InviteCode)

	groupBrief, err := client.GetGroupByInviteCode(context.Background(), &monify.GetGroupByInviteCodeRequest{InviteCode: code.InviteCode})
	assert.NoError(t, err)
	assert.Equal(t, group.GroupId, groupBrief.GroupId)
	assert.Equal(t, "test", groupBrief.Name)
	assert.Equal(t, "test123", groupBrief.Description)

	groupBrief, err = client.GetGroupInfo(context.Background(), &monify.GetGroupInfoRequest{GroupId: group.GroupId})
	assert.NoError(t, err)
	assert.Equal(t, group.GroupId, groupBrief.GroupId)
	assert.Equal(t, "test", groupBrief.Name)
	assert.Equal(t, "test123", groupBrief.Description)

	_, err = client.DeleteInviteCode(context.Background(), &monify.DeleteInviteCodeRequest{GroupId: group.GroupId})
	assert.NoError(t, err)
	inviteCode, err = client.GetInviteCode(context.Background(), &monify.GetInviteCodeRequest{GroupId: group.GroupId})
	assert.Error(t, err)
	assert.Nil(t, inviteCode)
}
