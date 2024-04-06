package message

type MessageElement interface {
	Tag() string
	Stringify() string
	Alias() []string
}

type noAliasMessageElement struct {
}

func (e *noAliasMessageElement) Alias() []string {
	return nil
}

type ChildrenMessageElement struct {
	Children []MessageElement
}

func (e *ChildrenMessageElement) stringifyChildren() string {
	if e == nil {
		return ""
	}
	if len(e.Children) == 0 {
		return ""
	}
	var result string
	for _, e := range e.Children {
		result += e.Stringify()
	}
	return result
}

func (e *ChildrenMessageElement) stringifyByTag(tag string) string {
	if e == nil || len(e.Children) == 0 {
		return "<" + tag + " />"
	}
	return "<" + tag + ">" + e.stringifyChildren() + "</" + tag + ">"
}
