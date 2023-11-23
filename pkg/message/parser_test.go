package message

import (
	"fmt"
	"testing"
)

func TestParse(t *testing.T) {
	elements, _ := Parse(
		`我是纯文本<b>我是加粗文本<b>套娃<b>套娃中的套娃</b></b>我是123</b>`)
	s := ""
	for _, e := range elements {
		s += e.Stringify()
	}
	fmt.Println(s)
}

func TestStringify(t *testing.T) {
	elements, _ := Parse(
		`我是纯文本<b>我是加粗文本<b>套娃<b>套娃中的套娃</b></b>我是123</b>`)
	s := ""
	for _, e := range elements {
		s += e.Stringify()
	}
	fmt.Println(s)
	var newS string
	newElements, _ := Parse(s)
	for _, e := range newElements {
		newS += e.Stringify()
	}
	fmt.Println(newS)
}
