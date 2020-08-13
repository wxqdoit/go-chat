package test

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"time"
)

//Claim是一些实体（通常指的用户）的状态和额外的元数据
type Claims struct {
	Username string `json:"username"`
	Password string `json:"password"`
	jwt.StandardClaims
}

// 指定加密密钥
var jwtSecret = []byte("go-chat")

// 根据用户的用户名和密码产生token
func GenerateToken(username, password string) (string, error) {
	fmt.Println("-------------" + username)
	//设置token有效时间
	nowTime := time.Now()
	expireTime := nowTime.Add(3 * time.Hour)

	claims := Claims{
		Username: username,
		Password: password,
		StandardClaims: jwt.StandardClaims{
			// 过期时间
			ExpiresAt: expireTime.Unix(),
			// 指定token发行人
			Issuer: "gin-blog",
		},
	}

	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodRS256, claims)
	//该方法内部生成签名字符串，再用于获取完整、已签名的token
	token, err := tokenClaims.SignedString(jwtSecret)

	return token, err
}
func ParseToken(tokenString string) (c jwt.Claims, err error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
		return jwtSecret, nil
	})
	c, ok := token.Claims.(jwt.Claims)
	fmt.Println("----", c, ok)

	if ok && token.Valid {
		// 获取分区内容
		return c, err
	} else {
		return c, err

	}
}
