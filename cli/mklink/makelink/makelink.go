package makelink

import (
	"bytes"
	"fmt"
	"io"

	"github.com/pkg/errors"
	"github.com/spiegel-im-spiegel/mklink"
)

//Context class is context for making link
type Context struct {
	linkStyle mklink.Style
	log       io.Writer
}

//New returns new Context instance
func New(s mklink.Style, log io.Writer) *Context {
	return &Context{linkStyle: s, log: log}
}

//MakeLink is making link
func (c *Context) MakeLink(url string) (io.Reader, error) {
	if c == nil {
		return nil, errors.New("nil pointer in makelink.Context.MakeLink() function")
	}
	lnk, err := mklink.New(url)
	if err != nil {
		return nil, err
	}

	rRes := lnk.Encode(c.linkStyle)
	if c.log == nil {
		return rRes, nil
	}
	buf := new(bytes.Buffer)
	if _, err := io.Copy(c.log, io.TeeReader(rRes, buf)); err != nil {
		return buf, err
	}
	fmt.Fprintln(c.log) //new line in logfile
	return buf, nil
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
