package test

import (
	"context"
	"github.com/stretchr/testify/assert"
	monify "monify/protobuf/gen/go"
	"testing"
)

func TestChangeUsername(t *testing.T) {
	client := GetTestClient(t)
	userId := client.CreateTestUser()
	ctx := context.Background()

	_, err := client.UpdateUserName(ctx, &monify.UpdateUserNameRequest{Name: "testname"})
	assert.NoError(t, err)

	info, err := client.GetUserInfo(ctx, &monify.GetUserInfoRequest{UserId: userId})
	assert.NoError(t, err)
	assert.Equal(t, "testname", info.Name)
}
