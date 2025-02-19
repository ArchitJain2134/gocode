package main

import (
	"fmt"
	"sync"
)

func main() {
	mutex := sync.Mutex{}
	wg:=sync.WaitGroup{}
	balance:=0

	for i:=1;i<=1000;i++{
		wg.Add(1)
		go paisabhadao(&balance , &wg, &mutex)		

	}
	wg.Wait()
	fmt.Println(balance)
}
	func paisabhadao (balance *int, wg *sync.WaitGroup, mutex *sync.Mutex) {
		defer wg.Done()
		mutex.Lock()
		// mutex mainly used in banking systems as above line is used to help in locking other transaction req. if any 
		// it also helps in preventing race conditions and provides atomicity
		*balance++
		mutex.Unlock()
	}

