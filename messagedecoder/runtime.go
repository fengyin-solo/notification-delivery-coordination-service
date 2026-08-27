package messagedecoder

// Decoder is stateless: every Decode call returns an independent slice, so
// results never alias one another or any internal buffer.
type Decoder struct{}

func NewDecoder() *Decoder { return &Decoder{} }

// Decode returns a freshly allocated slice holding value. The returned slice
// shares no state with the Decoder or with any prior Decode result, so callers
// may mutate it freely and later decodings will not retroactively alter
// earlier results (the equal-length payload aliasing bug).
func (d *Decoder) Decode(value string) []byte {
    out := make([]byte, len(value))
    copy(out, value)
    return out
}
