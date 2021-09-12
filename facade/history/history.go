package history

import (
	"bufio"
	"fmt"
	"io"

	"github.com/nyaosorg/go-readline-ny"
	"github.com/spiegel-im-spiegel/errs"
)

//History is a ring-buffer class for history (string) data.
type History struct {
	head, tail int
	buffer     []string
}

var _ readline.IHistory = (*History)(nil)

func New(size int) *History {
	if size < 1 {
		return nil
	}
	return &History{head: 0, tail: -1, buffer: make([]string, size)}
}

//Len method returns count of buffering strings.
func (hist *History) Len() int {
	if hist.Size() == 0 {
		return 0
	}
	if hist.tail < 0 {
		return 0
	}
	if hist.Size() == 1 {
		return 1
	}
	if hist.head == hist.tail {
		return hist.Size()
	}
	// if hist.head > hist.tail {
	// 	return (hist.tail + hist.Size()) - hist.head
	// }
	return hist.tail - hist.head
}

//At method returns .histroty (string) data at index.
func (hist *History) At(n int) string {
	if hist.Size() == 0 || n >= hist.Len() {
		return ""
	}
	i := (hist.head + n) % hist.Size()
	return hist.buffer[i]
}

//Size method returns size of history buffer.
func (hist *History) Size() int {
	if hist == nil {
		return 0
	}
	return len(hist.buffer)
}

//Add method adds new string in history buffer.
func (hist *History) Add(s string) {
	if hist.Size() == 0 || len(s) == 0 {
		return
	}
	if hist.tail < 0 {
		hist.buffer[0] = s
		hist.tail = 1 % hist.Size()
		return
	}
	if hist.Len() > 0 && hist.At(hist.Len()-1) == s {
		return
	}
	tail := hist.tail
	hist.buffer[tail] = s
	hist.tail = (hist.tail + 1) % hist.Size()
	if hist.head == tail {
		hist.head = hist.tail
	}
}

//Import method imports history data from reader.
func (hist *History) Import(r io.Reader) error {
	if hist.Size() == 0 {
		return nil
	}
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		hist.Add(scanner.Text())
	}
	return errs.Wrap(scanner.Err())
}

//Export method exports history data to writer.
func (hist *History) Export(w io.Writer) error {
	for i := 0; i < hist.Len(); i++ {
		if _, err := fmt.Fprintln(w, hist.At(i)); err != nil {
			return errs.Wrap(err)
		}
	}
	return nil
}

/* Copyright 2021 Spiegel
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
