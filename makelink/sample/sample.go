//go:build run
// +build run

package main

import (
	"context"
	"fmt"
	"io"
	"os"

	"github.com/goark/ml/makelink"
)

func main() {
	lnk, err := makelink.New(context.Background(), "https://git.io/vFR5M")
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}
	_, _ = io.Copy(os.Stdout, lnk.Encode(makelink.StyleMarkdown))
	// Output:
	// [GitHub - goark/ml: Make Link with Markdown Format](https://github.com/goark/ml)
}
