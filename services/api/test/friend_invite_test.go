package test

import (
	"context"
	monify "monify/protobuf/gen/go"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInviteFriend(t *testing.T) {
	// test invite
	client := GetTestClient(t)
	user1 := client.CreateTestUser()
	_, err := client.UpdateUserNickId(context.Background(), &monify.UpdateUserNickIdRequest{NickId: "test_nickId1"})
	assert.NoError(t, err)
	_, err = client.InviteFriend(context.TODO(), &monify.InviteFriendRequest{ReceiverNickId: "test_nickId1"})
	assert.Error(t, err) // cannot send invitation to yourself
	user2 := client.CreateTestUser()
	_, err = client.InviteFriend(context.TODO(), &monify.InviteFriendRequest{ReceiverNickId: "test_nickId1"})
	assert.NoError(t, err)
	user3 := client.CreateTestUser()
	_, err = client.InviteFriend(context.TODO(), &monify.InviteFriendRequest{ReceiverNickId: "test_nickId1"})
	assert.NoError(t, err)
	_ = client.CreateTestUser()
	_, err = client.InviteFriend(context.TODO(), &monify.InviteFriendRequest{ReceiverNickId: "test_nickId1"})
	assert.NoError(t, err)

	// test list invitation
	client.SetTestUser(user1)
	invitaions, err := client.ListFriendInvitation(context.TODO(), &monify.FriendEmpty{})
	assert.NoError(t, err)
	assert.Equal(t, len(invitaions.GetInvitation()), 3)

	// test accept invitation
	var inviteId []string
	for i := 0; i < len(invitaions.GetInvitation()); i++ {
		inviteId = append(inviteId, (invitaions.Invitation[i].InviteId))
	}
	_, err = client.AcceptInvitation(context.TODO(), &monify.AcceptInvitationRequest{User1Id: user1, User2Id: user2, InviteId: inviteId[0]})
	assert.NoError(t, err)
	_, err = client.AcceptInvitation(context.TODO(), &monify.AcceptInvitationRequest{User1Id: user1, User2Id: user3, InviteId: inviteId[1]})
	assert.NoError(t, err)

	// test reject invitation
	_, err = client.RejectInvitation(context.TODO(), &monify.RejectInvitationRequest{InviteId: inviteId[2]})
	assert.NoError(t, err)
	invitaions, err = client.ListFriendInvitation(context.TODO(), &monify.FriendEmpty{})
	assert.Equal(t, len(invitaions.GetInvitation()), 2)

	// test list friend
	friends, err := client.ListFriend(context.TODO(), &monify.FriendEmpty{})
	assert.NoError(t, err)
	assert.Equal(t, len(friends.GetFriends()), 2)

	// test delete friend
	var friendS_relationId []string
	for _, friend := range friends.GetFriends() {
		friendS_relationId = append(friendS_relationId, friend.RelationId)
	}
	_, err = client.DeleteFriend(context.TODO(), &monify.DeleteFriendRequest{RelationId: friendS_relationId[0]})
	assert.NoError(t, err)
	friends, err = client.ListFriend(context.TODO(), &monify.FriendEmpty{})
	assert.NoError(t, err)
	assert.Equal(t, len(friends.GetFriends()), 1)

}
