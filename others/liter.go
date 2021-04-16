// Package main provides ...
package main

import (
	"fmt"
)

type stu struct {
	Name string
	Age  int
}

func parse_stu() {
	m := make(map[string]*stu)
	stus := []stu{
		{Name: "zhou", Age: 14},
		{Name: "wang", Age: 15},
		{Name: "li", Age: 17},
	}

	for _, stu := range stus {
		m[stu.Name] = &stu
	}
	fmt.Println(stus)
}

func main() {
	parse_stu()
}
