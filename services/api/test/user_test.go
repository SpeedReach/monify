package test

import (
	"context"
	monify "monify/protobuf/gen/go"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestChangeUsernameAndUsernickId(t *testing.T) {
	client := GetTestClient(t)
	userId := client.CreateTestUser()
	ctx := context.Background()

	_, err := client.UpdateUserName(ctx, &monify.UpdateUserNameRequest{Name: "testname"})
	assert.NoError(t, err)

	_, err = client.UpdateNickId(ctx, &monify.UpdateNickIdRequest{NickId: "test_nickId"})
	assert.NoError(t, err)

	info, err := client.GetUserInfo(ctx, &monify.GetUserInfoRequest{UserId: userId})
	assert.NoError(t, err)
	assert.Equal(t, "testname", info.Name)
	assert.Equal(t, "test_nickId", info.NickId)
}
