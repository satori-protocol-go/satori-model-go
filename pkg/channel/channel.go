package channel

type ChannelType uint8 // 频道类型

const (
	ChannelTypeText     ChannelType = iota // 文本频道
	ChannelTypeDirect                      // 私聊频道
	ChannelTypeCategory                    // 分类频道
	ChannelTypeVoice                       // 语音频道
)

// 频道
type Channel struct {
	Id       string      `json:"id"`                  // 频道 ID
	Type     ChannelType `json:"type"`                // 频道类型
	Name     string      `json:"name,omitempty"`      // 频道名称
	ParentId string      `json:"parent_id,omitempty"` // 父频道 ID
}

// Channel 分页列表
type ChannelList struct {
	Data []*Channel `json:"data"`           // 数据
	Next string     `json:"next,omitempty"` // 下一页的令牌
}
