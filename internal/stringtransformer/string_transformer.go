package stringtransformer

import (
	"strconv"
	"strings"
	"unicode"
)

func Transform(str []byte) strings.Builder {
	var strAfterTransform strings.Builder
	strLen := len(str) - 1

	for i := 0; i <= strLen; i++ {
		if unicode.IsDigit(rune(str[i])) {
			continue
		}
		if i != strLen {
			if unicode.IsDigit(rune(str[i+1])) {
				res, _ := strconv.Atoi(string(str[i+1]))
				for j := 0; res > j; j++ {
					strAfterTransform.WriteByte(str[i])
				}
			} else {
				strAfterTransform.WriteByte(str[i])
			}
		} else {
			strAfterTransform.WriteByte(str[i])
		}
	}

	return strAfterTransform
}
