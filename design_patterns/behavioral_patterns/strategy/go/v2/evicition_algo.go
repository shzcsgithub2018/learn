package v2

import (
	"log"
)

func Lru(cache *Cache) {
	log.Println("Evicting by lru strategy")
}

func Lfu(cache *Cache) {
	log.Println("Evicting by lfu strategy")
}

func Fifo(cache *Cache) {
	log.Println("Evicting by fifo strategy")
}
