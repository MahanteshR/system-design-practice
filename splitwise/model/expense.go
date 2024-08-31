package model

type Expense struct {
	ID          string
	Desc        string
	Amt         float64
	PayerID     string
	GroupID     string
	SplitMethod string
}
