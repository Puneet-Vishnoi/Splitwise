package models

type SplitType string

const (
	Equal     SplitType = "EqualSplit"
	Exact     SplitType = "ExactSplit"
	Percetage SplitType = "PercentageSplit"
	Share     SplitType = "ShareSplit"
)

type Split interface {
	GetUser() *User
	GetAmount() float64
}

type EqualSplit struct {
	Amount float64
	User   *User
}

func NewEqualSplit(amount float64, user *User) *EqualSplit {
	return &EqualSplit{
		Amount: amount,
		User:   user,
	}
}

func (s *EqualSplit) GetUser() *User {
	return s.User
}

func (s *EqualSplit) GetAmount() float64 {
	return s.Amount
}

type ExactSplit struct {
	Amount float64
	User   *User
}

func NewExactSplit(amount float64, user *User) *ExactSplit {
	return &ExactSplit{
		Amount: amount,
		User:   user,
	}
}

func (s *ExactSplit) GetUser() *User {
	return s.User
}

func (s *ExactSplit) GetAmount() float64 {
	return s.Amount
}

type PercentageSplit struct {
	Amount    float64
	User      *User
	Percetage float64
}

func NewPercentageSplit(percentage float64, user *User) *PercentageSplit {
	return &PercentageSplit{
		Percetage: percentage,
		User:      user,
	}
}

func (s *PercentageSplit) GetAmount() float64 {
	return s.Amount
}

func (s *PercentageSplit) GetUser() *User {
	return s.User
}

type ShareSplit struct {
	Amount float64
	User   *User
	Share  int64
}

func NewShareSplit(share int64, user *User) *ShareSplit {
	return &ShareSplit{
		Share: share,
		User:  user,
	}
}

func (s *ShareSplit) GetAmount() float64 {
	return s.Amount
}

func (s *ShareSplit) GetUser() *User {
	return s.User
}
