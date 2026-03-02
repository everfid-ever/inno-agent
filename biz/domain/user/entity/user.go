package entity

import (
	"github.com/xh-polaris/inno_agent/biz/domain/user/dal/model"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// User 领域实体，组合数据实体并提供领域方法
type User struct {
	*model.User
}

func NewUser(m *model.User) *User {
	return &User{User: m}
}

// IdHex 返回用户 ObjectID 的 hex 字符串
func (u *User) IdHex() string {
	return u.ID.Hex()
}

// ObjectID 返回用户 ObjectID
func (u *User) ObjectID() primitive.ObjectID {
	return u.ID
}