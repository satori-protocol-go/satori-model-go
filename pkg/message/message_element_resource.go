package message

import (
	"fmt"
	"strconv"

	"golang.org/x/net/html"
)

// type ResourceRootMessageElement struct {
// 	Src     string
// 	Cache   bool
// 	Timeout string //ms
// }

// func parseResourceRootMessageElement(attrMap map[string]string) *ResourceRootMessageElement {
// 	result := &ResourceRootMessageElement{
// 		Src:     attrMap["src"],
// 		Cache:   false,
// 		Timeout: attrMap["timeout"],
// 	}
// 	cacheAttr, ok := attrMap["cache"]
// 	if ok {
// 		result.Cache = cacheAttr == "" || cacheAttr == "true" || cacheAttr == "1"
// 	}
// 	return result
// }

// func (e *ResourceRootMessageElement) attrString() string {
// 	if e == nil {
// 		return ""
// 	}
// 	result := ""
// 	if e.Src != "" {
// 		result += ` src="` + e.Src + `"`
// 	}
// 	if e.Cache {
// 		result += ` cache`
// 	}
// 	if e.Timeout != "" {
// 		result += ` timeout="` + e.Timeout + `"`
// 	}
// 	return result
// }

// func (e *ResourceRootMessageElement) stringifyByTag(tag string) string {
// 	return "<" + tag + e.attrString() + " />"
// }

type MessageElementImg struct {
	*noAliasMessageElement
	Src     string
	Cache   bool
	Timeout string //ms
	Width   uint32
	Height  uint32
}

func (e *MessageElementImg) Tag() string {
	return "img"
}

func (e *MessageElementImg) attrString() string {
	if e == nil {
		return ""
	}
	result := ""
	if e.Src != "" {
		result += ` src="` + escape(e.Src, true) + `"`
	}
	if e.Cache {
		result += ` cache`
	}
	if e.Timeout != "" {
		result += ` timeout="` + e.Timeout + `"`
	}
	return result
}

func (e *MessageElementImg) Stringify() string {
	result := "<" + e.Tag()
	attrStr := e.attrString()
	if e.Width > 0 {
		attrStr += fmt.Sprintf(" width=%d", e.Width)
	}
	if e.Height > 0 {
		attrStr += fmt.Sprintf(" height=%d", e.Height)
	}
	result += attrStr
	return result + " />"
}

func (e *MessageElementImg) Parse(n *html.Node) (MessageElement, error) {
	attrMap := attrList2MapVal(n.Attr)
	result := &MessageElementImg{
		Src:     attrMap["src"],
		Cache:   false,
		Timeout: attrMap["timeout"],
	}
	cacheAttr, ok := attrMap["cache"]
	if ok {
		result.Cache = cacheAttr == "" || cacheAttr == "true" || cacheAttr == "1"
	}
	if w, ok := attrMap["width"]; ok {
		width, e := strconv.Atoi(w)
		if e != nil {
			return nil, fmt.Errorf("width[%s] is illegal:%v", w, e)
		}
		result.Width = uint32(width)
	}
	if h, ok := attrMap["height"]; ok {
		height, e := strconv.Atoi(h)
		if e != nil {
			return nil, fmt.Errorf("height[%s] is illegal:%v", h, e)
		}
		result.Height = uint32(height)
	}
	return result, nil
}

type MessageElementAudio struct {
	*noAliasMessageElement
	Src     string
	Cache   bool
	Timeout string //ms
}

func (e *MessageElementAudio) Tag() string {
	return "audio"
}

func (e *MessageElementAudio) attrString() string {
	if e == nil {
		return ""
	}
	result := ""
	if e.Src != "" {
		result += ` src="` + escape(e.Src, true) + `"`
	}
	if e.Cache {
		result += ` cache`
	}
	if e.Timeout != "" {
		result += ` timeout="` + e.Timeout + `"`
	}
	return result
}

func (e *MessageElementAudio) Stringify() string {
	result := "<" + e.Tag()
	result += e.attrString()
	return result + " />"
}

func (e *MessageElementAudio) Parse(n *html.Node) (MessageElement, error) {
	attrMap := attrList2MapVal(n.Attr)
	result := &MessageElementAudio{
		Src:     attrMap["src"],
		Cache:   false,
		Timeout: attrMap["timeout"],
	}
	cacheAttr, ok := attrMap["cache"]
	if ok {
		result.Cache = cacheAttr == "" || cacheAttr == "true" || cacheAttr == "1"
	}
	return result, nil
}

type MessageElementVideo struct {
	*noAliasMessageElement
	Src     string
	Cache   bool
	Timeout string //ms
}

func (e *MessageElementVideo) Tag() string {
	return "video"
}

func (e *MessageElementVideo) attrString() string {
	if e == nil {
		return ""
	}
	result := ""
	if e.Src != "" {
		result += ` src="` + escape(e.Src, true) + `"`
	}
	if e.Cache {
		result += ` cache`
	}
	if e.Timeout != "" {
		result += ` timeout="` + e.Timeout + `"`
	}
	return result
}

func (e *MessageElementVideo) Stringify() string {
	result := "<" + e.Tag()
	result += e.attrString()
	return result + " />"
}

func (e *MessageElementVideo) Parse(n *html.Node) (MessageElement, error) {
	attrMap := attrList2MapVal(n.Attr)
	result := &MessageElementVideo{
		Src:     attrMap["src"],
		Cache:   false,
		Timeout: attrMap["timeout"],
	}
	cacheAttr, ok := attrMap["cache"]
	if ok {
		result.Cache = cacheAttr == "" || cacheAttr == "true" || cacheAttr == "1"
	}
	return result, nil
}

type MessageElementFile struct {
	*noAliasMessageElement
	Src     string
	Cache   bool
	Timeout string //ms
}

func (e *MessageElementFile) Tag() string {
	return "file"
}
func (e *MessageElementFile) attrString() string {
	if e == nil {
		return ""
	}
	result := ""
	if e.Src != "" {
		result += ` src="` + escape(e.Src, true) + `"`
	}
	if e.Cache {
		result += ` cache`
	}
	if e.Timeout != "" {
		result += ` timeout="` + e.Timeout + `"`
	}
	return result
}
func (e *MessageElementFile) Stringify() string {
	result := "<" + e.Tag()
	result += e.attrString()
	return result + " />"
}

func (e *MessageElementFile) Parse(n *html.Node) (MessageElement, error) {
	attrMap := attrList2MapVal(n.Attr)
	result := &MessageElementFile{
		Src:     attrMap["src"],
		Cache:   false,
		Timeout: attrMap["timeout"],
	}
	cacheAttr, ok := attrMap["cache"]
	if ok {
		result.Cache = cacheAttr == "" || cacheAttr == "true" || cacheAttr == "1"
	}
	return result, nil
}

func init() {
	RegsiterParserElement(&MessageElementImg{})
	RegsiterParserElement(&MessageElementAudio{})
	RegsiterParserElement(&MessageElementVideo{})
	RegsiterParserElement(&MessageElementFile{})
}
