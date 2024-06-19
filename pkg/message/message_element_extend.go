package message

import "golang.org/x/net/html"

type MessageElementExtend struct {
	*noAliasMessageElement
	*ChildrenMessageElement
	*ExtendAttributes
	tag string
}

func (e *MessageElementExtend) Tag() string {
	return e.tag
}

func (e *MessageElementExtend) Stringify() string {
	result := ""
	result += e.stringifyAttributes()
	childrenStr := e.stringifyChildren()
	if childrenStr == "" {
		return `<` + e.Tag() + result + `/>`
	}
	return `<` + e.Tag() + result + `>` + childrenStr + `</` + e.Tag() + `>`
}

func (e *MessageElementExtend) Parse(n *html.Node) (MessageElement, error) {
	attrMap := attrList2MapVal(n.Attr)
	result := &MessageElementExtend{
		tag: n.Data,
	}
	for key, value := range attrMap {
		result.ExtendAttributes = result.AddAttribute(key, value)
	}
	children, err := result.parseChildren(n)
	if err != nil {
		return nil, err
	}
	result.ChildrenMessageElement = children
	return result, nil
}

func NewMessageElementExtend(tag string, attrs map[string]string, children ...MessageElement) MessageElement {
	result := &MessageElementExtend{
		tag: tag,
	}
	for key, value := range attrs {
		result.ExtendAttributes = result.AddAttribute(key, value)
	}
	result.ChildrenMessageElement = &ChildrenMessageElement{
		children: children,
	}
	return result
}

var ExtendParser = &MessageElementExtend{}
