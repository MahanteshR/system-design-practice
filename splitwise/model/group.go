package model

type Group struct {
	GroupID   string
	Name      string
	MemberIDs []string
	Balances  map[string]float64
}
