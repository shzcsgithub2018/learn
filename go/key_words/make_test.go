package key_words

import (
	"sync"
	"testing"
)

func Test_slice(t *testing.T) {
	//s1 := make([]int64)   必须指定 size

	s2 := make([]int64, 10)
	for i := int64(0); i < 10; i++ {
		s2[i] = i
	}
	t.Log(s2)
	// [0 1 2 3 4 5 6 7 8 9]
}

func Test_map(t *testing.T) {
	m1 := make(map[int]string)
	m2 := make(map[int]string, 10)

	for i := 0; i < 10; i++ {
		m1[i] = string(rune(i) + 'a')
		m2[i] = string(rune(i) + 'b')
	}

	t.Log(m1, m2)
	// map[0:a 1:b 2:c 3:d 4:e 5:f 6:g 7:h 8:i 9:j]
	// map[0:b 1:c 2:d 3:e 4:f 5:g 6:h 7:i 8:j 9:k]
}

func Test_chan(t *testing.T) {
	// 同步阻塞
	ch1 := make(chan int)
	ch2 := make(chan int, 0)

	// 缓冲区未满 异步  满了 阻塞
	ch3 := make(chan int, 10)

	wg := sync.WaitGroup{}
	wg.Add(1)
	go func() {
		defer wg.Done()
		for i := 0; i < 10; i++ {
			t.Log("ch1 in ", i)
			ch1 <- i
		}
		close(ch1)
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		for num := range ch1 {
			t.Log("ch1 out ", num)
		}
	}()
	wg.Wait()
	t.Log("-----------------------------------")

	wg = sync.WaitGroup{}
	wg.Add(1)
	go func() {
		defer wg.Done()
		for i := 0; i < 10; i++ {
			t.Log("ch2 in ", i)
			ch2 <- i
		}
		close(ch2)
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		for num := range ch2 {
			t.Log("ch2 out ", num)
		}
	}()
	wg.Wait()
	t.Log("-----------------------------------")

	wg = sync.WaitGroup{}
	wg.Add(1)
	go func() {
		defer wg.Done()
		for i := 0; i < 10; i++ {
			t.Log("ch3 in ", i)
			ch3 <- i
		}
		close(ch3)
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		for num := range ch3 {
			t.Log("ch3 out ", num)
		}
	}()
	wg.Wait()
}
