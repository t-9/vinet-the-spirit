package auth

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"strconv"
	"time"
)

var accessKey string
var accessSecret string

func SetAccessKey(k string) {
	accessKey = k
}

func SetAccessSecret(s string) {
	accessSecret = s
}

func GetAccessKey() string {
	return accessKey
}

func GenerateTimestamp() string {
	return strconv.FormatInt(time.Now().Unix(), 10)
}

func GenerateSign(method string, path string, body string) string {
	plain := GenerateTimestamp() + method + path + body
	mac := hmac.New(sha256.New, []byte(accessSecret))
	mac.Write([]byte(plain))
	return hex.EncodeToString(mac.Sum(nil))
}
