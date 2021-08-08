// Exercise: rot13Reader
// A common pattern is an io.Reader that wraps another io.Reader,
// modifying the stream in some way.

// For example, the gzip.NewReader function takes an io.Reader
// (a stream of compressed data) and returns a *gzip.Reader
// that also implements io.Reader (a stream of the decompressed data).

// Implement a rot13Reader that implements io.Reader and reads
// from an io.Reader, modifying the stream by applying the rot13
// substitution cipher to all alphabetical characters.

// The rot13Reader type is provided for you.
// Make it an io.Reader by implementing its Read method.

package main

import (
	"io"
	"os"
	"strings"
)

type rot13Reader struct {
	r io.Reader
}

func (r *rot13Reader) Read(p []byte) (n int, err error) {
	b := make([]byte, 1024)

	n, err = r.r.Read(b)
	if err != nil {
		return 0, err
	}
	l := len(b)

	for i := 0; i < l; i++ {
		char := b[i]
		if (char >= 65 && char <= 77) || char >= 97 && char <= 109 {
			char += 13
		} else if (char >= 78 && char <= 90) || char >= 110 && char <= 122 {
			char -= 13
		}

		p[i] = char
	}

	return n, nil
}

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}
