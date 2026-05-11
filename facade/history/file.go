package history

import (
	"errors"
	"os"

	"github.com/goark/errs"
	"github.com/nyaosorg/go-readline-ny"
)

// HistoryFile stores history entries backed by a file.
type HistoryFile struct {
	*History
	path string
}

var _ readline.IHistory = (*HistoryFile)(nil)

// NewFile returns a new HistoryFile.
func NewFile(size int, path string) *HistoryFile {
	return &HistoryFile{History: New(size), path: path}
}

// Load imports history entries from the file.
func (hf *HistoryFile) Load() (err error) {
	if hf == nil || hf.Size() == 0 || len(hf.path) == 0 {
		return nil
	}
	file, err := os.Open(hf.path)
	if err != nil {
		// Missing history file on first run is expected.
		if errors.Is(err, os.ErrNotExist) {
			return nil
		}
		return errs.Wrap(err)
	}
	defer func() {
		if cerr := file.Close(); cerr != nil {
			err = errs.Join(err, cerr)
		}
	}()
	err = hf.Import(file)
	return
}

// Save writes history entries to the file.
func (hf *HistoryFile) Save() (err error) {
	if hf == nil || hf.Size() == 0 || len(hf.path) == 0 {
		return nil
	}
	// Use OpenFile to preserve explicit 0600 permission while truncating stale data.
	file, err := os.OpenFile(hf.path, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0600)
	if err != nil {
		return errs.Wrap(err)
	}
	defer func() {
		if cerr := file.Close(); cerr != nil {
			err = errs.Join(err, cerr)
		}
	}()
	err = hf.Export(file)
	return
}

/* Copyright 2021-2026 Spiegel
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
