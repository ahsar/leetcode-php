// Package jindian provides ...
package jindian

import (
	"fmt"
	"strconv"
	"strings"
	"testing"
)

func Test06(*testing.T) {
	var s string

	//case 1
	s = "aabcccccaaa"

	//case 2
	s = "abbccd"

	r := compressString(s)
	fmt.Println("res", r)
}

func compressString(S string) string {
	var rLen int = len(S)
	var s strings.Builder

	for i, j := 0, 1; i < len(S); i++ {
		j = 1

		tmp := S[i]
		if i < len(S)-1 && S[i+1] == tmp {
			for i < len(S)-1 && S[i+1] == tmp {
				i++
				j++
			}
		}

		s.WriteByte(tmp)
		s.WriteString(strconv.Itoa(j))
	}

	//fmt.Println("str", s.String())
	if len(s.String()) >= rLen {
		return S
	}

	return s.String()
}
