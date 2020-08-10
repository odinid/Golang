package main

import (
	"fmt"
	"strings"
)

func WordCount(s string) map[string]int {
	m := make(map[string]int)
	s = strings.ReplaceAll(s, ",", "")
	s = strings.ReplaceAll(s, ".", "")

	for _, a := range strings.Split(s, " ") {
		v, ok := m[a]

		if ok == true {
			m[a] = v + 1
		} else {
			m[a] = 1
		}
	}
	return m
}

func main() {
	s := "If it looks like a duck, swims like a duck, and quacks like a duck, then it probably is a duck."
	w := WordCount(s)
	fmt.Println(w)
}
