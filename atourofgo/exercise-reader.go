package main

import (
	"code.google.com/p/go-tour/reader"
)

type MyReader struct{}

// TODO: Add a Read([]byte) (int, error) method to MyReader.
func (MyReader) Read(bytes []byte) (int, error) {
	for x := range bytes {
		bytes[x] = 'A'
	}
  return len(bytes), nil
}

func main() {
	reader.Validate(MyReader{})
}
