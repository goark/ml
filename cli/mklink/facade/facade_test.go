package facade

import (
	"bytes"
	"fmt"
	"testing"

	"github.com/spiegel-im-spiegel/gocli/exitcode"
	"github.com/spiegel-im-spiegel/gocli/rwi"
)

func TestInteractiveError(t *testing.T) {
	outBuf := new(bytes.Buffer)
	outErrBuf := new(bytes.Buffer)
	ui := rwi.New(rwi.WithWriter(outBuf), rwi.WithErrorWriter(outErrBuf))
	args := []string{"-i"}

	exit := Execute(ui, args)
	if exit == exitcode.Normal {
		t.Errorf("Execute(markdown) = \"%v\", want \"%v\".", exit, exitcode.Abnormal)
	}
	fmt.Printf("Info: %v", outErrBuf.String())
}

func TestStyleMarkdown(t *testing.T) {
	outBuf := new(bytes.Buffer)
	outErrBuf := new(bytes.Buffer)
	ui := rwi.New(rwi.WithWriter(outBuf), rwi.WithErrorWriter(outErrBuf))
	args := []string{"https://text.baldanders.info"}

	exit := Execute(ui, args)
	if exit != exitcode.Normal {
		t.Errorf("Execute(markdown) = \"%v\", want \"%v\".", exit, exitcode.Normal)
	}
	str := outErrBuf.String()
	if str != "" {
		t.Errorf("Execute(markdown) = \"%v\", want \"%v\".", str, "")
	}
	str = outBuf.String()
	res := "[text.Baldanders.info](https://text.baldanders.info)\n"
	if str != res {
		t.Errorf("Execute(markdown) = \"%v\", want \"%v\".", str, res)
	}
}

func TestUrlErr(t *testing.T) {
	outBuf := new(bytes.Buffer)
	outErrBuf := new(bytes.Buffer)
	ui := rwi.New(rwi.WithWriter(outBuf), rwi.WithErrorWriter(outErrBuf))
	args := []string{"http://foo.bar"}

	exit := Execute(ui, args)
	if exit != exitcode.Abnormal {
		t.Errorf("Execute(markdown) = \"%v\", want \"%v\".", exit, exitcode.Abnormal)
	}
}

func TestStyleWiki(t *testing.T) {
	outBuf := new(bytes.Buffer)
	outErrBuf := new(bytes.Buffer)
	ui := rwi.New(rwi.WithWriter(outBuf), rwi.WithErrorWriter(outErrBuf))
	args := []string{"-s", "wiki", "https://text.baldanders.info"}

	exit := Execute(ui, args)
	if exit != exitcode.Normal {
		t.Errorf("Execute(wiki) = \"%v\", want \"%v\".", exit, exitcode.Normal)
	}
	str := outErrBuf.String()
	if str != "" {
		t.Errorf("Execute(wiki) = \"%v\", want \"%v\".", str, "")
	}
	str = outBuf.String()
	res := "[https://text.baldanders.info text.Baldanders.info]\n"
	if str != res {
		t.Errorf("Execute(wiki) = \"%v\", want \"%v\".", str, res)
	}
}

func TestPipe(t *testing.T) {
	inData := bytes.NewBufferString("https://text.baldanders.info")
	outBuf := new(bytes.Buffer)
	outErrBuf := new(bytes.Buffer)
	ui := rwi.New(rwi.WithReader(inData), rwi.WithWriter(outBuf), rwi.WithErrorWriter(outErrBuf))
	args := []string{}

	exit := Execute(ui, args)
	if exit != exitcode.Normal {
		t.Errorf("Execute(pipe) = \"%v\", want \"%v\".", exit, exitcode.Normal)
	}
	str := outErrBuf.String()
	if str != "" {
		t.Errorf("Execute(pipe) = \"%v\", want \"%v\".", str, "")
	}
	str = outBuf.String()
	res := "[text.Baldanders.info](https://text.baldanders.info)\n"
	if str != res {
		t.Errorf("Execute(pipe) = \"%v\", want \"%v\".", str, res)
	}
}

func TestPipeUrlErr(t *testing.T) {
	inData := bytes.NewBufferString("http://foo.bar")
	outBuf := new(bytes.Buffer)
	outErrBuf := new(bytes.Buffer)
	ui := rwi.New(rwi.WithReader(inData), rwi.WithWriter(outBuf), rwi.WithErrorWriter(outErrBuf))
	args := []string{}

	exit := Execute(ui, args)
	if exit != exitcode.Abnormal {
		t.Errorf("Execute(pipe) = \"%v\", want \"%v\".", exit, exitcode.Abnormal)
	}
}

/* Copyright 2017-2019 Spiegel
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 * 	http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */
