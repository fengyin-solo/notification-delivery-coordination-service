package topicdecoder

// Decoder 将字符串主题标识解码为字节切片。内部复用 buffer 以摊销分配，
// 但每次 Decode 都返回独立副本，避免连续解码时后一次结果覆盖前一次。
type Decoder struct{ buffer []byte }

func NewDecoder() *Decoder { return &Decoder{buffer: make([]byte, 0, 64)} }

// Decode 解码主题标识。返回的字节切片与内部 buffer 及后续调用互不共享内存，
// 调用方修改返回值不会影响解码器内部状态。
func (d *Decoder) Decode(value string) []byte {
	d.buffer = append(d.buffer[:0], value...)
	out := make([]byte, len(d.buffer))
	copy(out, d.buffer)
	return out
}
