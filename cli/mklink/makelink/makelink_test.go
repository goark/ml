package makelink

import (
	"bytes"
	"io"
	"testing"

	"github.com/spiegel-im-spiegel/mklink"
)

func TestMakeLink(t *testing.T) {
	logBuf := new(bytes.Buffer)
	rRes, err := New(mklink.StyleMarkdown, logBuf).MakeLink("https://git.io/vFR5M")
	if err != nil {
		t.Errorf("Error of Context.MakeLink()  = %v, want nil.", err)
	}
	outBuf := new(bytes.Buffer)
	if _, err := io.Copy(outBuf, rRes); err != nil {
		t.Errorf("Error of io.Copy()  = %v, want nil.", err)
	}

	res := "[GitHub - spiegel-im-spiegel/mklink: Make Link with Markdown Format](https://github.com/spiegel-im-spiegel/mklink)"
	str := outBuf.String()
	if str != res {
		t.Errorf("Context.MakeLink()  = \"%v\", want \"%v\".", str, res)
	}
	str = logBuf.String()
	if str != res+"\n" {
		t.Errorf("Context.MakeLink()  = \"%v\", want \"%v\".", str, res+"\n")
	}
}

func TestMakeLinkNil(t *testing.T) {
	rRes, err := New(mklink.StyleMarkdown, nil).MakeLink("https://git.io/vFR5M")
	if err != nil {
		t.Errorf("Context.MakeLink()  = \"%v\", want nil error.", err)
	}
	outBuf := new(bytes.Buffer)
	if _, err := io.Copy(outBuf, rRes); err != nil {
		t.Errorf("Error of io.Copy()  = %v, want nil.", err)
	}

	res := "[GitHub - spiegel-im-spiegel/mklink: Make Link with Markdown Format](https://github.com/spiegel-im-spiegel/mklink)"
	str := outBuf.String()
	if str != res {
		t.Errorf("Context.MakeLink()  = \"%v\", want \"%v\".", str, res)
	}
}

func TestMakeLinkErr(t *testing.T) {
	_, err := New(mklink.StyleMarkdown, nil).MakeLink("https://foo.bar")
	if err == nil {
		t.Error("Context.MakeLink() = nil error, not want nil error.")
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
