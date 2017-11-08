package mklink

import "testing"

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
			t.Errorf("GetStyles()  = \"%v\", want nil error.", err)
		} else if tps.String() != tst.t.String() {
			t.Errorf("GetStyles()  = \"%v\", want \"%v\".", tps, tst.t)
		}
	}
}

func TestGetStyleErr(t *testing.T) {
	tps, err := GetStyle("foobar")
	if err == nil {
		t.Error("GetStyles(foobar)  = nil error, not want nil error.")
	} else if tps.String() != "unknown" {
		t.Errorf("GetStyles(foobar)  = \"%v\", want \"unknown\".", tps)
	}
}

func TestStyleList(t *testing.T) {
	str := StyleList()
	res := "markdown wiki html csv"
	if str != res {
		t.Errorf("StylesList()  = \"%v\", want \"%v\".", str, res)
	}
}
