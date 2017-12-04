package mklink

import (
	"bytes"
	"io"

	"github.com/saintfish/chardet"
	"golang.org/x/text/encoding/japanese"
	"golang.org/x/text/transform"
)

//CharEncode is type of character encoding
type CharEncode int

const (
	//CharUnknown is unknown character
	CharUnknown CharEncode = iota
	//CharUTF8 is UTF-8
	CharUTF8
	//CharISO8859_1 is ISO-8859-1
	CharISO8859_1
	//CharShiftJIS is Shift-JIS
	CharShiftJIS
	//CharEUCJP is EUC-JP
	CharEUCJP
	//CharISO2022JP is ISO-2022-JP
	CharISO2022JP
)

var (
	charEncodeMap = map[CharEncode]string{
		CharUTF8:      "UTF-8",
		CharISO8859_1: "ISO-8859-1",
		CharShiftJIS:  "Shift_JIS",
		CharEUCJP:     "EUC-JP",
		CharISO2022JP: "ISO-2022-JP",
	}
)

//TypeofCharEncode returns CharEncode from string
func TypeofCharEncode(s string) CharEncode {
	for key, value := range charEncodeMap {
		if value == s {
			return key
		}
	}
	return CharUnknown
}

func (e CharEncode) String() string {
	if name, ok := charEncodeMap[e]; ok {
		return name
	}
	return "unknown"

}

//DetectCharEncode returns character encoding
func DetectCharEncode(body []byte) CharEncode {
	det := chardet.NewTextDetector()
	res, err := det.DetectBest(body)
	if err != nil {
		return CharUnknown
	}
	//fmt.Println(res.Charset)
	return TypeofCharEncode(res.Charset)
}

//ToUTF8 returns string with UTF-8 encoding
func ToUTF8(body []byte) string {
	var trans transform.Transformer
	switch DetectCharEncode(body) {
	case CharUTF8, CharISO8859_1:
		return string(body)
	case CharShiftJIS:
		trans = japanese.ShiftJIS.NewDecoder()
	case CharEUCJP:
		trans = japanese.EUCJP.NewDecoder()
	case CharISO2022JP:
		trans = japanese.ISO2022JP.NewDecoder()
	default:
		return ""
	}
	r := transform.NewReader(bytes.NewReader(body), trans)
	buf := new(bytes.Buffer)
	io.Copy(buf, r)
	return buf.String()
}
