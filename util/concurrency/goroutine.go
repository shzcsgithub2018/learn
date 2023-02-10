package concurrency

import (
	"sync"
)

type Operator func()

// UnSafe
// @Description: x^y goroutine execute op
// @param x
// @param y
// @param op
func UnSafe(x, y int, op Operator) {
	wg := sync.WaitGroup{}
	wg.Add(1)
	go func() {
		defer wg.Done()
		for x > 0 {
			x--
			wg.Add(1)
			go func(y int) {
				defer wg.Done()
				for y > 0 {
					y--
					op()
				}
			}(y)
		}
	}()
	wg.Wait()
}
