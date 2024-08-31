package model

import (
	"github.com/google/uuid"
)

type SplitwiseSystem struct {
	Users    map[string]User
	Groups   map[string]Group
	Expenses []Expense
}

func (s *SplitwiseSystem) AddExpense(desc, payerID, groupID string, amt float64) Expense {
	id := uuid.NewString()

	expense := &Expense{
		ID:          id,
		Desc:        desc,
		PayerID:     payerID,
		GroupID:     groupID,
		Amt:         amt,
		SplitMethod: "EQUAL",
	}

	s.Expenses = append(s.Expenses, *expense)
	s.calculateBalance(groupID, id)

	return *expense
}

func (s *SplitwiseSystem) calculateBalance(groupID, expenseID string) {
	group, exists := s.Groups[groupID]
	if !exists {
		return
	}

	for _, expense := range s.Expenses {
		if expense.GroupID != groupID && expense.ID != expenseID {
			continue
		}

		splitAmt := expense.Amt / float64(len(group.MemberIDs))

		for _, memberID := range group.MemberIDs {
			group.Balances[memberID] -= splitAmt

			if memberID == expense.PayerID {
				group.Balances[memberID] += splitAmt
			}
		}
	}
}
