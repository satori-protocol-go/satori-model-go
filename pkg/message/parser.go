package message

import (
	"bytes"
	"strings"

	"golang.org/x/net/html"
)

type messageElemenParser func(n *html.Node) (MessageELement, error)

var _parserOfTag = make(map[string]messageElemenParser)

func attrList2Map(attrs []html.Attribute) map[string]html.Attribute {
	var result = make(map[string]html.Attribute)
	for _, attr := range attrs {
		result[attr.Key] = attr
	}
	return result
}

func regsiterParser(tag string, parserFunc messageElemenParser) {
	_parserOfTag[tag] = parserFunc
}
func regsiterParserMulti(tags []string, parserFunc messageElemenParser) {
	for _, tag := range tags {
		_parserOfTag[tag] = parserFunc
	}
}

func init() {
}

func parseHtmlNode(n *html.Node, callback func(e MessageELement)) error {
	parsed := false
	if n.Type == html.ElementNode {
		var parserOfTagFunc messageElemenParser
		parserOfTagFunc, parsed = _parserOfTag[n.Data]
		if parsed {
			e, err := parserOfTagFunc(n)
			if err != nil {
				return err
			}
			callback(e)
		}
	} else if n.Type == html.TextNode {
		content := strings.TrimSpace(n.Data)
		if content != "" {
			callback(&MessageElementText{
				Content: content,
			})
		}
		parsed = true
	}
	if !parsed {
		parseHtmlChildrenNode(n, callback)
	}
	return nil
}
func parseHtmlChildrenNode(n *html.Node, callback func(e MessageELement)) error {
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		err := parseHtmlNode(c, callback)
		if err != nil {
			return err
		}
	}
	return nil
}

func Parse(source string) ([]MessageELement, error) {
	doc, _ := html.Parse(bytes.NewReader([]byte(source)))
	var result []MessageELement
	err := parseHtmlNode(doc, func(e MessageELement) {
		if e != nil {
			result = append(result, e)
		}
	})
	if err != nil {
		return nil, err
	}
	return result, nil
}

func Stringify([]MessageELement) (string, error) {
	return "", nil
}
