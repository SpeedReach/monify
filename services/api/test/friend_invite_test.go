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
	_ = client.CreateTestUser()
	_, err = client.InviteFriend(context.TODO(), &monify.InviteFriendRequest{ReceiverNickId: "test_nickId1"})
	assert.NoError(t, err)

	// test list invitation
	client.SetTestUser(user1)
	invitaions, err := client.ListFriendInvitation(context.TODO(), &monify.FriendEmpty{})
	assert.NoError(t, err)
	assert.Equal(t, len(invitaions.GetInvitation()), 2)

	// test accept invitation
	var inviteId []string
	for i := 0; i < len(invitaions.GetInvitation()); i++ {
		inviteId = append(inviteId, (invitaions.Invitation[i].InviteId))
	}
	_, err = client.AcceptInvitation(context.TODO(), &monify.AcceptInvitationRequest{User1Id: user1, User2Id: user2, InviteId: inviteId[0]})
	assert.NoError(t, err)
}
