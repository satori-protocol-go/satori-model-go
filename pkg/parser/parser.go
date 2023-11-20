package parser

import (
	"bytes"
	"fmt"
	"strconv"
	"strings"

	"github.com/dezhishen/satori-model-go/pkg/message"
	"golang.org/x/net/html"
)

type messageElemenParser func(n *html.Node) (message.MessageELement, error)

var parserOfTag = make(map[string]messageElemenParser)

func attrList2Map(attrs []html.Attribute) map[string]html.Attribute {
	var result = make(map[string]html.Attribute)
	for _, attr := range attrs {
		result[attr.Key] = attr
	}
	return result
}
func init() {
	parserOfTag["at"] = func(n *html.Node) (message.MessageELement, error) {
		attrMap := attrList2Map(n.Attr)
		return &message.MessageElementAt{
			Id:   attrMap["id"].Val,
			Name: attrMap["name"].Val,
			Role: attrMap["role"].Val,
			Type: attrMap["type"].Val,
		}, nil
	}
	parserOfTag["sharp"] = func(n *html.Node) (message.MessageELement, error) {
		attrMap := attrList2Map(n.Attr)
		return &message.MessageElementSharp{
			Id:   attrMap["id"].Val,
			Name: attrMap["name"].Val,
		}, nil
	}
	parserOfTag["a"] = func(n *html.Node) (message.MessageELement, error) {
		attrMap := attrList2Map(n.Attr)
		return &message.MessageElementA{
			Href: attrMap["href"].Val,
		}, nil
	}
	parserOfTag["img"] = func(n *html.Node) (message.MessageELement, error) {
		attrMap := attrList2Map(n.Attr)
		result := &message.MessageElementImg{
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
	}
	parserOfTag["audio"] = func(n *html.Node) (message.MessageELement, error) {
		attrMap := attrList2Map(n.Attr)
		result := &message.MessageElementAudio{
			Src:     attrMap["src"].Val,
			Cache:   false,
			Timeout: attrMap["timeout"].Val,
		}
		cacheAttr, ok := attrMap["cache"]
		if ok && cacheAttr.Val != "" {
			result.Cache = cacheAttr.Val == "true" || cacheAttr.Val == "1"
		}
		return result, nil
	}
	parserOfTag["video"] = func(n *html.Node) (message.MessageELement, error) {
		attrMap := attrList2Map(n.Attr)
		result := &message.MessageElementVedio{
			Src:     attrMap["src"].Val,
			Cache:   false,
			Timeout: attrMap["timeout"].Val,
		}
		cacheAttr, ok := attrMap["cache"]
		if ok && cacheAttr.Val != "" {
			result.Cache = cacheAttr.Val == "true" || cacheAttr.Val == "1"
		}
		return result, nil
	}
	parserOfTag["file"] = func(n *html.Node) (message.MessageELement, error) {
		attrMap := attrList2Map(n.Attr)
		result := &message.MessageElementFile{
			Src:     attrMap["src"].Val,
			Cache:   false,
			Timeout: attrMap["timeout"].Val,
		}
		cacheAttr, ok := attrMap["cache"]
		if ok && cacheAttr.Val != "" {
			result.Cache = cacheAttr.Val == "true" || cacheAttr.Val == "1"
		}
		return result, nil
	}
}

func parseHtmlNode(n *html.Node, callback func(e message.MessageELement)) error {
	if n.Type == html.ElementNode {
		parserOfTagFunc, ok := parserOfTag[n.Data]
		if ok {
			e, err := parserOfTagFunc(n)
			if err != nil {
				return err
			}
			callback(e)
		}
	} else if n.Type == html.TextNode {
		content := strings.TrimSpace(n.Data)
		if content != "" {
			callback(&message.MessageElementContent{
				Content: content,
			})
		}
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		err := parseHtmlNode(c, callback)
		if err != nil {
			return err
		}
	}
	return nil
}

func Parse(source string) ([]message.MessageELement, error) {
	doc, _ := html.Parse(bytes.NewReader([]byte(source)))
	var result []message.MessageELement
	err := parseHtmlNode(doc, func(e message.MessageELement) {
		if e != nil {
			result = append(result, e)
		}
	})
	if err != nil {
		return nil, err
	}
	return result, nil
}

func Stringify([]message.MessageELement) (string, error) {
	return "", nil
}
