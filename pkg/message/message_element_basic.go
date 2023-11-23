package message

import "golang.org/x/net/html"

type MessageElementText struct {
	*noAliasMessageElement
	Content string
}

func (e *MessageElementText) Tag() string {
	return "text"
}

func (e *MessageElementText) Stringify() string {
	return e.Content
}

type MessageElementAt struct {
	*noAliasMessageElement
	Id   string
	Name string //	收发	目标用户的名称
	Role string //	收发	目标角色
	Type string //	收发	特殊操作，例如 all 表示 @全体成员，here 表示 @在线成员
}

func (e *MessageElementAt) Tag() string {
	return "at"
}

func (e *MessageElementAt) Stringify() string {
	result := "<" + e.Tag()
	if e.Id != "" {
		result += ` id="` + e.Id + `"`
	}
	if e.Name != "" {
		result += ` name="` + e.Name + `"`
	}
	if e.Role != "" {
		result += ` role="` + e.Role + `"`
	}
	if e.Type != "" {
		result += ` type="` + e.Type + `"`
	}
	return result + " />"
}

type MessageElementSharp struct {
	*noAliasMessageElement
	Id   string //收发	目标频道的 ID
	Name string //收发	目标频道的名称
}

func (e *MessageElementSharp) Tag() string {
	return "sharp"
}

func (e *MessageElementSharp) Stringify() string {
	result := "<" + e.Tag()
	if e.Id != "" {
		result += ` id="` + e.Id + `"`
	}
	if e.Name != "" {
		result += ` name="` + e.Name + `"`
	}
	return result + " />"
}

type MessageElementA struct {
	*noAliasMessageElement
	Href string
}

func (e *MessageElementA) Tag() string {
	return "a"
}

func (e *MessageElementA) Stringify() string {
	result := "<" + e.Tag()
	if e.Href != "" {
		result += ` href="` + e.Href + `"`
	}
	return result + " />"
}

func init() {
	regsiterParser("at", func(n *html.Node) (MessageElement, error) {
		attrMap := attrList2Map(n.Attr)
		return &MessageElementAt{
			Id:   attrMap["id"].Val,
			Name: attrMap["name"].Val,
			Role: attrMap["role"].Val,
			Type: attrMap["type"].Val,
		}, nil
	})
	regsiterParser("sharp", func(n *html.Node) (MessageElement, error) {
		attrMap := attrList2Map(n.Attr)
		return &MessageElementSharp{
			Id:   attrMap["id"].Val,
			Name: attrMap["name"].Val,
		}, nil
	})
	regsiterParser("a", func(n *html.Node) (MessageElement, error) {
		attrMap := attrList2Map(n.Attr)
		return &MessageElementA{
			Href: attrMap["href"].Val,
		}, nil
	})
}
