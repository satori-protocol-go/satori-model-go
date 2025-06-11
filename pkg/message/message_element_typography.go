package message

import (
	"golang.org/x/net/html"
)

type MessageElementBr struct {
	*noAliasMessageElement
}

func (e *MessageElementBr) Tag() string {
	return "br"
}

func (e *MessageElementBr) Stringify() string {
	return "<br/>"
}

func (e *MessageElementBr) Parse(n *html.Node) (MessageElement, error) {
	return &MessageElementBr{}, nil
}

func (e *MessageElementBr) GetChildren() []MessageElement {
	return nil
}

type MessageElementP struct {
	*noAliasMessageElement
	*ChildrenMessageElement
	*ExtendAttributes
}

func (e *MessageElementP) Tag() string {
	return "p"
}

func (e *MessageElementP) Stringify() string {
	result := e.stringifyAttributes()
	childrenStr := e.stringifyChildren()
	if childrenStr == "" {
		return "<" + e.Tag() + result + "/>"
	}
	return "<" + e.Tag() + result + ">" + childrenStr + "</" + e.Tag() + ">"
}

func (e *MessageElementP) Parse(n *html.Node) (MessageElement, error) {
	attrMap := attrList2MapVal(n.Attr)
	result := &MessageElementP{}
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

type MessageElementMessage struct {
	*noAliasMessageElement
	*ChildrenMessageElement
	*ExtendAttributes
	Id      string // 发 消息的 ID
	Forward bool   // 发 是否为转发消息
}

func (e *MessageElementMessage) Tag() string {
	return "message"
}

func (e *MessageElementMessage) Stringify() string {
	result := ""
	if e.Id != "" {
		result += ` id="` + escape(e.Id, true) + `"`
	}
	if e.Forward {
		result += ` forward`
	}
	result += e.stringifyAttributes()
	childrenStr := e.stringifyChildren()
	if childrenStr == "" {
		return "<" + e.Tag() + result + "/>"
	}
	return "<" + e.Tag() + result + ">" + childrenStr + "</" + e.Tag() + ">"
}

func (e *MessageElementMessage) Parse(n *html.Node) (MessageElement, error) {
	attrMap := attrList2MapVal(n.Attr)
	result := &MessageElementMessage{
		Forward: false,
	}
	if id, ok := attrMap["id"]; ok {
		result.Id = id
	}
	if forwardAttr, ok := attrMap["forward"]; ok {
		result.Forward = forwardAttr == "" || forwardAttr == "true" || forwardAttr == "1"
	}
	for key, value := range attrMap {
		if key != "id" && key != "forward" {
			result.ExtendAttributes = result.AddAttribute(key, value)
		}
	}
	children, err := result.parseChildren(n)
	if err != nil {
		return nil, err
	}
	result.ChildrenMessageElement = children
	return result, nil
}

func init() {
	RegisterParserElement(&MessageElementBr{})
	RegisterParserElement(&MessageElementP{})
	RegisterParserElement(&MessageElementMessage{})
}
