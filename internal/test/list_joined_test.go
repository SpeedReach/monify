package test

import (
	"testing"
)

func TestListJoined(t *testing.T) {
	panic("panic")
	/*
		client := GetTestClient(t)
		userId1 := client.CreateTestUser()
		group, err := client.CreateGroup(context.TODO(), &monify.CreateGroupRequest{Name: "test"})
		assert.NoError(t, err)
		assert.NotEmpty(t, group.GroupId)
		userId2 := client.CreateTestUser()
		client.SetTestUser(userId1)
		_, err := client.ListJoinedGroups(context.TODO(), &monify.Empty{})
		assert.Error(t, err)
		assert.Equal()



		client := GetTestClient(t)
		userId1 := client.CreateTestUser()
		group, err := client.CreateGroup(context.TODO(), &monify.CreateGroupRequest{Name: "test"})
		assert.NoError(t, err)
		assert.NotEmpty(t, group.GroupId)
		userId2 := client.CreateTestUser()
		client.SetTestUser(userId1)

		_, err := client.EmailLogin(context.TODO(), &monify.EmailLoginRequest{Email: email})
		assert.Error(t, err)
		res1, err := client.EmailRegister(context.TODO(), &monify.EmailRegisterRequest{Email: email})
		assert.NoError(t, err)
		res2, err := client.EmailLogin(context.TODO(), &monify.EmailLoginRequest{Email: email})
		assert.NoError(t, err)
		assert.Equal(t, res1.UserId, res2.UserId)
	*/
}
