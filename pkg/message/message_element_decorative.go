package message

import (
	"golang.org/x/net/html"
)

type MessageElementStrong struct {
	*childrenMessageElement
}

func (e *MessageElementStrong) Tag() string {
	return "b"
}
func (e *MessageElementStrong) Alias() []string {
	return []string{"strong"}
}

func (e *MessageElementStrong) Stringify() string {
	return e.stringifyByTag(e.Tag())
}

func (e *MessageElementStrong) parse(n *html.Node) (MessageElement, error) {
	var children []MessageElement
	err := parseHtmlChildrenNode(n, func(e MessageElement) {
		children = append(children, e)
	})
	if err != nil {
		return nil, err
	}
	return &MessageElementStrong{
		&childrenMessageElement{
			Children: children,
		},
	}, nil
}

type MessageElementEm struct {
	*childrenMessageElement
}

func (e *MessageElementEm) Tag() string {
	return "i"
}
func (e *MessageElementEm) Alias() []string {
	return []string{"em"}
}

func (e *MessageElementEm) Stringify() string {
	return e.stringifyByTag(e.Tag())
}

func (e *MessageElementEm) parse(n *html.Node) (MessageElement, error) {
	var children []MessageElement
	err := parseHtmlChildrenNode(n, func(e MessageElement) {
		children = append(children, e)
	})
	if err != nil {
		return nil, err
	}
	return &MessageElementEm{
		&childrenMessageElement{
			Children: children,
		},
	}, nil
}

type MessageElementIns struct {
	*childrenMessageElement
}

func (e *MessageElementIns) Tag() string {
	return "s"
}

func (e *MessageElementIns) Alias() []string {
	return []string{"ins"}
}

func (e *MessageElementIns) Stringify() string {
	return e.stringifyByTag(e.Tag())
}

func (e *MessageElementIns) parse(n *html.Node) (MessageElement, error) {
	var children []MessageElement
	err := parseHtmlChildrenNode(n, func(e MessageElement) {
		children = append(children, e)
	})
	if err != nil {
		return nil, err
	}
	return &MessageElementIns{
		&childrenMessageElement{
			Children: children,
		},
	}, nil
}

type MessageElementDel struct {
	*childrenMessageElement
}

func (e *MessageElementDel) Tag() string {
	return "s"
}

func (e *MessageElementDel) Alias() []string {
	return []string{"del"}
}

func (e *MessageElementDel) Stringify() string {
	return e.stringifyByTag(e.Tag())
}

func (e *MessageElementDel) parse(n *html.Node) (MessageElement, error) {
	var children []MessageElement
	err := parseHtmlChildrenNode(n, func(e MessageElement) {
		children = append(children, e)
	})
	if err != nil {
		return nil, err
	}
	return &MessageElementDel{
		childrenMessageElement: &childrenMessageElement{
			Children: children,
		},
	}, nil
}

type MessageElementSpl struct {
	*noAliasMessageElement
	*childrenMessageElement
}

func (e *MessageElementSpl) Tag() string {
	return "spl"
}

func (e *MessageElementSpl) Stringify() string {
	return e.stringifyByTag(e.Tag())
}

func (e *MessageElementSpl) parse(n *html.Node) (MessageElement, error) {
	var children []MessageElement
	err := parseHtmlChildrenNode(n, func(e MessageElement) {
		children = append(children, e)
	})
	if err != nil {
		return nil, err
	}
	return &MessageElementSpl{
		childrenMessageElement: &childrenMessageElement{
			Children: children,
		},
	}, nil
}

type MessageElementCode struct {
	*childrenMessageElement
	*noAliasMessageElement
}

func (e *MessageElementCode) Tag() string {
	return "code"
}

func (e *MessageElementCode) Stringify() string {
	return e.stringifyByTag(e.Tag())
}

func (e *MessageElementCode) parse(n *html.Node) (MessageElement, error) {
	var children []MessageElement
	err := parseHtmlChildrenNode(n, func(e MessageElement) {
		children = append(children, e)
	})
	if err != nil {
		return nil, err
	}
	return &MessageElementCode{
		childrenMessageElement: &childrenMessageElement{
			Children: children,
		},
	}, nil
}

type MessageElementSup struct {
	*childrenMessageElement
	*noAliasMessageElement
}

func (e *MessageElementSup) Tag() string {
	return "sup"
}

func (e *MessageElementSup) Stringify() string {
	return e.stringifyByTag(e.Tag())
}

func (e *MessageElementSup) parse(n *html.Node) (MessageElement, error) {
	var children []MessageElement
	err := parseHtmlChildrenNode(n, func(e MessageElement) {
		children = append(children, e)
	})
	if err != nil {
		return nil, err
	}
	return &MessageElementSup{
		childrenMessageElement: &childrenMessageElement{
			Children: children,
		},
	}, nil
}

type MessageElementSub struct {
	*childrenMessageElement
	*noAliasMessageElement
}

func (e *MessageElementSub) Tag() string {
	return "sub"
}

func (e *MessageElementSub) Stringify() string {
	return e.stringifyByTag(e.Tag())
}

func (e *MessageElementSub) parse(n *html.Node) (MessageElement, error) {
	var children []MessageElement
	err := parseHtmlChildrenNode(n, func(e MessageElement) {
		children = append(children, e)
	})
	if err != nil {
		return nil, err
	}
	return &MessageElementSub{
		childrenMessageElement: &childrenMessageElement{
			Children: children,
		},
	}, nil
}

func init() {
	regsiterParserElement(&MessageElementStrong{})
	regsiterParserElement(&MessageElementEm{})
	regsiterParserElement(&MessageElementIns{})
	regsiterParserElement(&MessageElementDel{})
	regsiterParserElement(&MessageElementSpl{})
	regsiterParserElement(&MessageElementCode{})
	regsiterParserElement(&MessageElementSup{})
	regsiterParserElement(&MessageElementSub{})
}
