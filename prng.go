package dprng

import "io"

func NewRandReader(seed int64) io.Reader {
	if seed < 0 {
		seed = -seed
	}
	return &prngReader{s: int(seed % seedSize)}
}

// prngReader implements an io.Reader that returns pseudorandom bytes.
// ensure every call with will return same random pattern
type prngReader struct{ s int }

// Read implements io.Reader.Read.
func (r *prngReader) Read(p []byte) (int, error) {
	idx := r.s
	var copied int
	for copied < len(p) {
		copied += copy(p[copied:], seedData[idx:])
		idx += copied
		idx = idx % seedSize
	}
	return copied, nil
}
