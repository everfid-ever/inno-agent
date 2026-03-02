package synapse

import "context"

// LoginResult Synapse 登录返回结果
type LoginResult struct {
	BasicUserId string
	Token       string
	IsNew       bool
}

// RegisterResult Synapse 注册返回结果
type RegisterResult struct {
	Token string
}

// Client Synapse 账号中台 HTTP 客户端接口
type Client interface {
	// Login 调用 Synapse 登录接口
	Login(ctx context.Context, authType, authId, verify string) (*LoginResult, error)
	// Register 调用 Synapse 注册接口
	Register(ctx context.Context, authType, authId, verify, password string) (*RegisterResult, error)
	// ResetPassword 调用 Synapse 重置密码接口（需 Authorization header）
	ResetPassword(ctx context.Context, authHeader, newPassword string) error
}