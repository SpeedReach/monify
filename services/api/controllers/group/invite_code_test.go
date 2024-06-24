package group

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestGenerateInviteCodeUniqueness(t *testing.T) {
	inviteCode := generateInviteCode()
	assert.Len(t, inviteCode, inviteCodeLength, "Expected invite code to be 6 characters long")
	time.Sleep(time.Millisecond)
	inviteCode2 := generateInviteCode()
	assert.NotEqual(t, inviteCode, inviteCode2, "Expected two different invite codes")
}

func TestGenerateInviteCodeRandomness(t *testing.T) {
	//try to generate 1000 invite codes and check if 990 of them are unique
	failCount := 0
	tries := 1000
	for i := 0; i < tries; i++ {
		if generateInviteCode() == generateInviteCode() {
			failCount++
		}
	}
	assert.LessOrEqual(t, failCount, 10, "Expected 990 out of 1000 invite codes to be unique")
}
