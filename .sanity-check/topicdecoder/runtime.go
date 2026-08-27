package topicdecoder
type Decoder struct { buffer []byte }
func NewDecoder() *Decoder { return &Decoder{buffer: make([]byte, 0, 64)} }
func (d *Decoder) Decode(value string) []byte { d.buffer = append(d.buffer[:0], value...); return d.buffer }
