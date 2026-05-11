package makelink

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"strings"

	"github.com/goark/errs"
	"github.com/goark/webinfo"
)

// Link stores metadata of a URL.
type Link struct {
	URL         string `json:"url,omitempty"`
	Location    string `json:"location,omitempty"`
	Canonical   string `json:"canonical,omitempty"`
	Title       string `json:"title,omitempty"`
	Description string `json:"description,omitempty"`
}

// New fetches URL metadata and returns a Link instance.
func New(ctx context.Context, urlStr, userAgent string) (link *Link, err error) {
	link = &Link{URL: urlStr}
	info, err := webinfo.Fetch(ctx, urlStr, userAgent)
	if err != nil {
		return link, errs.Wrap(err, errs.WithContext("url", urlStr))
	}
	link.URL = trimString(info.URL)
	if len(link.URL) == 0 {
		link.URL = urlStr
	}
	link.Location = trimString(info.Location)
	link.Canonical = trimString(info.Canonical)
	link.Title = trimString(info.Title)
	link.Description = trimString(info.Description)
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

// TitleName returns the display title.
func (lnk *Link) TitleName() string {
	if lnk == nil {
		return ""
	}
	if len(lnk.Title) > 0 {
		return lnk.Title
	}
	return lnk.URL
}

// CanonicalURL returns the canonical URL when available.
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

// Encode encodes Link in the requested style.
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

// String returns Link as JSON text.
func (lnk *Link) String() string {
	if lnk == nil {
		return ""
	}
	return fmt.Sprint(lnk.Encode(StyleJSON))
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
