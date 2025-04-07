package wallet

import (
	"fmt"
	"log"
	"sync"
)

/**
How Locks Work?

Read Lock (RLock):
- Multiple goroutines can acquire a read lock simultaneously.
- It allows multiple readers to access the resource concurrently.

Write Lock (Lock):
- Only one goroutine can acquire a write lock at a time.
- It prevents any other goroutine from further reading or writing until this lock is released.
*/

var walletMutex = sync.RWMutex{}
var userWallets = make(map[string]float32) // wallets Map [user ID => balance]

func GetWalletBalance(userID string) (float32, error) {
	walletMutex.RLock()         // Lock the mutex for reading
	defer walletMutex.RUnlock() // Unlock the mutex after the function completes, it will be unlocked even if an error or early return occurs

	balance, exists := userWallets[userID]
	if !exists {
		return 0, fmt.Errorf("wallet not found for user ID: %s", userID)
	} else {
		return balance, nil
	}
}

func CreateWallet(userID string) error {
	walletMutex.Lock()         // Lock the mutex to prevent concurrent access
	defer walletMutex.Unlock() // Unlock the mutex after the function completes, it will be unlocked even if an error or early return occurs

	_, exists := userWallets[userID]
	if exists {
		return fmt.Errorf("wallet already exists for user ID: %v", userID)
	} else {
		userWallets[userID] = 0.0
		log.Printf("New Wallet Created for %v :%v", userID, userWallets[userID])
		return nil
	}
}

func UpdateWalletBalance(userID string, balance float32) (float32, error) {
	walletMutex.Lock()         // Lock the mutex to prevent concurrent access
	defer walletMutex.Unlock() // Unlock the mutex after the function completes, it will be unlocked even if an error or early return occurs

	_, exists := userWallets[userID]
	if !exists {
		return 0, fmt.Errorf("wallet not found for user ID: %s", userID)
	} else {
		userWallets[userID] = balance
		return balance, nil
	}
}
