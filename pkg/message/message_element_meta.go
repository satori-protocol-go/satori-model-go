package message

type MessageElementQuote struct {
	*noAliasMessageElement
	*childrenMessageElement
}

func (e *MessageElementQuote) Tag() string {
	return "quote"
}

func (e *MessageElementQuote) Stringify() string {
	return e.stringifyByTag(e.Tag())
}

type MessageElementAuthor struct {
	*noAliasMessageElement
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
		result += ` id="` + e.Id + `"`
	}
	if e.Name != "" {
		result += ` name="` + e.Name + `"`
	}
	if e.Avatar != "" {
		result += ` avatar="` + e.Avatar + `"`
	}
	return result + " />"
}
