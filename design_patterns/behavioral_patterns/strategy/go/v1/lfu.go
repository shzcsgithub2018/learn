package v1

import (
	"log"
)

type Lfu struct {
}

func (l *Lfu) evict(c *Cache) {
	log.Println("Evicting by lfu strategy")
}
