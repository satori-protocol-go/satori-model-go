package message

import (
	"fmt"

	"golang.org/x/net/html"
)

type MessageElmentBr struct {
	*noAliasMessageElement
}

func (e *MessageElmentBr) Tag() string {
	return "br"
}

func (e *MessageElmentBr) Stringify() string {
	return fmt.Sprintln()
}

func (e *MessageElmentBr) Parse(n *html.Node) (MessageElement, error) {
	return &MessageElmentBr{}, nil
}

type MessageElmentP struct {
	*noAliasMessageElement
	*childrenMessageElement
}

func (e *MessageElmentP) Tag() string {
	return "p"
}

func (e *MessageElmentP) Stringify() string {
	return e.stringifyByTag(e.Tag())
}

func (e *MessageElmentP) Parse(n *html.Node) (MessageElement, error) {
	var children []MessageElement
	err := parseHtmlChildrenNode(n, func(e MessageElement) {
		children = append(children, e)
	})
	if err != nil {
		return nil, err
	}
	return &MessageElmentP{
		childrenMessageElement: &childrenMessageElement{
			Children: children,
		},
	}, nil
}

type MessageElementMessage struct {
	*noAliasMessageElement
	*childrenMessageElement
}

func (e *MessageElementMessage) Tag() string {
	return "message"
}

func (e *MessageElementMessage) Stringify() string {
	return e.stringifyByTag(e.Tag())
}

func (e *MessageElementMessage) Parse(n *html.Node) (MessageElement, error) {
	var children []MessageElement
	err := parseHtmlChildrenNode(n, func(e MessageElement) {
		children = append(children, e)
	})
	if err != nil {
		return nil, err
	}
	return &MessageElementMessage{
		childrenMessageElement: &childrenMessageElement{
			Children: children,
		},
	}, nil
}

func init() {
	RegsiterParserElement(&MessageElmentBr{})
	RegsiterParserElement(&MessageElmentP{})
	RegsiterParserElement(&MessageElementMessage{})
}
