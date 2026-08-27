package templatecatalog

// Cache 缓存模板变量字节切片。Put 时持有一份独立副本，Get 时再返回一份副本，
// 使缓存内部存储与查询结果互不影响，避免查询结果被反向修改而改坏缓存。
type Cache struct{ values map[string][]byte }

func NewCache() *Cache { return &Cache{values: make(map[string][]byte)} }

// Put 存入变量值的一份独立副本，后续对入参切片的修改不会污染缓存。
func (c *Cache) Put(key string, value []byte) {
	dup := make([]byte, len(value))
	copy(dup, value)
	c.values[key] = dup
}

// Get 返回变量值的一份独立副本，对返回切片的修改不会反向写回缓存。
func (c *Cache) Get(key string) []byte {
	v, ok := c.values[key]
	if !ok {
		return nil
	}
	out := make([]byte, len(v))
	copy(out, v)
	return out
}
