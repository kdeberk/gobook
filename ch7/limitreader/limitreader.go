package main

import "io"

type reader struct {
	r    io.Reader
	left int64
}

func LimitReader(r io.Reader, n int64) io.Reader {
	return &reader{r: r, left: n}
}

func (r *reader) Read(p []byte) (n int, err error) {
	switch {
	case r.left <= 0:
		err = io.EOF
	case r.left < int64(len(p)):
		n, err = r.r.Read(p[:r.left])
		r.left = 0
	default:
		n, err = r.r.Read(p)
		r.left -= int64(len(p))
	}
	return
}
