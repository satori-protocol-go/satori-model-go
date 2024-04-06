package message

import (
	"golang.org/x/net/html"
)

type MessageElementStrong struct {
	*ChildrenMessageElement
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

func (e *MessageElementStrong) Parse(n *html.Node) (MessageElement, error) {
	var children []MessageElement
	err := parseHtmlChildrenNode(n, func(e MessageElement) {
		children = append(children, e)
	})
	if err != nil {
		return nil, err
	}
	return &MessageElementStrong{
		&ChildrenMessageElement{
			Children: children,
		},
	}, nil
}

type MessageElementEm struct {
	*ChildrenMessageElement
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

func (e *MessageElementEm) Parse(n *html.Node) (MessageElement, error) {
	var children []MessageElement
	err := parseHtmlChildrenNode(n, func(e MessageElement) {
		children = append(children, e)
	})
	if err != nil {
		return nil, err
	}
	return &MessageElementEm{
		&ChildrenMessageElement{
			Children: children,
		},
	}, nil
}

type MessageElementIns struct {
	*ChildrenMessageElement
}

func (e *MessageElementIns) Tag() string {
	return "u"
}

func (e *MessageElementIns) Alias() []string {
	return []string{"ins"}
}

func (e *MessageElementIns) Stringify() string {
	return e.stringifyByTag(e.Tag())
}

func (e *MessageElementIns) Parse(n *html.Node) (MessageElement, error) {
	var children []MessageElement
	err := parseHtmlChildrenNode(n, func(e MessageElement) {
		children = append(children, e)
	})
	if err != nil {
		return nil, err
	}
	return &MessageElementIns{
		&ChildrenMessageElement{
			Children: children,
		},
	}, nil
}

type MessageElementDel struct {
	*ChildrenMessageElement
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

func (e *MessageElementDel) Parse(n *html.Node) (MessageElement, error) {
	var children []MessageElement
	err := parseHtmlChildrenNode(n, func(e MessageElement) {
		children = append(children, e)
	})
	if err != nil {
		return nil, err
	}
	return &MessageElementDel{
		ChildrenMessageElement: &ChildrenMessageElement{
			Children: children,
		},
	}, nil
}

type MessageElementSpl struct {
	*noAliasMessageElement
	*ChildrenMessageElement
}

func (e *MessageElementSpl) Tag() string {
	return "spl"
}

func (e *MessageElementSpl) Stringify() string {
	return e.stringifyByTag(e.Tag())
}

func (e *MessageElementSpl) Parse(n *html.Node) (MessageElement, error) {
	var children []MessageElement
	err := parseHtmlChildrenNode(n, func(e MessageElement) {
		children = append(children, e)
	})
	if err != nil {
		return nil, err
	}
	return &MessageElementSpl{
		ChildrenMessageElement: &ChildrenMessageElement{
			Children: children,
		},
	}, nil
}

type MessageElementCode struct {
	*ChildrenMessageElement
	*noAliasMessageElement
}

func (e *MessageElementCode) Tag() string {
	return "code"
}

func (e *MessageElementCode) Stringify() string {
	return e.stringifyByTag(e.Tag())
}

func (e *MessageElementCode) Parse(n *html.Node) (MessageElement, error) {
	var children []MessageElement
	err := parseHtmlChildrenNode(n, func(e MessageElement) {
		children = append(children, e)
	})
	if err != nil {
		return nil, err
	}
	return &MessageElementCode{
		ChildrenMessageElement: &ChildrenMessageElement{
			Children: children,
		},
	}, nil
}

type MessageElementSup struct {
	*ChildrenMessageElement
	*noAliasMessageElement
}

func (e *MessageElementSup) Tag() string {
	return "sup"
}

func (e *MessageElementSup) Stringify() string {
	return e.stringifyByTag(e.Tag())
}

func (e *MessageElementSup) Parse(n *html.Node) (MessageElement, error) {
	var children []MessageElement
	err := parseHtmlChildrenNode(n, func(e MessageElement) {
		children = append(children, e)
	})
	if err != nil {
		return nil, err
	}
	return &MessageElementSup{
		ChildrenMessageElement: &ChildrenMessageElement{
			Children: children,
		},
	}, nil
}

type MessageElementSub struct {
	*ChildrenMessageElement
	*noAliasMessageElement
}

func (e *MessageElementSub) Tag() string {
	return "sub"
}

func (e *MessageElementSub) Stringify() string {
	return e.stringifyByTag(e.Tag())
}

func (e *MessageElementSub) Parse(n *html.Node) (MessageElement, error) {
	var children []MessageElement
	err := parseHtmlChildrenNode(n, func(e MessageElement) {
		children = append(children, e)
	})
	if err != nil {
		return nil, err
	}
	return &MessageElementSub{
		ChildrenMessageElement: &ChildrenMessageElement{
			Children: children,
		},
	}, nil
}

func init() {
	RegsiterParserElement(&MessageElementStrong{})
	RegsiterParserElement(&MessageElementEm{})
	RegsiterParserElement(&MessageElementIns{})
	RegsiterParserElement(&MessageElementDel{})
	RegsiterParserElement(&MessageElementSpl{})
	RegsiterParserElement(&MessageElementCode{})
	RegsiterParserElement(&MessageElementSup{})
	RegsiterParserElement(&MessageElementSub{})
}
