package service

import (
	"context"

	"github.com/xh-polaris/inno_agent/biz/domain/user/entity"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// UserDomainSVC 领域服务接口，向 Application 层提供 User 领域能力
type UserDomainSVC interface {
	// Login 通过 Synapse 完成登录，返回用户实体与是否新用户
	// authType: phone-password / phone-verify / email-password / email-verify / studentId-password / studentId-verify
	Login(ctx context.Context, authType, authId, verify string) (*entity.User, bool, error)

	// Register 通过 Synapse 完成注册
	Register(ctx context.Context, authType, authId, verify, password string) error

	// ResetPassword 重置密码（需透传 Authorization header）
	ResetPassword(ctx context.Context, authHeader, newPassword string) error

	// GetUser 根据本地 ID 获取用户实体
	GetUser(ctx context.Context, userId string) (*entity.User, error)

	// UpdateField 更新用户字段
	UpdateField(ctx context.Context, id primitive.ObjectID, fields bson.M) error
}