package topiccatalog
type Cache struct { values map[string][]byte }
func NewCache() *Cache { return &Cache{values: make(map[string][]byte)} }
func (c *Cache) Put(key string, value []byte) { c.values[key] = value }
func (c *Cache) Get(key string) []byte { return c.values[key] }
