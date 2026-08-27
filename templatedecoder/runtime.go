package templatedecoder

// Decoder 将模板变量值解码为字节切片。buffer 仅作复用的临时缓冲区，
// Decode 返回的切片必须与后续调用互相独立，避免等长变量连续解码时
// 因复用同一底层数组而把首组内容覆盖成后组内容。
type Decoder struct{ buffer []byte }

func NewDecoder() *Decoder { return &Decoder{buffer: make([]byte, 0, 64)} }

// Decode 解码变量值，返回一份独立副本，调用方修改不会影响后续解码结果。
func (d *Decoder) Decode(value string) []byte {
	d.buffer = append(d.buffer[:0], value...)
	out := make([]byte, len(d.buffer))
	copy(out, d.buffer)
	return out
}
