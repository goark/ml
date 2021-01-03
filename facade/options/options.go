package options

import (
	"bytes"
	"context"
	"fmt"
	"io"

	"github.com/spiegel-im-spiegel/errs"
	"github.com/spiegel-im-spiegel/ml/ecode"
	"github.com/spiegel-im-spiegel/ml/makelink"
)

//Options class is Options for making link
type Options struct {
	ctx       context.Context
	linkStyle makelink.Style
	log       io.Writer
}

//New returns new Options instance
func New(ctx context.Context, s makelink.Style, log io.Writer) *Options {
	return &Options{ctx: ctx, linkStyle: s, log: log}
}

//MakeLink is making link
func (c *Options) MakeLink(urlStr string) (io.Reader, error) {
	if c == nil {
		return nil, errs.Wrap(ecode.ErrNullPointer)
	}
	lnk, err := makelink.New(c.ctx, urlStr)
	if err != nil {
		return nil, errs.Wrap(err)
	}

	rRes := lnk.Encode(c.linkStyle)
	if c.log != nil {
		buf := &bytes.Buffer{}
		if _, err := io.Copy(c.log, io.TeeReader(rRes, buf)); err != nil {
			return buf, errs.New("error in logging", errs.WithCause(err))
		}
		fmt.Fprintln(c.log) //new line in logfile
		return buf, nil
	}
	return rRes, nil
}

//Context returns context.Context instance in Options.
func (c *Options) Context() context.Context {
	return c.ctx
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
