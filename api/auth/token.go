package auth

import (
	"github.com/dgrijalva/jwt-go"
	"go-micro-example/service/constant/micro_c"
	"go-micro-example/service/user/proto"
	"time"
)

var privateKey = []byte("`xs#a_1-!")

// 自定义的 metadata
// 在加密后作为 JWT 的第二部分返回给客户端
type CustomClaims struct {
	User *user.UserInfo
	jwt.StandardClaims
}

func Decode(tokenStr string) (*user.UserInfo, error) {
	t, err := jwt.ParseWithClaims(tokenStr, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return privateKey, nil
	})

	if err != nil {
		return nil, err
	}

	// 解密转换类型并返回
	if claims, ok := t.Claims.(*CustomClaims); ok && t.Valid {
		return claims.User, nil
	} else {
		return nil, err
	}
}

func Encode(user *user.UserInfo) (string, error) {
	// 三天后过期
	expireTime := time.Now().Add(time.Hour * 24 * 3).Unix()
	claims := CustomClaims{
		user,
		jwt.StandardClaims{
			Issuer:    micro_c.MicroNameUser, // 签发者
			ExpiresAt: expireTime,
		},
	}
	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return jwtToken.SignedString(privateKey)
}
