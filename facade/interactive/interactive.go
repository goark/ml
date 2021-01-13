package interactive

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"os"

	"github.com/atotto/clipboard"
	"github.com/spiegel-im-spiegel/errs"
	"github.com/spiegel-im-spiegel/gocli/signal"
	"github.com/spiegel-im-spiegel/ml/facade/options"
	"github.com/zetamatta/go-readline-ny"
)

func Do(opts *options.Options) error {
	editor := readline.Editor{
		Prompt:  func() (int, error) { return fmt.Print("ml> ") },
		History: opts.History(),
	}
	ctx := signal.Context(context.Background(), os.Interrupt)
	fmt.Println("Input 'q' or 'quit' to stop")
	for {
		text, err := editor.ReadLine(context.Background())
		if err != nil {
			return errs.Wrap(err)
		}
		if text == "q" || text == "quit" {
			return nil
		}
		r, err := opts.MakeLink(ctx, text)
		if err != nil {
			errStr := errs.Cause(err).Error()
			if errs.Is(err, context.Canceled) {
				return errs.Wrap(err)
			}
			fmt.Println(errStr)
		} else {
			buf := &bytes.Buffer{}
			if _, err := io.Copy(os.Stdout, io.TeeReader(r, buf)); err != nil {
				return errs.Wrap(err)
			}
			fmt.Println()
			_ = clipboard.WriteAll(buf.String())
		}
	}
}

/* Copyright 2019-2021 Spiegel
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
