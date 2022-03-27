package options

import (
	"context"
	"io"

	"github.com/goark/errs"
	"github.com/goark/ml/ecode"
	"github.com/goark/ml/facade/history"
	"github.com/goark/ml/makelink"
)

//Options class is Options for making link
type Options struct {
	linkStyle makelink.Style
	hist      *history.HistoryFile
	userAgent string
}

//New returns new Options instance
func New(s makelink.Style, hist *history.HistoryFile, userAgent string) *Options {
	if hist == nil {
		hist = history.NewFile(0, "")
	}
	return &Options{linkStyle: s, hist: hist, userAgent: userAgent}
}

//History method returns history.HistoryFile instance.
func (c *Options) History() *history.HistoryFile { return c.hist }

//MakeLink is making link
func (c *Options) MakeLink(ctx context.Context, urlStr string) (io.Reader, error) {
	if c == nil {
		return nil, errs.Wrap(ecode.ErrNullPointer)
	}
	c.History().Add(urlStr)
	lnk, err := makelink.New(ctx, urlStr, c.userAgent)
	if err != nil {
		return nil, errs.Wrap(err)
	}
	return lnk.Encode(c.linkStyle), nil
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
