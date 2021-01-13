package history

import (
	"os"

	"github.com/spiegel-im-spiegel/errs"
	"github.com/zetamatta/go-readline-ny"
)

//HistoryFile is a history file class.
type HistoryFile struct {
	*History
	path string
}

var _ readline.IHistory = (*HistoryFile)(nil)

//NewFile function returns new HistoryFile instance.
func NewFile(size int, path string) *HistoryFile {
	return &HistoryFile{History: New(size), path: path}
}

//Load method imports history data from file.
func (hf *HistoryFile) Load() error {
	if hf == nil || hf.Size() == 0 || len(hf.path) == 0 {
		return nil
	}
	file, err := os.Open(hf.path)
	if err != nil {
		return errs.Wrap(err)
	}
	defer file.Close()
	return hf.Import(file)
}

//Save method exports history data to file.
func (hf *HistoryFile) Save() error {
	if hf == nil || hf.Size() == 0 || len(hf.path) == 0 {
		return nil
	}
	// file, err := os.Create(hf.path)
	file, err := os.OpenFile(hf.path, os.O_RDWR|os.O_CREATE, 0600)
	if err != nil {
		return errs.Wrap(err)
	}
	defer file.Close()
	return hf.Export(file)
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
