package interaction

// 交互指令
type Argv struct {
	Name      string                 `json:"name"`      // 指令名称
	Arguments []interface{}          `json:"arguments"` // 参数
	Options   map[string]interface{} `json:"options"`   // 选项
}

// 交互按钮
type Button struct {
	Id string `json:"id"` // 按钮 ID
}
