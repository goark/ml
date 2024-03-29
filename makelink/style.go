package makelink

import (
	"strings"

	"github.com/goark/errs"
	"github.com/goark/ml/ecode"
)

//Style as link style
type Style int

const (
	//StyleUnknown is unknown link style
	StyleUnknown Style = iota
	//StyleMarkdown is unknown markdown style
	StyleMarkdown
	//StyleWiki is unknown wiki style
	StyleWiki
	//StyleHTML is unknown HTML anchor style
	StyleHTML
	//StyleCSV is CSV data format
	StyleCSV
	//StyleJSON is JSON format
	StyleJSON
)

var (
	styleMap = map[Style]string{
		StyleMarkdown: "markdown",
		StyleWiki:     "wiki",
		StyleHTML:     "html",
		StyleCSV:      "csv",
		StyleJSON:     "json",
	}
	styleList = []string{
		styleMap[StyleMarkdown],
		styleMap[StyleWiki],
		styleMap[StyleHTML],
		styleMap[StyleCSV],
		styleMap[StyleJSON],
	}
)

//StyleList returns string Style list
func StyleList() string {
	return strings.Join(styleList, "|")
}

//GetStyle returns Style from string
func GetStyle(s string) (Style, error) {
	for t, v := range styleMap {
		if strings.EqualFold(v, s) {
			return t, nil
		}
	}
	return StyleUnknown, errs.Wrap(ecode.ErrNoImplement, errs.WithContext("style", s))
}

func (t Style) String() string {
	if name, ok := styleMap[t]; ok {
		return name
	}
	return "unknown"
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
