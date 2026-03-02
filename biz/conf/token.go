package conf

import (
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"errors"
)

// Token JWT 签发配置
type Token struct {
	PublicKey string
	SecretKey string
	Expire    int64
}

// GetSecretKey 从 PEM 字符串解析 RSA 私钥
func GetSecretKey(sk string) (*rsa.PrivateKey, error) {
	b, _ := pem.Decode([]byte(sk))
	if b == nil {
		return nil, errors.New("failed to decode PEM block for secret key")
	}
	privateKey, err := x509.ParsePKCS8PrivateKey(b.Bytes)
	if err != nil {
		return nil, err
	}
	rsaKey, ok := privateKey.(*rsa.PrivateKey)
	if !ok {
		return nil, errors.New("not an RSA private key")
	}
	return rsaKey, nil
}

// GetPublicKey 从 PEM 字符串解析 RSA 公钥
func GetPublicKey(pk string) (*rsa.PublicKey, error) {
	b, _ := pem.Decode([]byte(pk))
	if b == nil {
		return nil, errors.New("failed to decode PEM block for public key")
	}
	pubKey, err := x509.ParsePKIXPublicKey(b.Bytes)
	if err != nil {
		return nil, err
	}
	rsaKey, ok := pubKey.(*rsa.PublicKey)
	if !ok {
		return nil, errors.New("not an RSA public key")
	}
	return rsaKey, nil
}
