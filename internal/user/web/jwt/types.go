package jwt

import "github.com/NotFound1911/mserver"

type Handler interface {
	ClearToken(ctx *mserver.Context) error
	ExtractToken(ctx *mserver.Context) string
	SetLoginToken(ctx *mserver.Context, uid int64) error
	SetJWTToken(ctx *mserver.Context, uid int64, ssid string) error
	CheckSession(ctx *mserver.Context, ssid string) error
}
