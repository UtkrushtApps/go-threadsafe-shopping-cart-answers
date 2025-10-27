package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"

	"go-threadsafe-shopping-cart/cart"
)

func simulateCustomer(id int, cart *cart.Cart, wg *sync.WaitGroup, items []string) {
	defer wg.Done()
	r := rand.New(rand.NewSource(time.Now().UnixNano() + int64(id)))
	for i := 0; i < 10; i++ {
		action := r.Intn(3)
		item := items[r.Intn(len(items))]
		price := float64(10 + r.Intn(90)) // Price between 10 and 99
		switch action {
		case 0:
			cart.AddItem(item, price)
			fmt.Printf("Customer %d added %s (%.2f)\n", id, item, price)
		case 1:
			cart.RemoveItem(item)
			fmt.Printf("Customer %d removed %s\n", id, item)
		case 2:
			// Just view total / count
			total := cart.Total()
			count := cart.Count()
			fmt.Printf("Customer %d checked cart: %d items, total $%.2f\n", id, count, total)
		}
		time.Sleep(time.Millisecond * time.Duration(10+r.Intn(40)))
	}
}

func main() {
	sharedCart := cart.NewCart()
	itemCatalog := []string{"apple", "banana", "orange", "pear", "mango", "grape", "carrot", "melon", "plum"}

	var wg sync.WaitGroup

	customerCount := 8
	for i := 0; i < customerCount; i++ {
		wg.Add(1)
		go simulateCustomer(i, sharedCart, &wg, itemCatalog)
	}

	wg.Wait()

	fmt.Println("\nFINAL CART STATE:")
	items := sharedCart.Snapshot()
	for item, price := range items {
		fmt.Printf("%s: $%.2f\n", item, price)
	}

	total := sharedCart.Total()
	count := sharedCart.Count()
	fmt.Printf("TOTAL: %d items, $%.2f\n", count, total)
}
