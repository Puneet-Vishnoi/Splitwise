package main

import (
	"fmt"

	"github.com/Puneet-Vishnoi/Splitwise/manager"
	"github.com/Puneet-Vishnoi/Splitwise/models"
)

func main() {

	em := manager.NewExpenseManager()

	u1 := &models.User{ID: "u1", Name: "user1"}
	u2 := &models.User{ID: "u2", Name: "user2"}
	u3 := &models.User{ID: "u3", Name: "user3"}
	u4 := &models.User{ID: "u4", Name: "user4"}

	em.AddUser(u1)
	em.AddUser(u2)
	em.AddUser(u3)
	em.AddUser(u4)

	m1 := models.NewGroup("g1", "group1", []*models.User{u1, u2, u3, u4})
	em.AddGroup(m1)

	//Equal

	ep, err := models.NewEqaulExpense(map[string]interface{}{
		"ID":        "e1",
		"Des":       "this is equaly devided",
		"Amount":    300.00,
		"PaidBy":    u1,
		"SplitType": models.Equal,
		"Users":     []*models.User{u1, u2, u4},
	})
	if err == nil{
		em.AddExpense(m1.Id, ep)
		em.ShowBalance(m1.Id)
	}




	m2 := models.NewGroup("g2", "group2", []*models.User{u1, u3, u4})
	em.AddGroup(m2)

	//Exact

	ep, err = models.NewExactExpense(map[string]interface{}{
		"ID":        "e2",
		"Des":       "this is exact devided",
		"Amount":    300.00,
		"PaidBy":    u3,
		"SplitType": models.Exact,
		"ExactSplits": []*models.ExactSplit{
			models.NewExactSplit(55, u3),
			models.NewExactSplit(125.86, u4),
			models.NewExactSplit(118.14, u1),
		},
	})
	fmt.Println(ep)
	em.AddExpense(m2.Id, ep)
	em.ShowBalance(m2.Id)


	m3 := models.NewGroup("g3", "group3", []*models.User{u3, u4})
	em.AddGroup(m3)

	// Percentage

	ep, err = models.NewPercentageExpense(map[string]interface{}{
		"ID":        "e3",
		"Des":       "this is percentage devided",
		"Amount":    300.00,
		"PaidBy":    u4,
		"SplitType": models.Percetage,
		"PercentageSplits": []*models.PercentageSplit{
			models.NewPercentageSplit(74.14, u3),
			models.NewPercentageSplit(25.86, u4),
		},
	})
	if err == nil{
		fmt.Println(ep)
		em.AddExpense(m3.Id, ep)
		em.ShowBalance(m3.Id)
	}
	


	m4 := models.NewGroup("g4", "group4", []*models.User{u2, u4})
	em.AddGroup(m4)

	// Share
	ep, err = models.NewShareExpense(map[string]interface{}{
		"ID":        "e4",
		"Des":       "this is share devided",
		"Amount":    300.00,
		"PaidBy":    u3,
		"SplitType": models.Share,
		"ShareSplits": []*models.ShareSplit{
			models.NewShareSplit(7, u3),
			models.NewShareSplit(5, u4),
			models.NewShareSplit(10, u1),
		},
	})
	if err == nil{
		fmt.Println(ep)
		em.AddExpense(m4.Id, ep)
		em.ShowBalance(m4.Id)
	}



	m5 := models.NewGroup("g5", "group5", []*models.User{u1, u2, u3, u4})
	em.AddGroup(m5)

	//Equal

	ep , err = models.NewEqaulExpense(map[string]interface{}{
		"ID":        "e5",
		"Des":       "this is equaly devided",
		"Amount":    300.00,
		"PaidBy":    u4,
		"SplitType": models.Equal,
		"Users":     []*models.User{u1, u2, u3, u4},
	})
	if err == nil{
		fmt.Println(ep)
		em.AddExpense(m5.Id, ep)
		em.ShowBalance(m5.Id)
	}
}
