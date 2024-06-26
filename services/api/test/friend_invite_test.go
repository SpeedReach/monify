package test

import (
	"context"
	monify "monify/protobuf/gen/go"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInviteFriend(t *testing.T) {
	client := GetTestClient(t)
	_ = client.CreateTestUser()
	_, err := client.UpdateUserNickId(context.Background(), &monify.UpdateUserNickIdRequest{NickId: "test_nickId"})
	assert.NoError(t, err)
	//_, err = client.InviteFriend(context.TODO(), &monify.InviteFriendRequest{ReceiverNickId: "test_nickId"})
	//assert.NoError(t, err)

}
