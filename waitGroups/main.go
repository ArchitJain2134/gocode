// Add(n) → Increases the counter by n (number of goroutines).
// Done() → Decreases the counter when a goroutine finishes.
// Wait() → Blocks execution until the counter becomes 0
package main

import (
	"fmt"
	"sync"
)

func a(wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("a wala func chl rha h")
}

func main() {
	wg := sync.WaitGroup{}
	wg.Add(1)
	go a(&wg)
	wg.Add(1)
	go a(&wg)
	fmt.Println("**********")
	wg.Wait()
	fmt.Println("----------")

}
