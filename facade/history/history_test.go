package history_test

import (
	"bytes"
	"strings"
	"testing"

	"github.com/spiegel-im-spiegel/ml/facade/history"
)

func TestHistory(t *testing.T) {
	testCases := []struct {
		size int
		inp  []string
		adds []string
		out  []string
	}{
		{size: 5, inp: []string{"alice", "bob"}, adds: []string{"chris", "dan"}, out: []string{"alice", "bob", "chris", "dan", ""}},
		{size: 4, inp: []string{"alice", "bob"}, adds: []string{"chris", "dan"}, out: []string{"alice", "bob", "chris", "dan", ""}},
		{size: 3, inp: []string{"alice", "bob"}, adds: []string{"chris", "dan"}, out: []string{"bob", "chris", "dan", ""}},
		{size: 3, inp: []string{"alice", "bob"}, adds: []string{"chris", "dan", "dan"}, out: []string{"bob", "chris", "dan", ""}},
		{size: 3, inp: []string{"alice", "bob"}, adds: []string{"chris", "dan", "elen"}, out: []string{"chris", "dan", "elen", ""}},
		{size: 1, inp: []string{"alice"}, adds: []string{"chris", "dan"}, out: []string{"dan", ""}},
		{size: 1, inp: []string{}, adds: []string{"alice"}, out: []string{"alice", ""}},
		{size: 0, inp: []string{}, adds: []string{"alice"}, out: []string{}},
	}
	for _, tc := range testCases {
		hist := history.New(tc.size)
		if err := hist.Import(strings.NewReader(strings.Join(tc.inp, "\n"))); err != nil {
			t.Errorf("History.Import(%v) is \"%+v\", want nil", tc.inp, err)
		} else {
			if hist.Len() != len(tc.inp) {
				t.Errorf("History.Len() is %v, want %v", hist.Len(), len(tc.inp))
			}
			if len(tc.inp) > 0 && hist.At(0) != tc.inp[0] {
				t.Errorf("History.At(0) is %v, want %v", hist.At(0), tc.inp[0])
			}
			for _, add := range tc.adds {
				hist.Add(add)
			}
			buf := &bytes.Buffer{}
			if err := hist.Export(buf); err != nil {
				t.Errorf("History.Export() is \"%+v\", want nil", err)
			} else if buf.String() != strings.Join(tc.out, "\n") {
				t.Errorf("History.Export() is \"%v\" (%v), want \"%v\"", buf.String(), hist.Len(), strings.Join(tc.out, "\n"))
			}
		}
	}
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
