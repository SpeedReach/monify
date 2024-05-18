package test

import (
	"context"
	"github.com/stretchr/testify/assert"
	monify "monify/protobuf"
	"testing"
)

func TestLogin(t *testing.T) {
	email := "brian030128@gmail.com"
	client := GetTestClient(t)
	_, err := client.EmailLogin(context.TODO(), &monify.EmailLoginRequest{Email: email})
	assert.Error(t, err)
	res1, err := client.EmailRegister(context.TODO(), &monify.EmailRegisterRequest{Email: email})
	assert.NoError(t, err)
	res2, err := client.EmailLogin(context.TODO(), &monify.EmailLoginRequest{Email: email})
	assert.NoError(t, err)
	assert.Equal(t, res1.UserId, res2.UserId)
}
