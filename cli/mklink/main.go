package main

import (
	"os"

	"github.com/spiegel-im-spiegel/gocli/rwi"
	"github.com/spiegel-im-spiegel/mklink/cli/mklink/facade"
)

func main() {
	facade.Execute(
		rwi.New(
			rwi.Reader(os.Stdin),
			rwi.Writer(os.Stdout),
			rwi.ErrorWriter(os.Stderr),
		),
		os.Args[1:],
	).Exit()
}
