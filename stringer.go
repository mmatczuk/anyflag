package anyflag

import (
	"fmt"
	"strings"
)

type ints interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 | ~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64
}

// StringerParser returns parse function for enums processed with stringer tool.
func StringerParser[T ints](index []uint8, name string) func(val string) (T, error) {
	return func(val string) (T, error) {
		var idx uint8 = 0
		for i, p := range index[1:] {
			if strings.EqualFold(name[idx:p], val) {
				return T(i), nil
			}
			idx = p
		}
		return T(0), fmt.Errorf("invalid value: %q", val)
	}
}
