package interactive

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"

	isatty "github.com/mattn/go-isatty"
	"github.com/pkg/errors"
	"github.com/spiegel-im-spiegel/mklink"
	"github.com/spiegel-im-spiegel/mklink/cli/mklink/makelink"
)

const (
	headerStr = "Input 'q' or 'quit' to stop"
	promptStr = "mklink> "
)

//Context class is context for making link
type Context struct {
	*makelink.Context
	scanner *bufio.Scanner
	writer  io.Writer
}

//New returns new Context instance
func New(s mklink.Style, log io.Writer) (*Context, error) {
	if !isTerminal() {
		return nil, errors.New("not terminal (or pipe?)")
	}
	cxt := &Context{Context: makelink.New(s, os.Stdout, log), scanner: bufio.NewScanner(os.Stdin), writer: os.Stdout}
	cxt.EnableClipboard()
	return cxt, nil
}

func isTerminal() bool {
	if !isatty.IsTerminal(os.Stdin.Fd()) && !isatty.IsCygwinTerminal(os.Stdin.Fd()) {
		return false
	}
	if !isatty.IsTerminal(os.Stdout.Fd()) && !isatty.IsCygwinTerminal(os.Stdout.Fd()) {
		return false
	}
	return true
}

//Run is running interactive mode
func (c *Context) Run() error {
	if c == nil {
		return errors.New("nil pointer in interactive.Context.Run() function")
	}

	fmt.Fprintln(c.writer, headerStr)
	for true {
		res, ok := c.prompt()
		if !ok {
			break
		}
		if err := c.MakeLink(res); err != nil {
			fmt.Fprintln(c.writer, err)
		}
	}

	if err := c.scanner.Err(); err != nil {
		return err
	}
	return nil
}

func (c *Context) prompt() (string, bool) {
	if c == nil {
		return "", false
	}
	fmt.Fprint(c.writer, promptStr)
	if c.scanner.Scan() {
		res := strings.Trim(c.scanner.Text(), "\t ")
		switch res {
		case "q", "quit":
			return res, false
		default:
			return res, true
		}
	}
	return "", false
}
