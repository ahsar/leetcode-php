// Package jindian provides ...
package jindian

import (
	"fmt"
	"strings"
	"testing"
)

func Test03(*testing.T) {
	var (
		s      string
		length int
	)

	// case 1
	s = "Mr John Smith    "
	length = 13

	// case 2
	//s = "               "
	//length = 5

	r := replaceSpaces(s, length)
	fmt.Println("res:", r)
}

/**
 * 正序
 */
func replaceSpaces(S string, length int) string {
	var (
		s   strings.Builder
		tmp string
	)

	for i := 0; i < length; i++ {
		tmp = string(S[i])
		if tmp == " " {
			tmp = "%20"
		}
		s.WriteString(tmp)
	}

	return s.String()
}

/**
 * 倒叙
 */
func replaceSpaces1(S string, length int) string {
	var res []string = make([]string, length)

	var i int = len(S) - 1
	for i >= length && string(S[i]) == " " {
		i--
	}

	var tmp string
	for ; i >= 0; i-- {
		tmp = string(S[i])
		if tmp == " " {
			tmp = "%20"
		}

		res[i] = tmp
	}
	return strings.Join(res, "")
}
