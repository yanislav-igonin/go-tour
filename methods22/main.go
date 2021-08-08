// Exercise: Readers
// Implement a Reader type that emits an infinite stream of the ASCII character 'A'.

package main

import (
	"golang.org/x/tour/reader"
)

type MyReader struct{}

func (r MyReader) Read(p []byte) (int, error) {
	if c := len(p); c > 0 {
		for i := 0; i < c; i++ {
			p[i] = 65
		}
		return c, nil
	} else {
		return 0, nil
	}
}

func main() {
	reader.Validate(MyReader{})
}
