package group

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestIndexToChar(t *testing.T) {
	assert.Equal(t, byte('0'), indexToChar(0), "Expected '0' for index 0")
	assert.Equal(t, byte('1'), indexToChar(1), "Expected '1' for index 1")
	assert.Equal(t, byte('2'), indexToChar(2), "Expected '2' for index 2")
	assert.Equal(t, byte('8'), indexToChar(8), "Expected '8' for index 8")
	assert.Equal(t, byte('a'), indexToChar(10), "Expected 'a' for index 10")
	assert.Equal(t, byte('b'), indexToChar(11), "Expected 'b' for index 11")
	assert.Equal(t, byte('z'), indexToChar(35), "Expected 'z' for index 35")
	assert.Equal(t, byte('A'), indexToChar(36), "Expected 'A' for index 36")
	assert.Equal(t, byte('B'), indexToChar(37), "Expected 'B' for index 37")
	assert.Equal(t, byte('Z'), indexToChar(61), "Expected 'Z' for index 61")
}

func TestGenerateInviteCodeUniqueness(t *testing.T) {
	inviteCode := generateInviteCode()
	assert.Len(t, inviteCode, inviteCodeLength, "Expected invite code to be 6 characters long")
	for _, char := range inviteCode {
		assert.True(t, char >= '0' && char <= '9' || char >= 'a' && char <= 'z' || char >= 'A' && char <= 'Z', "Expected invite code to contain only alphanumeric characters")
	}
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
