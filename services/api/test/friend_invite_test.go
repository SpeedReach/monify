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
	assert.Error(t, err)
	_ = client.CreateTestUser()
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
}
