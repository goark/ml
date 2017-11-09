package makelink

import (
	"bytes"
	"fmt"
	"io"

	"github.com/atotto/clipboard"
	"github.com/pkg/errors"
	"github.com/spiegel-im-spiegel/mklink"
)

//Context class is context for making link
type Context struct {
	linkStyle   mklink.Style
	writer      io.Writer
	log         io.Writer
	clipbrdFlag bool
}

//New returns new Context instance
func New(s mklink.Style, writer io.Writer, log io.Writer) *Context {
	return &Context{linkStyle: s, writer: writer, log: log, clipbrdFlag: false}
}

//MakeLink is making link
func (c *Context) MakeLink(url string) error {
	if c == nil {
		return errors.New("nil pointer in makelink.Context.MakeLink() function")
	}
	lnk, err := mklink.New(url)
	if err != nil {
		return err
	}

	r := lnk.Encode(c.linkStyle)
	buf := new(bytes.Buffer)
	io.Copy(c.writer, io.TeeReader(r, buf))
	strLink := buf.String()
	if c.clipbrdFlag {
		clipboard.WriteAll(strLink)
	}
	if c.log != nil {
		fmt.Fprint(c.log, strLink)
	}
	return nil
}

//EnableClipboard is enabling to clipboard
func (c *Context) EnableClipboard() {
	if c != nil {
		c.clipbrdFlag = true
	}
}

//DisableClipboard is disabling to clipboard
func (c *Context) DisableClipboard() {
	if c != nil {
		c.clipbrdFlag = false
	}
}
