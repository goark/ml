package makelink

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"os"
	"testing"
)

var typesTests2 = []typesTestCase{
	{"[GitHub - spiegel-im-spiegel/ml: Make Link with Markdown Format](https://github.com/spiegel-im-spiegel/ml)", StyleMarkdown},
	{"[https://github.com/spiegel-im-spiegel/ml GitHub - spiegel-im-spiegel/ml: Make Link with Markdown Format]", StyleWiki},
	{"<a href=\"https://github.com/spiegel-im-spiegel/ml\">GitHub - spiegel-im-spiegel/ml: Make Link with Markdown Format</a>", StyleHTML},
	{"\"https://git.io/vFR5M\",\"https://github.com/spiegel-im-spiegel/ml\",\"\",\"GitHub - spiegel-im-spiegel/ml: Make Link with Markdown Format\",\"ml - Make Link with Markdown Format\"", StyleCSV},
	{"", StyleUnknown},
}

func TestEncode(t *testing.T) {
	lnk := &Link{URL: "https://git.io/vFR5M", Location: "https://github.com/spiegel-im-spiegel/ml", Title: "GitHub - spiegel-im-spiegel/ml: Make Link with Markdown Format", Description: "ml - Make Link with Markdown Format"}
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
	{"[https://git.io/vFR5M](https://github.com/spiegel-im-spiegel/ml)", StyleMarkdown},
	{"[https://github.com/spiegel-im-spiegel/ml https://git.io/vFR5M]", StyleWiki},
	{"<a href=\"https://github.com/spiegel-im-spiegel/ml\">https://git.io/vFR5M</a>", StyleHTML},
	{"\"https://git.io/vFR5M\",\"https://github.com/spiegel-im-spiegel/ml\",\"https://github.com/spiegel-im-spiegel/ml\",\"\",\"\"", StyleCSV},
	{"", StyleUnknown},
}

func TestEncodeNoUTF8(t *testing.T) {
	lnk := &Link{URL: "https://git.io/vFR5M", Location: "https://github.com/spiegel-im-spiegel/ml", Canonical: "https://github.com/spiegel-im-spiegel/ml", Title: "", Description: ""}
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
	lnk := &Link{URL: "https://git.io/vFR5M", Location: "https://github.com/spiegel-im-spiegel/ml", Title: "GitHub - spiegel-im-spiegel/ml: Make Link with Markdown Format", Description: "ml - Make Link with Markdown Format"}
	str := lnk.String()
	res := `{"url":"https://git.io/vFR5M","location":"https://github.com/spiegel-im-spiegel/ml","title":"GitHub - spiegel-im-spiegel/ml: Make Link with Markdown Format","description":"ml - Make Link with Markdown Format"}
`
	if str != res {
		t.Errorf("New()  = \"%v\", want \"%v\".", str, res)
	}
}

func TestNewErr(t *testing.T) {
	_, err := New(context.Background(), "https://foo.bar")
	if err == nil {
		t.Error("New()  = nil error, not want nil error.")
	} else {
		fmt.Fprintf(os.Stderr, "info: %+v\n", err)
	}
}

func ExampleNew() {
	link, err := New(context.Background(), "https://git.io/vFR5M")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(link.Encode(StyleMarkdown))
	// Output:
	// [GitHub - spiegel-im-spiegel/ml: Make Link with Markdown Format](https://github.com/spiegel-im-spiegel/ml)
}

/* Copyright 2017-2021 Spiegel
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
