package options

import (
	"context"
	"io"

	"github.com/spiegel-im-spiegel/errs"
	"github.com/spiegel-im-spiegel/ml/ecode"
	"github.com/spiegel-im-spiegel/ml/facade/history"
	"github.com/spiegel-im-spiegel/ml/makelink"
)

//Options class is Options for making link
type Options struct {
	ctx       context.Context
	linkStyle makelink.Style
	hist      *history.HistoryFile
}

//New returns new Options instance
func New(ctx context.Context, s makelink.Style, hist *history.HistoryFile) *Options {
	if hist == nil {
		hist = history.NewFile(0, "")
	}
	return &Options{ctx: ctx, linkStyle: s, hist: hist}
}

//History method returns history.HistoryFile instance.
func (c *Options) History() *history.HistoryFile { return c.hist }

//MakeLink is making link
func (c *Options) MakeLink(urlStr string) (io.Reader, error) {
	if c == nil {
		return nil, errs.Wrap(ecode.ErrNullPointer)
	}
	c.History().Add(urlStr)
	lnk, err := makelink.New(c.ctx, urlStr)
	if err != nil {
		return nil, errs.Wrap(err)
	}
	return lnk.Encode(c.linkStyle), nil
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
