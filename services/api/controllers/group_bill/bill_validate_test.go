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
		SplitPeople: []*monify.InsertSplitPerson{
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
		SplitPeople: []*monify.InsertSplitPerson{
			{
				MemberId: "123",
				Amount:   50,
			},
			{
				MemberId: "123",
				Amount:   50,
			},
		},
		PrepaidPeople: []*monify.InsertPrepaidPerson{
			{
				MemberId: "123",
				Amount:   100,
			},
		},
	}
	err = validateGroupBill(&test2)
	assert.NoError(t, err)

	test3 := monify.CreateGroupBillRequest{
		GroupId:     "123",
		TotalMoney:  250,
		Title:       "test",
		Description: "test",
		SplitPeople: []*monify.InsertSplitPerson{
			{
				MemberId: "",
				Amount:   100,
			},
			{
				MemberId: "",
				Amount:   150,
			},
		},
		PrepaidPeople: []*monify.InsertPrepaidPerson{
			{
				MemberId: "",
				Amount:   250,
			},
		},
	}

	err = validateGroupBill(&test3)
	assert.NoError(t, err)
}
