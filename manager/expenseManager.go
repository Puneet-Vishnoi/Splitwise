package manager

import (
	"fmt"
	"sync"

	"github.com/Puneet-Vishnoi/Splitwise/models"
)

type ExpenseManager struct {
	Users        map[string]*models.User
	Group        map[string]*models.Group
	BalanceSheet map[string]map[string]float64
	mu           sync.RWMutex
}

var once sync.Once
var instance *ExpenseManager

func NewExpenseManager() *ExpenseManager {
	once.Do(func() {
		instance = &ExpenseManager{
			Users:        make(map[string]*models.User),
			Group:        make(map[string]*models.Group),
			BalanceSheet: make(map[string]map[string]float64),
		}
	})
	return instance
}

func (em *ExpenseManager) AddUser(u *models.User) {
	em.mu.Lock()
	defer em.mu.Unlock()
	em.Users[u.ID] = u
}

func (em *ExpenseManager) AddGroup(g *models.Group) {
	em.mu.Lock()
	defer em.mu.Unlock()
	em.Group[g.Id] = g
}

func (em *ExpenseManager) AddExpense(groupId string, expense *models.Expense) {
	em.mu.Lock()
	defer em.mu.Unlock()
	group, ok := em.Group[groupId]
	if !ok {
		fmt.Println("Please Enter valid GroupID")
		return
	}
	group.Expenses = append(group.Expenses, expense)

	payer := expense.PaidBy.ID
	for _, s := range expense.Splits {
		amount := s.GetAmount()
		userId := s.GetUser().ID

		if payer == userId {
			continue
		}
		if em.BalanceSheet[userId] == nil {
			em.BalanceSheet[userId] = make(map[string]float64)
		}
		if em.BalanceSheet[payer] == nil {
			em.BalanceSheet[payer] = make(map[string]float64)
		}

		em.BalanceSheet[userId][payer] += amount
		em.BalanceSheet[payer][userId] -= amount
	}
}

func (em *ExpenseManager) ShowBalance(groupId string) {
	em.mu.RLock()
	defer em.mu.RUnlock()
	group, ok := em.Group[groupId]
	if !ok {
		fmt.Println("Please Enter valid GroupID")
		return
	}
	for _, m1 := range group.Members {
		for _, m2 := range group.Members {
			amount := em.BalanceSheet[m1.ID][m2.ID]
			if amount > 0 {
				fmt.Printf("%s owes %.2f to %s\n", m1.Name, amount, m2.Name)
			}
		}
	}
}