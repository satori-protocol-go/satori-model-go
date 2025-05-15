package meta

import "github.com/satori-protocol-go/satori-model-go/pkg/login"

// 元信息
type Meta struct {
	Logins    []*login.Login `json:"logins"`     // 登录信息
	ProxyUrls []string       `json:"proxy_urls"` // 代理路由 列表
}
