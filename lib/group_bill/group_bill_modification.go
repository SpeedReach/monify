package group_bill

import "github.com/google/uuid"

type ModificationType = int

const (
	Create ModificationType = 0
	Delete ModificationType = 1
	Update ModificationType = 2
)

type GroupBillModification struct {
	Ty               ModificationType
	GroupId          uuid.UUID
	OperatorMemberId uuid.UUID
	BillId           uuid.UUID
	Title            string
}
