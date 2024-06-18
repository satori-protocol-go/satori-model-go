package guildrole

// 群组角色
type GuildRole struct {
	Id   string `json:"id"`             // 角色 ID
	Name string `json:"name,omitempty"` // 角色名称
}

// GuildRole 分页列表
type GuildRoleList struct {
	Data []*GuildRole `json:"data"`           // 数据
	Next string       `json:"next,omitempty"` // 下一页的令牌
}
