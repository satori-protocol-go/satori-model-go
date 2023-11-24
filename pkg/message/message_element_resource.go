package message

import (
	"fmt"
	"strconv"

	"golang.org/x/net/html"
)

type resourceRootMessageElement struct {
	Src     string
	Cache   bool
	Timeout string //ms
}

func parseResourceRootMessageElement(attrMap map[string]string) *resourceRootMessageElement {
	result := &resourceRootMessageElement{
		Src:     attrMap["src"],
		Cache:   false,
		Timeout: attrMap["timeout"],
	}
	cacheAttr, ok := attrMap["cache"]
	if ok && cacheAttr != "" {
		result.Cache = cacheAttr == "true" || cacheAttr == "1"
	}
	return result
}

func (e *resourceRootMessageElement) attrString() string {
	result := ""
	if e.Src != "" {
		result += ` src="` + e.Src + `"`
	}
	if e.Cache {
		result += ` cache`
	}
	if e.Timeout != "" {
		result += ` timeout="` + e.Timeout + `"`
	}
	return result
}

func (e *resourceRootMessageElement) stringifyByTag(tag string) string {
	return "<" + tag + e.attrString() + "/>"
}

type MessageElementImg struct {
	*noAliasMessageElement
	*resourceRootMessageElement
	Width  uint32
	Height uint32
}

func (e *MessageElementImg) Tag() string {
	return "img"
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
	return result + " />"
}

func (e *MessageElementImg) parse(n *html.Node) (MessageElement, error) {
	attrMap := attrList2MapVal(n.Attr)
	root := parseResourceRootMessageElement(attrMap)
	result := &MessageElementImg{
		resourceRootMessageElement: root,
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
	*resourceRootMessageElement
}

func (e *MessageElementAudio) Tag() string {
	return "audio"
}

func (e *MessageElementAudio) Stringify() string {
	return e.stringifyByTag(e.Tag())
}

func (e *MessageElementAudio) parse(n *html.Node) (MessageElement, error) {
	attrMap := attrList2MapVal(n.Attr)
	return &MessageElementAudio{
		resourceRootMessageElement: parseResourceRootMessageElement(attrMap),
	}, nil
}

type MessageElementVedio struct {
	*noAliasMessageElement
	*resourceRootMessageElement
}

func (e *MessageElementVedio) Tag() string {
	return "vedio"
}

func (e *MessageElementVedio) Stringify() string {
	return e.stringifyByTag(e.Tag())
}

func (e *MessageElementVedio) parse(n *html.Node) (MessageElement, error) {
	attrMap := attrList2MapVal(n.Attr)
	return &MessageElementVedio{
		resourceRootMessageElement: parseResourceRootMessageElement(attrMap),
	}, nil
}

type MessageElementFile struct {
	*noAliasMessageElement
	*resourceRootMessageElement
}

func (e *MessageElementFile) Tag() string {
	return "file"
}

func (e *MessageElementFile) Stringify() string {
	return e.stringifyByTag(e.Tag())
}

func (e *MessageElementFile) parse(n *html.Node) (MessageElement, error) {
	attrMap := attrList2MapVal(n.Attr)
	return &MessageElementFile{
		resourceRootMessageElement: parseResourceRootMessageElement(attrMap),
	}, nil
}

func init() {
	regsiterParserElement(&MessageElementImg{})
	regsiterParserElement(&MessageElementAudio{})
	regsiterParserElement(&MessageElementVedio{})
	regsiterParserElement(&MessageElementFile{})
}
