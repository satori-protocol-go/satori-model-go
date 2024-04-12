package message

import "golang.org/x/net/html"

type MessageElementQuote struct {
	*noAliasMessageElement
	*ChildrenMessageElement
}

func (e *MessageElementQuote) Tag() string {
	return "quote"
}

func (e *MessageElementQuote) Stringify() string {
	return e.stringifyByTag(e.Tag())
}

func (e *MessageElementQuote) Parse(n *html.Node) (MessageElement, error) {
	var children []MessageElement
	err := parseHtmlChildrenNode(n, func(e MessageElement) {
		children = append(children, e)
	})
	if err != nil {
		return nil, err
	}
	return &MessageElementQuote{
		ChildrenMessageElement: &ChildrenMessageElement{
			Children: children,
		},
	}, nil
}

type MessageElementAuthor struct {
	*noAliasMessageElement
	*ChildrenMessageElement
	Id     string
	Name   string
	Avatar string
}

func (e *MessageElementAuthor) Tag() string {
	return "author"
}

func (e *MessageElementAuthor) Stringify() string {
	result := "<" + e.Tag()
	if e.Id != "" {
		result += ` id="` + escape(e.Id, true) + `"`
	}
	if e.Name != "" {
		result += ` name="` + escape(e.Name, true) + `"`
	}
	if e.Avatar != "" {
		result += ` avatar="` + escape(e.Avatar, true) + `"`
	}
	// result += ">"
	childrenStr := e.stringifyChildren()
	if childrenStr == "" {
		return result + `/>`
	}
	return result + `>` + childrenStr + `</` + e.Tag() + `>`
}

func (e *MessageElementAuthor) Parse(n *html.Node) (MessageElement, error) {
	attrMap := attrList2MapVal(n.Attr)
	result := &MessageElementAuthor{
		Id:     attrMap["id"],
		Name:   attrMap["name"],
		Avatar: attrMap["avatar"],
	}
	err := parseHtmlChildrenNode(n, func(e MessageElement) {
		result.Children = append(result.Children, e)
	})
	if err != nil {
		return nil, err
	}
	return result, nil
}
func init() {
	RegsiterParserElement(&MessageElementQuote{})
	RegsiterParserElement(&MessageElementAuthor{})
}
