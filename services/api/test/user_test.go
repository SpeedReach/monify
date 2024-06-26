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

	_, err := client.UpdateUserName(context.Background(), &monify.UpdateUserNameRequest{Name: "testname"})
	assert.NoError(t, err)

	_, err = client.UpdateUserNickId(context.TODO(), &monify.UpdateUserNickIdRequest{NickId: "test_nickId"})
	assert.NoError(t, err)

	info, err := client.GetUserInfo(context.TODO(), &monify.GetUserInfoRequest{UserId: userId})
	assert.NoError(t, err)
	assert.Equal(t, "testname", info.Name)
	assert.Equal(t, "test_nickId", info.NickId)
}
