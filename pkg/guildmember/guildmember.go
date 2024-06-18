package guildmember

import "github.com/satori-protocol-go/satori-model-go/pkg/user"

// 群组成员
type GuildMember struct {
	User     *user.User `json:"user,omitempty"`      // 用户对象
	Nick     string     `json:"nick,omitempty"`      // 用户在群组中的名称
	Avatar   string     `json:"avatar,omitempty"`    // 用户在群组中的头像
	JoinedAt int64      `json:"joined_at,omitempty"` // 加入时间
}

// GuildMember 分页列表
type GuildMemberList struct {
	Data []*GuildMember `json:"data"`           // 数据
	Next string         `json:"next,omitempty"` // 下一页的令牌
}
