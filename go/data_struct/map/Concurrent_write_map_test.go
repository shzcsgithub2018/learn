package map_test

import (
	"fmt"
	"sync"
	"testing"
)

func TestConcurrentWriteMap(t *testing.T) {
	m := make(map[int]string, 10)

	wg := sync.WaitGroup{}

	for i := 0; i <= 10000; i++ {
		wg.Add(1)
		go func(i int) {
			defer func() {
				wg.Done()
				if r := recover(); r != nil {
					fmt.Println("Recovered. Error:\n", r)
				}
			}()
			m[i] = "hello world" + fmt.Sprintf("%d", i)
		}(i)
	}
	wg.Wait()
}
