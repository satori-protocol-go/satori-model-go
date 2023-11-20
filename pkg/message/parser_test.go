package message

import (
	"fmt"
	"testing"
)

func TestParse(t *testing.T) {
	elements, _ := Parse(
		`我是纯文本1<at id="1" name="123456" />我是纯文本2<at id="2" name="123456" />`)
	s := ""
	for _, e := range elements {
		s += e.Stringify()
	}
	fmt.Println(s)
}

func TestStringify(t *testing.T) {
	elements := []MessageELement{
		&MessageElementText{
			Content: "我是纯文本",
		},
		&MessageElementAt{
			Id:   "test",
			Name: "test",
			Role: "test",
			Type: "",
		},
		&MessageElementAudio{
			Src:     "https://127.0.0.1/mp3.mp3",
			Cache:   false,
			Timeout: "",
		},
	}
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
