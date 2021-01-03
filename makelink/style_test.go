package makelink

import (
	"errors"
	"fmt"
	"testing"

	"github.com/spiegel-im-spiegel/ml/ecode"
)

type typesTestCase struct {
	name string
	t    Style
}

var typesTests = []typesTestCase{
	{"markdown", StyleMarkdown},
	{"wiki", StyleWiki},
	{"html", StyleHTML},
	{"csv", StyleCSV},
}

func TestGetStyle(t *testing.T) {
	for _, tst := range typesTests {
		tps, err := GetStyle(tst.name)
		if err != nil {
			t.Errorf("GetStyles() = \"%+v\".", err)
		} else if tps.String() != tst.t.String() {
			t.Errorf("GetStyles() = \"%v\", want \"%v\".", tps, tst.t)
		}
	}
}

func TestGetStyleErr(t *testing.T) {
	tps, err := GetStyle("foobar")
	fmt.Printf("Info(TestGetStyleErr): %+v\n", err)
	if !errors.Is(err, ecode.ErrNoImplement) {
		t.Errorf("GetStyles(foobar) error = \"%v\", want \"%v\".", err, ecode.ErrNoImplement)
	} else if tps.String() != "unknown" {
		t.Errorf("GetStyles(foobar) = \"%v\", want \"unknown\".", tps)
	}
}

func TestStyleList(t *testing.T) {
	str := StyleList()
	res := "markdown|wiki|html|csv|json"
	if str != res {
		t.Errorf("StylesList() = \"%v\", want \"%v\".", str, res)
	}
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
