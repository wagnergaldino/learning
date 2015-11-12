package main

import (
	"io"
	"os"
	"strings"
)

type rot13Reader struct {
	r io.Reader
}

func (rot *rot13Reader) Read(bytes []byte) (n int, err error) {
	n, err = rot.r.Read(bytes)
	for x := range bytes {
		if (bytes[x] >= 'A' && bytes[x] < 'N') || (bytes[x] >='a' && bytes[x] < 'n') {
			bytes[x] += 13
		} else if (bytes[x] > 'M' && bytes[x] <= 'Z') || (bytes[x] > 'm' && bytes[x] <= 'z') {
			bytes[x] -= 13
		}
	}
	return
}

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}
