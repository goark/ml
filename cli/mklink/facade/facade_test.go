package facade

import (
	"bytes"
	"testing"

	"github.com/spiegel-im-spiegel/gocli"
)

func TestStyleMarkdown(t *testing.T) {
	outBuf := new(bytes.Buffer)
	outErrBuf := new(bytes.Buffer)
	ui := gocli.NewUI(gocli.Writer(outBuf), gocli.ErrorWriter(outErrBuf))
	args := []string{"http://text.baldanders.info"}

	clearFlags()
	exit := Execute(ui, args)
	if exit != ExitNormal {
		t.Errorf("Execute(markdown) = \"%v\", want \"%v\".", exit, ExitNormal)
	}
	str := outErrBuf.String()
	if str != "" {
		t.Errorf("Execute(markdown) = \"%v\", want \"%v\".", str, "")
	}
	str = outBuf.String()
	res := "[text.Baldanders.info](http://text.baldanders.info)\n"
	if str != res {
		t.Errorf("Execute(markdown) = \"%v\", want \"%v\".", str, res)
	}
}

func TestUrlErr(t *testing.T) {
	outBuf := new(bytes.Buffer)
	outErrBuf := new(bytes.Buffer)
	ui := gocli.NewUI(gocli.Writer(outBuf), gocli.ErrorWriter(outErrBuf))
	args := []string{"http://foo.bar"}

	clearFlags()
	exit := Execute(ui, args)
	if exit != ExitAbnormal {
		t.Errorf("Execute(markdown) = \"%v\", want \"%v\".", exit, ExitAbnormal)
	}
}

func TestStyleWiki(t *testing.T) {
	outBuf := new(bytes.Buffer)
	outErrBuf := new(bytes.Buffer)
	ui := gocli.NewUI(gocli.Writer(outBuf), gocli.ErrorWriter(outErrBuf))
	args := []string{"-s", "wiki", "http://text.baldanders.info"}

	clearFlags()
	exit := Execute(ui, args)
	if exit != ExitNormal {
		t.Errorf("Execute(wiki) = \"%v\", want \"%v\".", exit, ExitNormal)
	}
	str := outErrBuf.String()
	if str != "" {
		t.Errorf("Execute(wiki) = \"%v\", want \"%v\".", str, "")
	}
	str = outBuf.String()
	res := "[http://text.baldanders.info text.Baldanders.info]\n"
	if str != res {
		t.Errorf("Execute(wiki) = \"%v\", want \"%v\".", str, res)
	}
}

func TestPipe(t *testing.T) {
	inData := bytes.NewBufferString("http://text.baldanders.info")
	outBuf := new(bytes.Buffer)
	outErrBuf := new(bytes.Buffer)
	ui := gocli.NewUI(gocli.Reader(inData), gocli.Writer(outBuf), gocli.ErrorWriter(outErrBuf))
	args := []string{}

	clearFlags()
	exit := Execute(ui, args)
	if exit != ExitNormal {
		t.Errorf("Execute(pipe) = \"%v\", want \"%v\".", exit, ExitNormal)
	}
	str := outErrBuf.String()
	if str != "" {
		t.Errorf("Execute(pipe) = \"%v\", want \"%v\".", str, "")
	}
	str = outBuf.String()
	res := "[text.Baldanders.info](http://text.baldanders.info)\n"
	if str != res {
		t.Errorf("Execute(pipe) = \"%v\", want \"%v\".", str, res)
	}
}

func TestUrlErr2(t *testing.T) {
	inData := bytes.NewBufferString("http://foo.bar")
	outBuf := new(bytes.Buffer)
	outErrBuf := new(bytes.Buffer)
	ui := gocli.NewUI(gocli.Reader(inData), gocli.Writer(outBuf), gocli.ErrorWriter(outErrBuf))
	args := []string{}

	clearFlags()
	exit := Execute(ui, args)
	if exit != ExitAbnormal {
		t.Errorf("Execute(pipe) = \"%v\", want \"%v\".", exit, ExitAbnormal)
	}
}

func clearFlags() {
	rootCmd.Flag("version").Value.Set("false")
	rootCmd.Flag("interactive").Value.Set("false")
	rootCmd.Flag("style").Value.Set(defaultStyle)
	rootCmd.Flag("log").Value.Set("")
}
