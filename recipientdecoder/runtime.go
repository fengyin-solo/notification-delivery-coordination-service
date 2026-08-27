package recipientdecoder

type Decoder struct { buffer []byte }

func NewDecoder() *Decoder { return &Decoder{buffer: make([]byte, 0, 64)} }

func (d *Decoder) Decode(value string) []byte {
    d.buffer = append(d.buffer[:0], value...)
    // 返回独立副本，避免复用内部 buffer 导致连续解析两批地址时
    // 首批内容被后批覆盖（同一底层数组别名）。
    out := make([]byte, len(d.buffer))
    copy(out, d.buffer)
    return out
}
