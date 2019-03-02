package mklink

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/PuerkitoBio/goquery"
	encoding "github.com/mattn/go-encoding"
	"github.com/spiegel-im-spiegel/mklink/errs"
	"golang.org/x/net/html/charset"
)

//Link class is information of URL
type Link struct {
	URL         string `json:"url,omitempty"`
	Location    string `json:"location,omitempty"`
	Title       string `json:"title,omitempty"`
	Description string `json:"description,omitempty"`
}

//New returns new Link instance
func New(url string) (*Link, error) {
	link := &Link{URL: trimString(url)}
	resp, err := http.Get(url)
	if err != nil {
		return link, errs.Wrap(err, "Create mklink.Link instance")
	}
	defer resp.Body.Close()

	link.Location = resp.Request.URL.String()

	br := bufio.NewReader(resp.Body)
	var r io.Reader = br
	if data, err2 := br.Peek(1024); err2 == nil { //next 1024 bytes without advancing the reader.
		enc, name, _ := charset.DetermineEncoding(data, resp.Header.Get("content-type"))
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
		return link, errs.Wrap(err, "Create mklink.Link instance")
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
	})

	return link, nil
}

var replacer = strings.NewReplacer(
	"\r\n", " ",
	"\r", " ",
	"\n", " ",
)

func trimString(s string) string {
	return strings.TrimSpace(replacer.Replace(s))
}

//JSON returns string (io.Reader) with JSON format
func (lnk *Link) JSON() (io.Reader, error) {
	if lnk == nil {
		return ioutil.NopCloser(bytes.NewReader(nil)), nil
	}
	buf := new(bytes.Buffer)
	encoder := json.NewEncoder(buf)
	encoder.SetIndent("", "  ")
	if err := encoder.Encode(lnk); err != nil {
		return ioutil.NopCloser(bytes.NewReader(nil)), errs.Wrap(err, "error in encoding JSON")
	}
	return buf, nil
}

//TitleName returns string of title name
func (lnk *Link) TitleName() string {
	if lnk == nil {
		return ""
	}
	if len(lnk.Title) > 0 {
		return lnk.Title
	}
	return lnk.URL
}

//Encode returns string (io.Reader) with other style
func (lnk *Link) Encode(t Style) io.Reader {
	if lnk == nil {
		return ioutil.NopCloser(bytes.NewReader(nil))
	}
	buf := new(bytes.Buffer)
	switch t {
	case StyleMarkdown:
		fmt.Fprintf(buf, "[%s](%s)", lnk.TitleName(), lnk.Location)
	case StyleWiki:
		fmt.Fprintf(buf, "[%s %s]", lnk.Location, lnk.TitleName())
	case StyleHTML:
		fmt.Fprintf(buf, "<a href=\"%s\">%s</a>", lnk.Location, lnk.TitleName())
	case StyleCSV:
		fmt.Fprintf(buf, "\"%s\",\"%s\",\"%s\",\"%s\"", escapeQuoteCsv(lnk.URL), escapeQuoteCsv(lnk.Location), escapeQuoteCsv(lnk.Title), escapeQuoteCsv(lnk.Description))
	}
	return buf
}
func escapeQuoteCsv(s string) string {
	return strings.Replace(s, "\"", "\"\"", -1)
}

func (lnk *Link) String() string {
	if lnk == nil {
		return ""
	}
	r, _ := lnk.JSON()
	return fmt.Sprint(r)
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
