package topiccatalog

// Cache 缓存解码后的主题标识。所有进出缓存的字节切片都会被复制，
// 确保每个主题标识互不共享底层内存：调用方修改返回值不会污染缓存，
// 写入后修改源切片也不会影响缓存内部状态。
type Cache struct{ values map[string][]byte }

func NewCache() *Cache { return &Cache{values: make(map[string][]byte)} }

// Put 写入主题标识。复制传入切片，避免外部后续修改影响缓存内部状态。
func (c *Cache) Put(key string, value []byte) {
	cp := make([]byte, len(value))
	copy(cp, value)
	c.values[key] = cp
}

// Get 取出主题标识。返回缓存内部的副本，调用方修改返回值不会影响缓存。
// 缺失键返回 nil。
func (c *Cache) Get(key string) []byte {
	v := c.values[key]
	if v == nil {
		return nil
	}
	out := make([]byte, len(v))
	copy(out, v)
	return out
}
