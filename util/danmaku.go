package util

import (
	"encoding/json"
	// "fmt"
	"fmt"
	"github.com/dgrijalva/jwt-go"
)

type Key struct {
	Key string
}
type DanmakuTokenToken struct {
	DanmakuToken string
}
type Token struct {
	Token string `json:"token"`
}

func GetDanmakuToken(userId int) string {
	baseURL := "http://192.168.10.111:8010"
	response := HTTPGet(baseURL+"/auth/key", map[string]string{})
	var key Key
	json.Unmarshal(response, &key)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userId": userId,
	})
	fmt.Println("key:", key)
	// Sign and get the complete encoded token as a string using the secret
	secret := "Cp@2112%" + key.Key
	tokenString, error := token.SignedString([]byte(secret))
	fmt.Println("token string:", tokenString)
	fmt.Println("error:", error)
	t := Token{
		Token: tokenString,
	}
	res := HTTPPost(baseURL+"/auth/autoLogin?key="+key.Key, t)
	var danmakuToken DanmakuTokenToken
	json.Unmarshal(res, &danmakuToken)
	return danmakuToken.DanmakuToken
}
