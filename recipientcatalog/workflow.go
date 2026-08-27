package recipientcatalog

type Cache struct { values map[string][]byte }

func NewCache() *Cache { return &Cache{values: make(map[string][]byte)} }

// Put 写入时拷贝，避免调用方后续修改入参字节污染缓存。
func (c *Cache) Put(key string, value []byte) {
    stored := make([]byte, len(value))
    copy(stored, value)
    c.values[key] = stored
}

// Get 读取时拷贝，避免读取方修改返回字节污染缓存。
func (c *Cache) Get(key string) []byte {
    v, ok := c.values[key]
    if !ok {
        return nil
    }
    out := make([]byte, len(v))
    copy(out, v)
    return out
}
