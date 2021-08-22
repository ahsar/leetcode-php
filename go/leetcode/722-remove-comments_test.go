package leetcode

import (
	"fmt"
	"testing"
)

func Test722(*testing.T) {
	var source []string

	// case 1
	source = []string{"/*Test program */", "int main()", "{ ", "  // variable declaration ", "int a, b, c;", "/* This is a test", "   multiline  ", "   comment for ", "   testing */", "a = b + c;", "}"}

	// case 2
	source = []string{"a/*comment", "line", "more_comment*/b"}

	// case 3
	source = []string{"struct Node{", "abcd/*/bdeclarenmembers;/**/", "    int size;", "    /**/int val;", "};"}

	// case 4
	//source = []string{"main() {", "/* here is commments", "  // still comments */", "   double s = 33;", "   cout << s;", "}"}

	// case 5
	source = []string{"void func(int k) {", "// this function does nothing /*", "   k = k*2/4;", "   k = k/2;*/", "}"}

	// case 6
	//source = []string{"a//*b//*c", "blank", "d/*/e*//f"}

	r := removeComments(source)
	fmt.Println("res:")
	for _, v := range r {
		fmt.Printf("%s\n", v)
	}
	/*/ declare members;/**/
}

func removeComments(source []string) []string {
	var (
		r    []string
		flag bool
		code string
		char string
	)

	for _, v := range source {
		if !flag {
			code = ""
		}
		char = ""
		fmt.Println("start:", v)
		for i := 0; i < len(v); i++ {
			char = string(v[i])

			// 行注释
			if i > 0 && char == "/" && string(v[i-1]) == "/" && !flag {
				//fmt.Println("block", i, code, flag)
				//code = code[:len(code)-1]
				break
				//continue
			}

			// 块注释结束 */
			if i > 0 && char == "/" && string(v[i-1]) == "*" && flag {
				flag = false
				i++
				//continue
			}

			// 块注释开始 /*
			if i > 0 && char == "*" && string(v[i-1]) == "/" && !flag {
				code = code[:len(code)-1]
				flag = true
				i++
				//fmt.Println("code", code, flag)
			}

			if !flag {
				code += char
			}
		}
		if len(code) == 0 {
			continue
		}
		if !flag {
			r = append(r, code)
		}
	}
	return r
}
