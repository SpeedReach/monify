package test

import (
	"context"
	"github.com/stretchr/testify/assert"
	monify "monify/protobuf"
	"testing"
	"time"
)

func TestLogin(t *testing.T) {
	time.Sleep(time.Second)
	client := GetTestClient()
	_, err := client.EmailLogin(context.TODO(), &monify.EmailLoginRequest{Email: "brian030128@gmail.com"})
	assert.Error(t, err)

}
