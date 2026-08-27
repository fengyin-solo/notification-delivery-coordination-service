package channelcatalog

type Cache struct { values map[string][]byte }

func NewCache() *Cache { return &Cache{values: make(map[string][]byte)} }

// Put 写入渠道参数快照。对入参做深拷贝，保证后续对源切片的改动不会改写已缓存的内容。
func (c *Cache) Put(key string, value []byte) {
    snapshot := make([]byte, len(value))
    copy(snapshot, value)
    c.values[key] = snapshot
}

// Get 读取渠道参数快照。返回独立副本，保证调用方无法改写缓存内的内容。
func (c *Cache) Get(key string) []byte {
    v, ok := c.values[key]
    if !ok {
        return nil
    }
    out := make([]byte, len(v))
    copy(out, v)
    return out
}
