package makelink

import (
	"bytes"
	"io"

	"github.com/spiegel-im-spiegel/mklink"
)

//Context class is context for making link
type Context struct {
	linkStyle mklink.Style
	writer    io.Writer
	log       io.Writer
}

//New returns new Context instance
func New(s mklink.Style, writer io.Writer, log io.Writer) *Context {
	return &Context{linkStyle: s, writer: writer, log: log}
}

//MakeLink is making link
func (c *Context) MakeLink(url string) error {
	lnk, err := mklink.New(url)
	if err != nil {
		return err
	}

	r := lnk.Encode(c.linkStyle)
	if c.log == nil {
		io.Copy(c.writer, r)
	} else {
		buf := new(bytes.Buffer)
		io.Copy(c.writer, io.TeeReader(r, buf))
		io.Copy(c.log, buf)
	}
	return nil
}
