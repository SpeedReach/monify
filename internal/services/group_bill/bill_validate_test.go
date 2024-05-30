package group_bill

import (
	"github.com/stretchr/testify/assert"
	monify "monify/protobuf/gen/go"
	"testing"
)

func TestBillValidate(t *testing.T) {
	test1 := monify.CreateGroupBillRequest{
		GroupId:     "123",
		TotalMoney:  100,
		Title:       "test",
		Description: "test",
		SplitPeople: []*monify.SplitPerson{
			{
				MemberId: "123",
				Amount:   100,
			},
		},
	}

	err := validateGroupBill(&test1)
	assert.Error(t, err)
	test2 := monify.CreateGroupBillRequest{
		GroupId:     "123",
		TotalMoney:  100,
		Title:       "test",
		Description: "test",
		SplitPeople: []*monify.SplitPerson{
			{
				MemberId: "123",
				Amount:   50,
			},
			{
				MemberId: "123",
				Amount:   50,
			},
		},
		PrepaidPeople: []*monify.PrepaidPerson{
			{
				MemberId: "123",
				Amount:   100,
			},
		},
	}
	err = validateGroupBill(&test2)
	assert.NoError(t, err)
}
