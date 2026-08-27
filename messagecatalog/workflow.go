package messagecatalog

// Cache stores independent copies of every payload. Neither Put nor Get
// shares mutable backing storage with the caller: the value Put is copied in,
// and the slice Get returns is a fresh copy, so caller mutations never bleed
// back into the cache and never alias another cached entry.
type Cache struct {
    values map[string][]byte
}

func NewCache() *Cache { return &Cache{values: make(map[string][]byte)} }

// Put stores a private copy of value. Mutating value after Put has no effect
// on the cached entry, and two Puts that share backing arrays cannot alias.
func (c *Cache) Put(key string, value []byte) {
    out := make([]byte, len(value))
    copy(out, value)
    c.values[key] = out
}

// Get returns a fresh copy of the cached value, or nil. Mutating the returned
// slice never affects the cache or any other Get result.
func (c *Cache) Get(key string) []byte {
    v, ok := c.values[key]
    if !ok {
        return nil
    }
    out := make([]byte, len(v))
    copy(out, v)
    return out
}
