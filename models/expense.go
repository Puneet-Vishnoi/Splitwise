package models

import (
	"errors"
	"fmt"
)

type Expense struct {
	Id          string
	Description string
	PaidBy      *User
	SplitType   SplitType
	Splits      []Split
	Amount      float64
}

func NewEqaulExpense(mp map[string]interface{}) (*Expense, error) {
	id, _ := mp["ID"].(string)

	des, ok := mp["Des"].(string)
	if !ok || des == "" {
		des = "No description provided"
	}

	amount, amountOk := mp["Amount"].(float64)
	paidby, paidbyOk := mp["PaidBy"].(*User)
	splitType, splitTypeOk := mp["SplitType"].(SplitType)
	users, usersOk := mp["Users"].([]*User)

	if !amountOk || !paidbyOk || !splitTypeOk || !usersOk || len(users) == 0 {
		fmt.Println("Invalid or missing fields in expense creation")
		return nil, errors.New("Invalid or missing fields in percentage split expense")
	}

	var s []Split
	s = make([]Split, 0)

	for _, u := range users {
		each := amount / float64(len(users))
		s = append(s, NewEqualSplit(each, u))
	}
	return &Expense{
		Id:          id,
		Description: des,
		Amount:      amount,
		PaidBy:      paidby,
		SplitType:   splitType,
		Splits:      s,
	}, nil
}

func NewExactExpense(mp map[string]interface{}) (*Expense, error) {
	id, _ := mp["ID"].(string)

	des, ok := mp["Des"].(string)
	if !ok || des == "" {
		des = "No description provided"
	}

	amount, amountOk := mp["Amount"].(float64)
	paidby, paidbyOk := mp["PaidBy"].(*User)
	splitType, splitTypeOk := mp["SplitType"].(SplitType)
	exactSplits, exactexactSplitsOK := mp["ExactSplits"].([]*ExactSplit)

	if !amountOk || !paidbyOk || !splitTypeOk || !exactexactSplitsOK || len(exactSplits) == 0 {
		fmt.Println("Invalid or missing fields in expense creation")
		return nil, errors.New("Invalid or missing fields in expense creation")
	}

	var s []Split
	s = make([]Split, len(exactSplits))

	for i, es := range exactSplits {
		s[i] = es
	}
	return &Expense{
		Id:          id,
		Description: des,
		Amount:      amount,
		PaidBy:      paidby,
		SplitType:   splitType,
		Splits:      s,
	}, nil
}
func NewPercentageExpense(mp map[string]interface{}) (*Expense, error) {
	id, _ := mp["ID"].(string)

	des, ok := mp["Des"].(string)
	if !ok || des == "" {
		des = "No description provided"
	}

	amount, amountOk := mp["Amount"].(float64)
	paidBy, paidByOk := mp["PaidBy"].(*User)
	splitType, splitTypeOk := mp["SplitType"].(SplitType)
	percentSplits, splitsOk := mp["PercentageSplits"].([]*PercentageSplit)

	if !amountOk || !paidByOk || !splitTypeOk || !splitsOk || len(percentSplits) == 0 {
		fmt.Println("Invalid or missing fields in percentage split expense")
		return nil, errors.New("Invalid or missing fields in percentage split expense")
	}

	var s []Split
	for _, ps := range percentSplits {
		ps.Amount = ps.Percetage * amount / 100
		s = append(s, ps)
	}

	return &Expense{
		Id:          id,
		Description: des,
		Amount:      amount,
		PaidBy:      paidBy,
		SplitType:   splitType,
		Splits:      s,
	}, nil
}

func NewShareExpense(mp map[string]interface{}) (*Expense, error) {
	id, _ := mp["ID"].(string)

	des, ok := mp["Des"].(string)
	if !ok || des == "" {
		des = "No description provided"
	}

	amount, amountOk := mp["Amount"].(float64)
	paidBy, paidByOk := mp["PaidBy"].(*User)
	splitType, splitTypeOk := mp["SplitType"].(SplitType)
	shareSplits, splitsOk := mp["ShareSplits"].([]*ShareSplit)

	if !amountOk || !paidByOk || !splitTypeOk || !splitsOk || len(shareSplits) == 0 {
		fmt.Println("Invalid or missing fields in share split expense")
		return nil, errors.New("Invalid or missing fields in share split expense")
	}

	var totalShares int64
	for _, ss := range shareSplits {
		totalShares += ss.Share
	}

	if totalShares == 0 {
		fmt.Println("Total shares cannot be zero")
		return nil, errors.New("Total shares cannot be zero")
	}

	var s []Split
	for _, ss := range shareSplits {
		ss.Amount = float64(ss.Share) * amount / float64(totalShares)
		s = append(s, ss)
	}

	return &Expense{
		Id:          id,
		Description: des,
		Amount:      amount,
		PaidBy:      paidBy,
		SplitType:   splitType,
		Splits:      s,
	}, nil
}