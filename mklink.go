package mklink

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/pkg/errors"
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
	link := &Link{URL: strings.Trim(url, "\t \n")}
	doc, err := goquery.NewDocument(link.URL)
	if err != nil {
		return link, err
	}
	link.Location = doc.Url.String()

	doc.Find("head").Each(func(_ int, s *goquery.Selection) {
		s.Find("title").Each(func(_ int, s *goquery.Selection) {
			link.Title = strings.Trim(s.Text(), "\t \n")
		})
		s.Find("meta[name='description']").Each(func(_ int, s *goquery.Selection) {
			if v, ok := s.Attr("content"); ok {
				link.Description = strings.Trim(v, "\t \n")
			}
		})
	})

	return link, nil
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
		return ioutil.NopCloser(bytes.NewReader(nil)), errors.Wrap(err, "error in mklink.Link.JSON() function")
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
		fmt.Fprintf(buf, "[%s](%s)\n", lnk.TitleName(), lnk.Location)
	case StyleWiki:
		fmt.Fprintf(buf, "[%s %s]\n", lnk.Location, lnk.TitleName())
	case StyleHTML:
		fmt.Fprintf(buf, "<a href=\"%s\">%s</a>\n", lnk.Location, lnk.TitleName())
	case StyleCSV:
		fmt.Fprintf(buf, "\"%s\",\"%s\",\"%s\",\"%s\"\n", lnk.URL, lnk.Location, strings.Replace(lnk.Title, "\"", "\"\"", -1), strings.Replace(lnk.Description, "\"", "\"\"", -1))
	}
	return buf
}

func (lnk *Link) String() string {
	if lnk == nil {
		return ""
	}
	r, _ := lnk.JSON()
	return fmt.Sprint(r)
}
