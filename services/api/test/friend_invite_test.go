package test

import (
	"context"
	monify "monify/protobuf/gen/go"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInviteFriend(t *testing.T) {
	// Register a receiver
	email := "daniel0702chien@gmail.com"
	password := "qwer1234"
	client := GetTestClient(t)
	_, err := client.EmailRegister(context.TODO(), &monify.EmailRegisterRequest{Email: email, Password: password})
	assert.NoError(t, err)

	_ = client.CreateTestUser()
	_, err = client.InviteFriend(context.TODO(), &monify.InviteFriendRequest{ReceiverEmail: email})
	assert.NoError(t, err)

}
