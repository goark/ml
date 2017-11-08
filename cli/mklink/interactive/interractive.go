package interactive

import (
	"bufio"
	"fmt"
	"io"
	"os"

	"github.com/spiegel-im-spiegel/mklink"
	"github.com/spiegel-im-spiegel/mklink/cli/mklink/makelink"
)

const (
	header = "Press Ctrl+C to stop"
	prompt = "mklink> "
)

//Context class is context for making link
type Context struct {
	*makelink.Context
}

//New returns new Context instance
func New(s mklink.Style, log io.Writer) *Context {
	return &Context{makelink.New(s, os.Stdout, log)}
}

//Run is running interactive mode
func (c *Context) Run() {
	reader := os.Stdin
	writer := os.Stdout
	fmt.Fprintln(writer, header)
	fmt.Fprint(writer, prompt)

	scanner := bufio.NewScanner(reader)
	for scanner.Scan() {
		err := c.MakeLink(scanner.Text())
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
		}
		fmt.Fprint(writer, prompt)
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}
}
