package mklink

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"testing"
)

var typesTests2 = []typesTestCase{
	{"[GitHub - spiegel-im-spiegel/mklink: Make Link with Markdown Format](https://github.com/spiegel-im-spiegel/mklink)", StyleMarkdown},
	{"[https://github.com/spiegel-im-spiegel/mklink GitHub - spiegel-im-spiegel/mklink: Make Link with Markdown Format]", StyleWiki},
	{"<a href=\"https://github.com/spiegel-im-spiegel/mklink\">GitHub - spiegel-im-spiegel/mklink: Make Link with Markdown Format</a>", StyleHTML},
	{"\"https://git.io/vFR5M\",\"https://github.com/spiegel-im-spiegel/mklink\",\"GitHub - spiegel-im-spiegel/mklink: Make Link with Markdown Format\",\"mklink - Make Link with Markdown Format\"", StyleCSV},
	{"", StyleUnknown},
}

func TestEncode(t *testing.T) {
	lnk := &Link{URL: "https://git.io/vFR5M", Location: "https://github.com/spiegel-im-spiegel/mklink", Title: "GitHub - spiegel-im-spiegel/mklink: Make Link with Markdown Format", Description: "mklink - Make Link with Markdown Format"}
	for _, tst := range typesTests2 {
		r := lnk.Encode(tst.t)
		buf := new(bytes.Buffer)
		if _, err := io.Copy(buf, r); err != nil {
			t.Errorf("io.Copy()  = %v, want nil.", err)
		}
		str := buf.String()
		if str != tst.name {
			t.Errorf("Encode(%v)  = \"%v\", want \"%v\".", tst.t, str, tst.name)
		}
	}
}

var typesTests3 = []typesTestCase{
	{"[https://git.io/vFR5M](https://github.com/spiegel-im-spiegel/mklink)", StyleMarkdown},
	{"[https://github.com/spiegel-im-spiegel/mklink https://git.io/vFR5M]", StyleWiki},
	{"<a href=\"https://github.com/spiegel-im-spiegel/mklink\">https://git.io/vFR5M</a>", StyleHTML},
	{"\"https://git.io/vFR5M\",\"https://github.com/spiegel-im-spiegel/mklink\",\"\",\"\"", StyleCSV},
	{"", StyleUnknown},
}

func TestEncodeNoUTF8(t *testing.T) {
	lnk := &Link{URL: "https://git.io/vFR5M", Location: "https://github.com/spiegel-im-spiegel/mklink", Title: "", Description: ""}
	for _, tst := range typesTests3 {
		r := lnk.Encode(tst.t)
		buf := new(bytes.Buffer)
		if _, err := io.Copy(buf, r); err != nil {
			t.Errorf("io.Copy()  = %v, want nil.", err)
		}
		str := buf.String()
		if str != tst.name {
			t.Errorf("Encode(%v)  = \"%v\", want \"%v\".", tst.t, str, tst.name)
		}
	}
}

func TestString(t *testing.T) {
	lnk := &Link{URL: "https://git.io/vFR5M", Location: "https://github.com/spiegel-im-spiegel/mklink", Title: "GitHub - spiegel-im-spiegel/mklink: Make Link with Markdown Format", Description: "mklink - Make Link with Markdown Format"}
	str := lnk.String()
	res := `{
  "url": "https://git.io/vFR5M",
  "location": "https://github.com/spiegel-im-spiegel/mklink",
  "title": "GitHub - spiegel-im-spiegel/mklink: Make Link with Markdown Format",
  "description": "mklink - Make Link with Markdown Format"
}
`
	if str != res {
		t.Errorf("New()  = \"%v\", want \"%v\".", str, res)
	}
}

func TestNewErr(t *testing.T) {
	_, err := New("https://foo.bar")
	if err == nil {
		t.Error("New()  = nil error, not want nil error.")
	} else {
		fmt.Fprintf(os.Stderr, "info: %v\n", err)
	}
}

func ExampleNew() {
	link, err := New("https://git.io/vFR5M")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(link.Encode(StyleMarkdown))
	// Output:
	// [GitHub - spiegel-im-spiegel/mklink: Make Link with Markdown Format](https://github.com/spiegel-im-spiegel/mklink)
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
