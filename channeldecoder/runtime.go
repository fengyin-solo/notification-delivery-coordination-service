package channeldecoder

type Decoder struct { buffer []byte }

func NewDecoder() *Decoder { return &Decoder{buffer: make([]byte, 0, 64)} }

func (d *Decoder) Decode(value string) []byte {
    d.buffer = append(d.buffer[:0], value...)
    // 返回独立副本，避免后续 Decode 复用底层数组覆盖本次结果。
    out := make([]byte, len(d.buffer))
    copy(out, d.buffer)
    return out
}
