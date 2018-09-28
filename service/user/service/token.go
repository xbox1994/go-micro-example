package service

import (
	"GoMicroExample/service/user/proto"
	"github.com/dgrijalva/jwt-go"
	"time"
)

type Authable interface {
	Decode(tokenStr string) (*CustomClaims, error)
	Encode(user *user.User) (string, error)
}

var privateKey = []byte("`xs#a_1-!")

// 自定义的 metadata
// 在加密后作为 JWT 的第二部分返回给客户端
type CustomClaims struct {
	User *user.User
	jwt.StandardClaims
}

func (srv *TokenService) Decode(tokenStr string) (*CustomClaims, error) {
	t, err := jwt.ParseWithClaims(tokenStr, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return privateKey, nil
	})

	if err != nil {
		return nil, err
	}

	// 解密转换类型并返回
	if claims, ok := t.Claims.(*CustomClaims); ok && t.Valid {
		return claims, nil
	} else {
		return nil, err
	}
}

type TokenService struct {
}

//
// 将 User 用户信息加密为 JWT 字符串
//
func (srv *TokenService) Encode(user *user.User) (string, error) {
	// 三天后过期
	expireTime := time.Now().Add(time.Hour * 24 * 3).Unix()
	claims := CustomClaims{
		user,
		jwt.StandardClaims{
			Issuer:    "go.micro.srv.user", // 签发者
			ExpiresAt: expireTime,
		},
	}
	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return jwtToken.SignedString(privateKey)
}
