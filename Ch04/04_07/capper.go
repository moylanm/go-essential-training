package main

import (
	"fmt"
	"io"
	"os"
)

// Capper implements io.Writer and turns everything to uppercase
type Capper struct {
	wtr io.Writer
}

func (c *Capper) Write(data []byte) (int, error) {
	diff := byte('a' - 'A')
	out := make([]byte, len(data))

	for i, c := range data {
		if c >= 'a' && c <= 'z' {
			c -= diff
		}
		out[i] = c
	}

	return c.wtr.Write(out)
}

func main() {
	c := &Capper{wtr: os.Stdout}
	fmt.Fprintln(c, "Hello there")
}
