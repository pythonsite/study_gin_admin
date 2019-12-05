package app

import (
	"github.com/LyricTian/captcha"
	"github.com/LyricTian/captcha/store"
	"study_gin_admin/internal/app/config"
	"study_gin_admin/pkg/logger"
)

// InitCaptcha 初始化图形验证码
func InitCaptcha() {
	cfg := config.Global().Captcha
	if cfg.Store == "redis" {
		rc := config.Global().Redis
		captcha.SetCustomStore(store.NewRedisStore(&store.RedisOptions{
			Addr:     rc.Addr,
			Password: rc.Password,
			DB:       cfg.RedisDB,
		}, captcha.Expiration, logger.StandardLogger(), cfg.RedisPrefix))
	}
}