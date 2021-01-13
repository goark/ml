package options_test

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"os"
	"testing"

	"github.com/spiegel-im-spiegel/ml/facade/history"
	"github.com/spiegel-im-spiegel/ml/facade/options"
	"github.com/spiegel-im-spiegel/ml/makelink"
)

func TestMakeLink(t *testing.T) {
	urlStr := "https://git.io/vFR5M"
	opt := options.New(makelink.StyleMarkdown, history.NewFile(1, ""))
	rRes, err := opt.MakeLink(context.Background(), urlStr)
	if err != nil {
		t.Errorf("Error in Context.MakeLink(): %+v", err)
	}
	outBuf := &bytes.Buffer{}
	if _, err := io.Copy(outBuf, rRes); err != nil {
		t.Errorf("Error in io.Copy(): %+v", err)
	}

	res := "[GitHub - spiegel-im-spiegel/ml: Make Link with Markdown Format](https://github.com/spiegel-im-spiegel/ml)"
	str := outBuf.String()
	if str != res {
		t.Errorf("Context.MakeLink() = \"%v\", want \"%v\".", str, res)
	}
	h := opt.History().At(0)
	if h != urlStr {
		t.Errorf("Histtory(0) = \"%v\" (%v), want \"%v\".", h, opt.History().Len(), urlStr)
	}
}

func TestMakeLinkNil(t *testing.T) {
	rRes, err := options.New(makelink.StyleMarkdown, nil).MakeLink(context.Background(), "https://git.io/vFR5M")
	if err != nil {
		t.Errorf("Error in Context.MakeLink(): %+v", err)
	}
	outBuf := new(bytes.Buffer)
	if _, err := io.Copy(outBuf, rRes); err != nil {
		t.Errorf("Error in io.Copy(): %+v", err)
	}

	res := "[GitHub - spiegel-im-spiegel/ml: Make Link with Markdown Format](https://github.com/spiegel-im-spiegel/ml)"
	str := outBuf.String()
	if str != res {
		t.Errorf("Context.MakeLink()  = \"%v\", want \"%v\".", str, res)
	}
}

func TestMakeLinkErr(t *testing.T) {
	_, err := options.New(makelink.StyleMarkdown, nil).MakeLink(context.Background(), "https://foo.bar")
	if err == nil {
		t.Error("Context.MakeLink() = nil error, not want nil error.")
	} else {
		fmt.Fprintf(os.Stderr, "info: %+v\n", err)
	}
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
