package constant

type (
	LoanStatus int
	LoanAction int
)

const (
	StatusProposed  LoanStatus = 1
	StatusApproved  LoanStatus = 2
	StatusInvested  LoanStatus = 3
	StatusDisbursed LoanStatus = 4

	ActionApprove  LoanAction = 1
	ActionInvest   LoanAction = 2
	ActionDisburse LoanAction = 3
)

func (s LoanStatus) Int() int {
	return int(s)
}

func (a LoanAction) Int() int {
	return int(a)
}
