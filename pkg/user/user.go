package user

// 用户
type User struct {
	Id     string `json:"id"`               // 用户 ID
	Name   string `json:"name,omitempty"`   // 用户名称
	Nick   string `json:"nick,omitempty"`   // 用户昵称
	Avatar string `json:"avatar,omitempty"` // 用户头像
	IsBot  bool   `json:"is_bot,omitempty"` // 是否为机器人
}

// User 分页列表
type UserList struct {
	Data []User `json:"data"`           // 数据
	Next string `json:"next,omitempty"` // 下一页的令牌
}
