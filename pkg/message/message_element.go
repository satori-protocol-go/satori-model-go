package message

import (
	"fmt"
)

type MessageELement interface {
	Tag() string
	Stringify() string
}

type MessageElementText struct {
	Content string
}

func (e *MessageElementText) Tag() string {
	return "text"
}

func (e *MessageElementText) Stringify() string {
	return e.Content
}

type MessageElementAt struct {
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

type MessageElementImg struct {
	Src     string
	Cache   bool
	Timeout string //ms
	Width   uint32
	Height  uint32
}

func (e *MessageElementImg) Tag() string {
	return "img"
}
func (e *MessageElementImg) Stringify() string {
	result := "<" + e.Tag()
	if e.Src != "" {
		result += ` src="` + e.Src + `"`
	}
	if e.Cache {
		result += ` cache`
	}
	if e.Timeout != "" {
		result += ` timeout="` + e.Timeout + `"`
	}
	if e.Width > 0 {
		result += fmt.Sprintf(" width=%d", e.Width)
	}
	if e.Height > 0 {
		result += fmt.Sprintf(" height=%d", e.Height)
	}
	return result + " />"
}

type MessageElementAudio struct {
	Src     string
	Cache   bool
	Timeout string //ms
}

func (e *MessageElementAudio) Tag() string {
	return "audio"
}

func (e *MessageElementAudio) Stringify() string {
	result := "<" + e.Tag()
	if e.Src != "" {
		result += ` src="` + e.Src + `"`
	}
	if e.Cache {
		result += ` cache`
	}
	if e.Timeout != "" {
		result += ` timeout="` + e.Timeout + `"`
	}
	return result + " />"
}

type MessageElementVedio struct {
	Src     string
	Cache   bool
	Timeout string //ms
}

func (e *MessageElementVedio) Tag() string {
	return "vedio"
}

func (e *MessageElementVedio) Stringify() string {
	result := "<" + e.Tag()
	if e.Src != "" {
		result += ` src="` + e.Src + `"`
	}
	if e.Cache {
		result += ` cache`
	}
	if e.Timeout != "" {
		result += ` timeout="` + e.Timeout + `"`
	}
	return result + " />"
}

type MessageElementFile struct {
	Src     string
	Cache   bool
	Timeout string //ms
}

func (e *MessageElementFile) Tag() string {
	return "file"
}

func (e *MessageElementFile) Stringify() string {
	result := "<" + e.Tag()
	if e.Src != "" {
		result += ` src="` + e.Src + `"`
	}
	if e.Cache {
		result += ` cache`
	}
	if e.Timeout != "" {
		result += ` timeout="` + e.Timeout + `"`
	}
	return result + " />"
}
