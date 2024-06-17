package login

import "github.com/satori-protocol-go/satori-model-go/pkg/user"

type LoginStatus uint8

const (
	StatusOffline    LoginStatus = iota // 离线
	StatusOnline                        // 在线
	StatusConnect                       // 连接中
	StatusDisconnect                    // 断开连接
	StatusReconnect                     // 重新连接
)

type Login struct {
	User      *user.User  `json:"user,omitempty"`     // 用户对象
	SelfId    string      `json:"self_id,omitempty"`  // 平台账号
	Platform  string      `json:"platform,omitempty"` // 平台名称
	Status    LoginStatus `json:"status"`             // 登录状态
	Features  []string    `json:"features"`           // 平台特性 列表
	ProxyUrls []string    `json:"proxy_urls"`         // 代理路由 列表
}
