package v1

import (
	"log"
)

type Fifo struct {
}

func (l *Fifo) evict(c *Cache) {
	log.Println("Evicting by fifo strategy")
}
