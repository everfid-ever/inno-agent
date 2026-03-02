package token

import (
	"errors"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/xh-polaris/inno_agent/biz/conf"
	"github.com/xh-polaris/inno_agent/biz/types/cst"
)

// Info JWT 载荷信息
type Info struct {
	RawToken    string
	BasicUserId string         `json:"basic_user_id"` // 用户ID
	Code        string         `json:"code"`          // 学号
	Phone       string         `json:"phone"`         // 手机号
	Email       string         `json:"email"`         // 邮箱
	LoginTime   int64          `json:"login_time"`    // 登录时间(秒时间戳)
	AuthType    string         `json:"auth_type"`     // 登录类型
	Extra       map[string]any `json:"extra"`         // 额外信息
}

// SignJWT 使用 RSA 私钥签发 JWT
func SignJWT(tokenConf *conf.Token, info *Info) (string, error) {
	now := time.Now().UTC()

	claims := jwt.MapClaims{
		"iat":                now.Unix(),
		"nbf":                now.Unix(),
		"exp":                now.Add(time.Duration(tokenConf.Expire) * time.Second).Unix(),
		cst.TokenBasicUserID: info.BasicUserId,
		cst.TokenCode:        info.Code,
		cst.TokenPhone:       info.Phone,
		cst.TokenEmail:       info.Email,
		cst.TokenLoginTime:   info.LoginTime,
		cst.TokenAuthType:    info.AuthType,
		cst.TokenExtra:       info.Extra,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodRS256, claims)
	sk, err := conf.GetSecretKey(tokenConf.SecretKey)
	if err != nil {
		return "", err
	}
	return token.SignedString(sk)
}

// ParseJWT 使用 RSA 公钥验证并解析 JWT
func ParseJWT(tokenConf *conf.Token, str string) (*Info, error) {
	token, err := jwt.Parse(str, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodRSA); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return conf.GetPublicKey(tokenConf.PublicKey)
	})
	if err != nil {
		return nil, err
	}
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		extra, _ := claims[cst.TokenExtra].(map[string]any)
		info := &Info{
			BasicUserId: claims[cst.TokenBasicUserID].(string),
			Code:        claims[cst.TokenCode].(string),
			Phone:       claims[cst.TokenPhone].(string),
			Email:       claims[cst.TokenEmail].(string),
			LoginTime:   int64(claims[cst.TokenLoginTime].(float64)),
			AuthType:    claims[cst.TokenAuthType].(string),
			Extra:       extra,
			RawToken:    str,
		}
		return info, nil
	}
	return nil, errors.New("invalid token")
}
