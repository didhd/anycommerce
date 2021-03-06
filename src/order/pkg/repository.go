package pkg

import (
	"strconv"
)

var currentID int

var orders Orders = Orders{}

// Init
func init() {
}

// RepoFindOrderByID Function
func RepoFindOrderByID(id string) Order {
	for _, t := range orders {
		if t.ID == id {
			return t
		}
	}
	// return empty Order if not found
	return Order{}
}

// RepoFindOrdersByUsername Function
func RepoFindOrdersByUsername(username string) Orders {
	var o Orders = Orders{}

	for _, t := range orders {
		if t.Username == username {
			o = append(o, t)
		}
	}

	return o
}

func RepoUpdateOrder(t Order) Order {
	for i := 0; i < len(orders); i++ {
		o := &orders[i]
		if o.ID == t.ID {
			orders[i] = t
			return RepoFindOrderByID(t.ID)
		}
	}

	// return empty Order if not found
	return Order{}
}

// RepoCreateOrder Function
func RepoCreateOrder(t Order) Order {
	currentID++
	t.ID = strconv.Itoa(currentID)
	orders = append(orders, t)
	return t
}
