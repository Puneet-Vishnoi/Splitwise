package models

type Group struct {
	Id       string
	Members  []*User
	Name     string
	Expenses []*Expense
}

func NewGroup(Id, Name string, members []*User) *Group {
	return &Group{
		Id:       Id,
		Name:     Name,
		Members:  members,
		Expenses: make([]*Expense, 0),
	}
}
