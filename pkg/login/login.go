package login

import "github.com/satori-protocol-go/satori-model-go/pkg/user"

type LoginStatus uint8

const (
	StatusOffline    LoginStatus = iota // 离线
	StatusOnline                        // 在线
	StatusConnect                       // 正在连接
	StatusDisconnect                    // 正在断开连接
	StatusReconnect                     // 正在重新连接
)

type Login struct {
	Sn       int64       `json:"sn"`                 // 序列号
	Platform string      `json:"platform,omitempty"` // 平台名称
	User     *user.User  `json:"user,omitempty"`     // 用户对象
	Status   LoginStatus `json:"status"`             // 登录状态
	Adapter  string      `json:"adapter"`            // 适配器名称
	Features []string    `json:"features,omitempty"` // 平台特性 列表
}
