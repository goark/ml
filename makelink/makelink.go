package makelink

import (
	"bufio"
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/goark/errs"
	"github.com/goark/fetch"
	encoding "github.com/mattn/go-encoding"
	"golang.org/x/net/html/charset"
)

// Link class is information of URL
type Link struct {
	URL         string `json:"url,omitempty"`
	Location    string `json:"location,omitempty"`
	Canonical   string `json:"canonical,omitempty"`
	Title       string `json:"title,omitempty"`
	Description string `json:"description,omitempty"`
}

// New returns new Link instance
func New(ctx context.Context, urlStr, userAgent string) (link *Link, err error) {
	link = &Link{URL: urlStr}
	u, err := fetch.URL(urlStr)
	if err != nil {
		return link, errs.Wrap(err, errs.WithContext("url", urlStr))
	}
	if len(userAgent) == 0 {
		userAgent = "goark/ml (+https://github.com/goark/ml)" //dummy user-agent string
	}
	resp, err := fetch.New(fetch.WithHTTPClient(&http.Client{})).GetWithContext(
		ctx,
		u,
		fetch.WithRequestHeaderSet("User-Agent", userAgent),
	)
	if err != nil {
		return link, errs.Wrap(err, errs.WithContext("url", urlStr))
	}
	defer func() {
		if cerr := resp.Close(); cerr != nil {
			err = errs.Join(err, cerr)
		}
	}()

	link.Location = resp.Request().URL.String()

	br := bufio.NewReader(resp.Body())
	var r io.Reader = br
	if data, err2 := br.Peek(1024); err2 == nil { //next 1024 bytes without advancing the reader.
		enc, name, _ := charset.DetermineEncoding(data, resp.Header().Get("content-type"))
		if enc != nil {
			r = enc.NewDecoder().Reader(br)
		} else if len(name) > 0 {
			if enc := encoding.GetEncoding(name); enc != nil {
				r = enc.NewDecoder().Reader(br)
			}
		}
	}
	doc, err := goquery.NewDocumentFromReader(r)
	if err != nil {
		err = errs.Wrap(err)
		return
	}

	doc.Find("head").Each(func(_ int, s *goquery.Selection) {
		s.Find("title").Each(func(_ int, s *goquery.Selection) {
			t := s.Text()
			if len(t) > 0 {
				link.Title = trimString(t)
			}
		})
		s.Find("meta[name='description']").Each(func(_ int, s *goquery.Selection) {
			if v, ok := s.Attr("content"); ok {
				if len(v) > 0 {
					link.Description = trimString(v)
				}
			}
		})
		s.Find("link[rel='canonical']").Each(func(_ int, s *goquery.Selection) {
			if v, ok := s.Attr("href"); ok {
				if len(v) > 0 {
					link.Canonical = trimString(v)
				}
			}
		})
	})
	return
}

var replacer = strings.NewReplacer(
	"\r\n", " ",
	"\r", " ",
	"\n", " ",
)

func trimString(s string) string {
	return strings.TrimSpace(replacer.Replace(s))
}

// TitleName returns string of title name
func (lnk *Link) TitleName() string {
	if lnk == nil {
		return ""
	}
	if len(lnk.Title) > 0 {
		return lnk.Title
	}
	return lnk.URL
}

// CanonicalURL returns the canonical URL.
func (lnk *Link) CanonicalURL() string {
	if lnk == nil {
		return ""
	}
	if len(lnk.Canonical) > 0 {
		return lnk.Canonical
	}
	if len(lnk.Location) > 0 {
		return lnk.Location
	}
	return lnk.URL
}

// Encode returns string (io.Reader) with other style
func (lnk *Link) Encode(t Style) io.Reader {
	if lnk == nil {
		return io.NopCloser(bytes.NewReader(nil))
	}
	buf := &bytes.Buffer{}
	switch t {
	case StyleMarkdown:
		fmt.Fprintf(buf, "[%s](%s)", lnk.TitleName(), lnk.CanonicalURL())
	case StyleWiki:
		fmt.Fprintf(buf, "[%s %s]", lnk.CanonicalURL(), lnk.TitleName())
	case StyleHTML:
		fmt.Fprintf(buf, "<a href=\"%s\">%s</a>", lnk.CanonicalURL(), lnk.TitleName())
	case StyleCSV:
		fmt.Fprintf(buf, "\"%s\",\"%s\",\"%s\",\"%s\",\"%s\"", escapeQuoteCsv(lnk.URL), escapeQuoteCsv(lnk.Location), escapeQuoteCsv(lnk.Canonical), escapeQuoteCsv(lnk.Title), escapeQuoteCsv(lnk.Description))
	case StyleJSON:
		_ = json.NewEncoder(buf).Encode(lnk)
	}
	return buf
}
func escapeQuoteCsv(s string) string {
	return strings.ReplaceAll(s, "\"", "\"\"")
}

func (lnk *Link) String() string {
	if lnk == nil {
		return ""
	}
	return fmt.Sprint(lnk.Encode(StyleJSON))
}

/* Copyright 2017-2025 Spiegel
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
