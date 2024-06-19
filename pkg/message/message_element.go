package message

import (
	"golang.org/x/net/html"
)

type MessageElement interface {
	Tag() string
	Stringify() string
	Alias() []string
	GetChildren() []MessageElement
}

type noAliasMessageElement struct {
}

func (e *noAliasMessageElement) Alias() []string {
	return nil
}

type ChildrenMessageElement struct {
	children []MessageElement
}

func (e *ChildrenMessageElement) stringifyChildren() string {
	if e == nil {
		return ""
	}
	if len(e.children) == 0 {
		return ""
	}
	var result string
	for _, e := range e.children {
		result += e.Stringify()
	}
	return result
}

func (e *ChildrenMessageElement) parseChildren(n *html.Node) (*ChildrenMessageElement, error) {
	var children []MessageElement
	err := parseHtmlChildrenNode(n, func(e MessageElement) {
		children = append(children, e)
	})
	if err != nil {
		return nil, err
	}
	result := &ChildrenMessageElement{
		children: children,
	}
	return result, nil
}

func (e *ChildrenMessageElement) GetChildren() []MessageElement {
	if e == nil {
		return nil
	}
	return e.children
}

func (e *ChildrenMessageElement) SetChildren(children []MessageElement) *ChildrenMessageElement {
	e.children = children
	return e
}

type ExtendAttributes struct {
	attributes map[string]string
}

func (e *ExtendAttributes) AddAttribute(key, value string) *ExtendAttributes {
	result := e
	if result == nil {
		result = &ExtendAttributes{
			attributes: make(map[string]string),
		}
	}
	result.attributes[key] = value
	return result
}

func (e *ExtendAttributes) DelAttribute(key string) *ExtendAttributes {
	if e == nil {
		return nil
	}
	delete(e.attributes, key)
	return e
}

func (e *ExtendAttributes) Get(key string) (string, bool) {
	if e == nil {
		return "", false
	}
	v, ok := e.attributes[key]
	return v, ok
}

func (e *ExtendAttributes) stringifyAttributes() string {
	if e == nil || len(e.attributes) == 0 {
		return ""
	}
	var result string
	for k, v := range e.attributes {
		if v == "" {
			result += " " + k
		} else {
			result += " " + k + `="` + escape(v, true) + `"`
		}
	}
	return result
}
