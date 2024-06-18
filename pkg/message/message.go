package message

import (
	"github.com/satori-protocol-go/satori-model-go/pkg/channel"
	"github.com/satori-protocol-go/satori-model-go/pkg/guild"
	"github.com/satori-protocol-go/satori-model-go/pkg/guildmember"
	"github.com/satori-protocol-go/satori-model-go/pkg/user"
)

// 消息
type Message struct {
	Id       string                   `json:"id"`                  // 消息 ID
	Content  string                   `json:"content"`             // 消息内容
	Channel  *channel.Channel         `json:"channel,omitempty"`   // 频道对象
	Guild    *guild.Guild             `json:"guild,omitempty"`     // 群组对象
	Member   *guildmember.GuildMember `json:"member,omitempty"`    // 群组成员对象
	User     *user.User               `json:"user,omitempty"`      // 用户对象
	CreateAt int64                    `json:"create_at,omitempty"` // 消息发送的时间戳
	UpdateAt int64                    `json:"update_at,omitempty"` // 消息修改的时间戳
}

// Message 双向分页列表
type MessageBidiList struct {
	Data []Message `json:"data"`           // 数据
	Prev string    `json:"prev,omitempty"` // 上一页的令牌
	Next string    `json:"next,omitempty"` // 下一页的令牌
}

func (m *Message) Decode(elements []MessageElement) error {
	raw := ""
	for _, e := range elements {
		raw += e.Stringify()
	}
	m.Content = raw
	return nil
}

func (m *Message) Encode() ([]MessageElement, error) {
	if m.Content == "" {
		return nil, nil
	}
	return Parse(m.Content)
}

// Select 选取特定的消息元素
func Select(element MessageElement, tag string) []MessageElement {
	var selected []MessageElement

	if element.Tag() == tag {
		selected = append(selected, element)
	}

	selected = selectFromSlide(element.GetChildren(), tag, selected)

	return selected
}

// selectFromSlide 从列表中选取特定的消息元素
func selectFromSlide(elements []MessageElement, tag string, selected []MessageElement) []MessageElement {
	for _, element := range elements {
		if element.Tag() == tag {
			selected = append(selected, element)
		}
	}

	for _, element := range elements {
		selected = append(selected, selectFromSlide(element.GetChildren(), tag, selected)...)
	}

	return selected
}
