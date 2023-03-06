package v1

type EvictionAlgo interface {
	evict(c *Cache)
}
