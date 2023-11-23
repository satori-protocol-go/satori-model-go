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

type MessageElementSpl struct {
	*childrenMessageElement
	*noAliasMessageElement
}

func (e *MessageElementSpl) Tag() string {
	return "spl"
}

func (e *MessageElementSpl) Stringify() string {
	return e.stringifyByTag(e.Tag())
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

func init() {
	regsiterParserMulti([]string{"b", "strong"}, func(n *html.Node) (MessageElement, error) {
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
	})
	regsiterParserMulti([]string{"i", "em"}, func(n *html.Node) (MessageElement, error) {
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
	})
	regsiterParserMulti([]string{"s", "ins"}, func(n *html.Node) (MessageElement, error) {
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
	})
	regsiterParser("spl", func(n *html.Node) (MessageElement, error) {
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
	})
	regsiterParser("code", func(n *html.Node) (MessageElement, error) {
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
	})
	regsiterParser("sup", func(n *html.Node) (MessageElement, error) {
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
	})

}
