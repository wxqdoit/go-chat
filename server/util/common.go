package util

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"net/http"
	"strings"
	"time"
)

//字段 首字母一定要大写
type BaseResponse struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data,omitempty"`
}

//Claim是一些实体（通常指的用户）的状态和额外的元数据
type Claims struct {
	Mobile   string `json:"mobile"`
	Password string `json:"password"`
	jwt.StandardClaims
}

// 指定加密密钥
var jwtSecret = []byte("go-chat")

func Resp(writer http.ResponseWriter, code int, data interface{}, msg string) {
	result := BaseResponse{
		Data: data,
		Code: code,
		Msg:  msg,
	}
	resultStrJson, err := json.Marshal(result)
	if err != nil {
		fmt.Println(err)
	}
	_, _ = writer.Write(resultStrJson)
}

func Md5Encode(data string) string {
	h := md5.New()
	h.Write([]byte(data)) // 需要加密的字符串为 123456
	cipherStr := h.Sum(nil)
	return hex.EncodeToString(cipherStr)
}

func MD5Encode(data string) string {
	return strings.ToUpper(Md5Encode(data))
}

func ValidatePassword(plainPassword, salt, password string) bool {
	return Md5Encode(plainPassword+salt) == password
}
func MakePassword(plainPassword, salt string) string {
	return Md5Encode(plainPassword + salt)
}

// 根据用户的用户名和密码产生token
func GenerateToken(mobile string, password string) (string, error) {
	//设置token有效时间
	nowTime := time.Now()
	//expireTime := nowTime.Add(24 * 30 * time.Hour)
	expireTime := nowTime.Add(10)
	claims := Claims{
		Mobile:   mobile,
		Password: password,
		StandardClaims: jwt.StandardClaims{
			// 过期时间
			ExpiresAt: expireTime.Unix(),
			// 指定token发行人
			Issuer: "wxqdoit",
		},
	}
	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	//该方法内部生成签名字符串，再用于获取完整、已签名的token
	token, err := tokenClaims.SignedString(jwtSecret)
	return token, err
}
func ParseToken(tokenString string) (c jwt.Claims, err error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})
	c, ok := token.Claims.(jwt.Claims)
	fmt.Println("----:", token.Valid)
	if ok && token.Valid {
		// 获取分区内容
		return c, err
	} else {
		return c, err
	}
}
