# Solution Steps

1. Create a new Go package named 'cart' and in cart.go, define a Cart struct with a map for items and a sync.RWMutex.

2. Implement the AddItem method: lock the cart for writing, add or update the item in the map, and unlock.

3. Implement the RemoveItem method: lock the cart for writing, delete the item from the map, and unlock.

4. Implement the Total method: lock the cart for reading, sum the prices of all items, and unlock.

5. Implement the Count method: lock the cart for reading, return the length of the items map, and unlock.

6. Implement a Snapshot method in Cart to provide a copy of the map for safe external reading.

7. In main.go, import the cart package. Create the shared cart and a list of sample item names.

8. Write a simulateCustomer function that randomly adds, removes, or views items in the cart, simulating concurrent access.

9. Launch multiple goroutines running simulateCustomer, each with a unique id and WaitGroup for concurrency management.

10. After all goroutines finish, display the final state of the cart and validate it by calling Cart methods from main.

