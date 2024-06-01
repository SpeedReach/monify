package test

import (
	"context"
	monify "monify/protobuf/gen/go"

	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLogin(t *testing.T) {
	email := "brian030128@gmail.com"
	password := "qwer1234"
	client := GetTestClient(t)
	_, err := client.EmailLogin(context.TODO(), &monify.EmailLoginRequest{Email: email, Password: password})
	assert.Error(t, err)
	res1, err := client.EmailRegister(context.TODO(), &monify.EmailRegisterRequest{Email: email, Password: password})
	assert.NoError(t, err)
	res2, err := client.EmailLogin(context.TODO(), &monify.EmailLoginRequest{Email: email, Password: password})
	assert.NoError(t, err)
	assert.Equal(t, res1.UserId, res2.UserId)
	res2, err = client.EmailLogin(context.TODO(), &monify.EmailLoginRequest{Email: email, Password: "badpassword"})
	assert.Error(t, err)
}
func TestRefreshToken(t *testing.T) {
	email := "a0905373664@gmail.com"
	password := "qwer1234"
	client := GetTestClient(t)
	res1, err := client.EmailRegister(context.TODO(), &monify.EmailRegisterRequest{Email: email, Password: password})
	assert.NoError(t, err)
	res2, err := client.EmailLogin(context.TODO(), &monify.EmailLoginRequest{Email: email, Password: password})
	assert.NoError(t, err)
	assert.Equal(t, res1.UserId, res2.UserId)

	tokens1, err := client.RefreshToken(context.TODO(), &monify.RefreshTokenRequest{RefreshToken: res2.RefreshToken})
	assert.NoError(t, err)
	assert.NotEmpty(t, tokens1)

	tokens2, err := client.RefreshToken(context.TODO(), &monify.RefreshTokenRequest{RefreshToken: tokens1.RefreshToken})
	assert.NoError(t, err)
	assert.NotEmpty(t, tokens2)

}
