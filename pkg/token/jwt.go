package token

import (
	"crypto/rsa"
	"github.com/dgrijalva/jwt-go"
	"strings"
	"time"
)

// CustomClaims 自定义的 metadata在加密后作为 JWT 的第二部分返回给客户端
type CustomClaims struct {
	UserName string `json:"user_name"`
	UserID   int64  `json:"uid"`
	Role     string `json:"role"`
	jwt.StandardClaims
}

// Token jwt
type Token struct {
	publicKey  *rsa.PublicKey
	privateKey *rsa.PrivateKey
}

func New(privateKeyByte, publicKeyByte []byte) (*Token, error) {
	publicKey, err := jwt.ParseRSAPublicKeyFromPEM(publicKeyByte)
	if err != nil {
		return nil, err
	}
	privateKey, err := jwt.ParseRSAPrivateKeyFromPEM(privateKeyByte)
	if err != nil {
		return nil, err
	}
	return &Token{
		privateKey: privateKey,
		publicKey:  publicKey,
	}, nil
}

//Decode 解码
func (srv *Token) Decode(tokenStr string) (*CustomClaims, error) {
	tokenStr = strings.ReplaceAll(tokenStr, "Bearer ", "")
	t, err := jwt.ParseWithClaims(tokenStr, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return srv.publicKey, nil
	})
	if err != nil {
		return nil, err
	}
	// 解密转换类型并返回
	if claims, ok := t.Claims.(*CustomClaims); ok && t.Valid {
		return claims, nil
	}
	return nil, err
}

// Encode 将 User 用户信息加密为 JWT 字符串
// expireTime := time.Now().Add(time.Hour * 24 * 3).Unix() 三天后过期
func (srv *Token) Encode(userName string, userId int64, role string, expireTime int64) (string, error) {
	claims := CustomClaims{
		userName,
		userId,
		role,
		jwt.StandardClaims{
			Issuer:    "whole.dev",
			NotBefore: time.Now().Unix(),
			IssuedAt:  time.Now().Unix(),
			ExpiresAt: expireTime,
			Subject:   "AuthToken",
		},
	}
	jwtToken := jwt.NewWithClaims(jwt.SigningMethodRS256, claims)
	tk, err := jwtToken.SignedString(srv.privateKey)
	if err != nil {
		return "", err
	}
	return tk, nil
}
