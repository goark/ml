package facade

import (
	"bytes"
	"io"

	"github.com/atotto/clipboard"
	"github.com/spiegel-im-spiegel/gocli/prompt"
	"github.com/spiegel-im-spiegel/gocli/rwi"
	"github.com/spiegel-im-spiegel/mklink/cli/mklink/makelink"
	"github.com/spiegel-im-spiegel/mklink/errs"
	errors "golang.org/x/xerrors"
)

func interactiveMode(ui *rwi.RWI, cxt *makelink.Context) error {
	p := prompt.New(
		rwi.New(
			rwi.WithReader(ui.Reader()),
			rwi.WithWriter(ui.Writer()),
		),
		func(url string) (string, error) {
			if url == "q" || url == "quit" {
				return "", prompt.ErrTerminate
			}
			r, err := cxt.MakeLink(url)
			if err != nil {
				return errs.Cause(err).Error(), nil
			}
			buf := &bytes.Buffer{}
			if _, err := io.Copy(buf, r); err != nil {
				return "", errs.Wrap(err, "error when output result")
			}
			res := buf.String()
			return res, errs.Wrap(clipboard.WriteAll(res), "error when output result")
		},
		prompt.WithPromptString("mklink> "),
		prompt.WithHeaderMessage("Input 'q' or 'quit' to stop"),
	)
	if !p.IsTerminal() {
		return errors.New("not terminal (or pipe?)")
	}
	return errs.Wrap(p.Run(), "error in interactive mode")
}

/* Copyright 2019 Spiegel
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