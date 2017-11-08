package makelink

import (
	"bytes"
	"testing"

	"github.com/spiegel-im-spiegel/mklink"
)

func TestMakeLink(t *testing.T) {
	outBuf := new(bytes.Buffer)
	logBuf := new(bytes.Buffer)
	cxt := New(mklink.StyleMarkdown, outBuf, logBuf)
	err := cxt.MakeLink("https://git.io/vFR5M")
	if err != nil {
		t.Errorf("MakeLink()  = \"%v\", want nil error.", err)
	}
	res := "[GitHub - spiegel-im-spiegel/mklink: Make Link with Markdown Format](https://github.com/spiegel-im-spiegel/mklink)\n"
	str := outBuf.String()
	if str != res {
		t.Errorf("MakeLink()  = \"%v\", want \"%v\".", str, res)
	}
	str = logBuf.String()
	if str != res {
		t.Errorf("MakeLink()  = \"%v\", want \"%v\".", str, res)
	}
}

func TestMakeLinkNil(t *testing.T) {
	outBuf := new(bytes.Buffer)
	cxt := New(mklink.StyleMarkdown, outBuf, nil)
	err := cxt.MakeLink("https://git.io/vFR5M")
	if err != nil {
		t.Errorf("MakeLink()  = \"%v\", want nil error.", err)
	}
	res := "[GitHub - spiegel-im-spiegel/mklink: Make Link with Markdown Format](https://github.com/spiegel-im-spiegel/mklink)\n"
	str := outBuf.String()
	if str != res {
		t.Errorf("MakeLink()  = \"%v\", want \"%v\".", str, res)
	}
}

func TestMakeLinkErr(t *testing.T) {
	outBuf := new(bytes.Buffer)
	cxt := New(mklink.StyleMarkdown, outBuf, nil)
	err := cxt.MakeLink("https://foo.bar")
	if err == nil {
		t.Error("MakeLink()  = nil error, not want nil error.")
	}
}
