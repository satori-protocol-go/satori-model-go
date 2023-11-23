package message

import (
	"fmt"
	"strconv"

	"golang.org/x/net/html"
)

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
func (e *MessageElementImg) Stringify() string {
	result := "<" + e.Tag()
	if e.Src != "" {
		result += ` src="` + e.Src + `"`
	}
	if e.Cache {
		result += ` cache`
	}
	if e.Timeout != "" {
		result += ` timeout="` + e.Timeout + `"`
	}
	if e.Width > 0 {
		result += fmt.Sprintf(" width=%d", e.Width)
	}
	if e.Height > 0 {
		result += fmt.Sprintf(" height=%d", e.Height)
	}
	return result + " />"
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

func (e *MessageElementAudio) Stringify() string {
	result := "<" + e.Tag()
	if e.Src != "" {
		result += ` src="` + e.Src + `"`
	}
	if e.Cache {
		result += ` cache`
	}
	if e.Timeout != "" {
		result += ` timeout="` + e.Timeout + `"`
	}
	return result + " />"
}

type MessageElementVedio struct {
	*noAliasMessageElement
	Src     string
	Cache   bool
	Timeout string //ms
}

func (e *MessageElementVedio) Tag() string {
	return "vedio"
}

func (e *MessageElementVedio) Stringify() string {
	result := "<" + e.Tag()
	if e.Src != "" {
		result += ` src="` + e.Src + `"`
	}
	if e.Cache {
		result += ` cache`
	}
	if e.Timeout != "" {
		result += ` timeout="` + e.Timeout + `"`
	}
	return result + " />"
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

func (e *MessageElementFile) Stringify() string {
	result := "<" + e.Tag()
	if e.Src != "" {
		result += ` src="` + e.Src + `"`
	}
	if e.Cache {
		result += ` cache`
	}
	if e.Timeout != "" {
		result += ` timeout="` + e.Timeout + `"`
	}
	return result + " />"
}

func init() {
	regsiterParser("img", func(n *html.Node) (MessageElement, error) {
		attrMap := attrList2Map(n.Attr)
		result := &MessageElementImg{
			Src:     attrMap["src"].Val,
			Cache:   false,
			Timeout: attrMap["timeout"].Val,
		}
		cacheAttr, ok := attrMap["cache"]
		if ok && cacheAttr.Val != "" {
			result.Cache = cacheAttr.Val == "true" || cacheAttr.Val == "1"
		}
		if w, ok := attrMap["width"]; ok {
			width, e := strconv.Atoi(w.Val)
			if e != nil {
				return nil, fmt.Errorf("width[%s] is illegal:%v", w.Val, e)
			}
			result.Width = uint32(width)
		}
		if h, ok := attrMap["height"]; ok {
			height, e := strconv.Atoi(h.Val)
			if e != nil {
				return nil, fmt.Errorf("height[%s] is illegal:%v", h.Val, e)
			}
			result.Height = uint32(height)
		}
		return result, nil
	})
	regsiterParser("audio", func(n *html.Node) (MessageElement, error) {
		attrMap := attrList2Map(n.Attr)
		result := &MessageElementAudio{
			Src:     attrMap["src"].Val,
			Cache:   false,
			Timeout: attrMap["timeout"].Val,
		}
		cacheAttr, ok := attrMap["cache"]
		if ok && cacheAttr.Val != "" {
			result.Cache = cacheAttr.Val == "true" || cacheAttr.Val == "1"
		}
		return result, nil
	})
	regsiterParser("video", func(n *html.Node) (MessageElement, error) {
		attrMap := attrList2Map(n.Attr)
		result := &MessageElementVedio{
			Src:     attrMap["src"].Val,
			Cache:   false,
			Timeout: attrMap["timeout"].Val,
		}
		cacheAttr, ok := attrMap["cache"]
		if ok && cacheAttr.Val != "" {
			result.Cache = cacheAttr.Val == "true" || cacheAttr.Val == "1"
		}
		return result, nil
	})
	regsiterParser("file", func(n *html.Node) (MessageElement, error) {
		attrMap := attrList2Map(n.Attr)
		result := &MessageElementFile{
			Src:     attrMap["src"].Val,
			Cache:   false,
			Timeout: attrMap["timeout"].Val,
		}
		cacheAttr, ok := attrMap["cache"]
		if ok && cacheAttr.Val != "" {
			result.Cache = cacheAttr.Val == "true" || cacheAttr.Val == "1"
		}
		return result, nil
	})
}
