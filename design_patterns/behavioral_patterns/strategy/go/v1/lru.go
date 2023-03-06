package v1

import (
	"log"
)

type Lru struct {
}

func (l *Lru) evict(c *Cache) {
	log.Println("Evicting by lru strategy")
}
