package guild

// 群组
type Guild struct {
	Id     string `json:"id"`               // 群组 ID
	Name   string `json:"name,omitempty"`   // 群组名称
	Avatar string `json:"avatar,omitempty"` // 群组头像
}

// Guild 分页列表
type GuildList struct {
	Data []*Guild `json:"data"`           // 数据
	Next string   `json:"next,omitempty"` // 下一页的令牌
}
