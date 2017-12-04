package mklink

import (
	"github.com/spiegel-im-spiegel/text/decode"
)

//ToUTF8 returns string with UTF-8 encoding
func ToUTF8(body []byte) string {
	if len(body) == 0 {
		return ""
	}
	utf8Text, err := decode.ToUTF8ja(body)
	if err != nil {
		return ""
	}
	return string(utf8Text)
}
