package options_test

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/goark/ml/facade/history"
	"github.com/goark/ml/facade/options"
	"github.com/goark/ml/makelink"
)

func newTestServer() *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		_, _ = io.WriteString(w, `<!doctype html><html><head><title>Example Title</title><meta name="description" content="Example Description"><link rel="canonical" href="https://example.com/canonical"></head><body></body></html>`)
	}))
}

func TestMakeLink(t *testing.T) {
	srv := newTestServer()
	defer srv.Close()

	urlStr := srv.URL
	opt := options.New(makelink.StyleMarkdown, history.NewFile(1, ""), "")
	rRes, err := opt.MakeLink(context.Background(), urlStr)
	if err != nil {
		t.Errorf("Error in Context.MakeLink(): %+v", err)
	}
	outBuf := &bytes.Buffer{}
	if _, err := io.Copy(outBuf, rRes); err != nil {
		t.Errorf("Error in io.Copy(): %+v", err)
	}

	res := "[Example Title](https://example.com/canonical)"
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
	srv := newTestServer()
	defer srv.Close()

	rRes, err := options.New(makelink.StyleMarkdown, nil, "").MakeLink(context.Background(), srv.URL)
	if err != nil {
		t.Errorf("Error in Context.MakeLink(): %+v", err)
	}
	outBuf := new(bytes.Buffer)
	if _, err := io.Copy(outBuf, rRes); err != nil {
		t.Errorf("Error in io.Copy(): %+v", err)
	}

	res := "[Example Title](https://example.com/canonical)"
	str := outBuf.String()
	if str != res {
		t.Errorf("Context.MakeLink()  = \"%v\", want \"%v\".", str, res)
	}
}

func TestMakeLinkErr(t *testing.T) {
	_, err := options.New(makelink.StyleMarkdown, nil, "").MakeLink(context.Background(), "://bad-url")
	if err == nil {
		t.Error("Context.MakeLink() = nil error, not want nil error.")
	} else {
		fmt.Fprintf(os.Stderr, "info: %+v\n", err)
	}
}

/* Copyright 2017-2026 Spiegel
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
