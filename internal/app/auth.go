package app

import (
	"study_gin_admin/internal/app/config"
	"study_gin_admin/pkg/auth"
)

// InitAuth 初始化用户认证
func InitAuth()(auth.Auther, error) {
	cfg := config.Global().JWTAuth


}