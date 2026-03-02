package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

const CollectionName = "user"

// User MongoDB 数据实体
type User struct {
	ID          primitive.ObjectID `bson:"_id,omitempty"`
	BasicUserId string             `bson:"basicUserId"`         // Synapse 账号中台 ID
	Name        string             `bson:"name"`                // 昵称
	Avatar      string             `bson:"avatar"`              // 头像 URL
	Phone       string             `bson:"phone,omitempty"`     // 手机号（phone 登录时写入）
	Email       string             `bson:"email,omitempty"`     // 邮箱（email 登录时写入）
	StudentId   string             `bson:"studentId,omitempty"` // 学号（studentId 登录时写入）
	Profile     *UserProfile       `bson:"profile,omitempty"`   // 用户资料（暂时留空）
	CreatedAt   time.Time          `bson:"createdAt"`
	UpdatedAt   time.Time          `bson:"updatedAt"`
}

// UserProfile 用户资料（字段暂时留空）
type UserProfile struct {
}