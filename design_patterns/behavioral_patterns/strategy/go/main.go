package main

import (
	v1 "github.com/shzgithub2018/learn/design_patterns/behavioral_patterns/strategy/go/v1"
	v2 "github.com/shzgithub2018/learn/design_patterns/behavioral_patterns/strategy/go/v2"
)

func main() {
	v1Test()
	/*Output
	Evicting by lfu strategy
	Evicting by lru strategy
	Evicting by fifo strategy
	*/

	v2Test()
	/*Output
	Evicting by lfu strategy
	Evicting by lru strategy
	Evicting by fifo strategy
	*/
}

func v1Test() {
	lfu := &v1.Lfu{}
	cache := v1.InitCache(lfu)

	cache.Add("a", "1")
	cache.Add("b", "2")

	cache.Add("c", "3")

	lru := &v1.Lru{}
	cache.SetEvictionAlgo(lru)

	cache.Add("d", "4")

	fifo := &v1.Fifo{}
	cache.SetEvictionAlgo(fifo)

	cache.Add("e", "5")
}

func v2Test() {
	lfu := v2.Lfu
	cache := v2.InitCache(lfu)

	cache.Add("a", "1")
	cache.Add("b", "2")

	cache.Add("c", "3")

	lru := v2.Lru
	cache.SetEvictionAlgo(lru)

	cache.Add("d", "4")

	fifo := v2.Fifo
	cache.SetEvictionAlgo(fifo)

	cache.Add("e", "5")
}
