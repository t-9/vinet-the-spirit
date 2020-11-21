package auth

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/hmac"
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"io"
	"strconv"
	"time"
)

var accessKey string
var cipheredAccessSecret []byte
var aesKey string
var iv []byte
var secretLen int

func setAccessKey(k string) {
	accessKey = k
}

func setAccessSecret(s string) error {
	b := make([]byte, 16)

	_, err := io.ReadFull(rand.Reader, b)
	if err != nil {
		return err
	}
	aesKey = base64.StdEncoding.EncodeToString(b)

	c, err := aes.NewCipher([]byte(aesKey))
	if err != nil {
		return err
	}

	iv = make([]byte, 16)
	_, err = io.ReadFull(rand.Reader, iv)
	if err != nil {
		return err
	}

	cfb := cipher.NewCFBEncrypter(c, iv)
	secretLen = len(s)
	ciphered := make([]byte, secretLen)
	cfb.XORKeyStream(ciphered, []byte(s))

	cipheredAccessSecret = ciphered
	return nil
}

func GetAccessKey() string {
	return accessKey
}

func GenerateTimestamp() string {
	return strconv.FormatInt(time.Now().Unix(), 10)
}

func GenerateSign(timestamp string, method string, path string, body string) (string, error) {
	c, err := aes.NewCipher([]byte(aesKey))
	if err != nil {
		return "", err
	}

	cfb := cipher.NewCFBDecrypter(c, iv)
	secret := make([]byte, secretLen)
	cfb.XORKeyStream(secret, cipheredAccessSecret)

	plain := timestamp + method + path + body

	mac := hmac.New(sha256.New, []byte(secret))

	mac.Write([]byte(plain))
	return hex.EncodeToString(mac.Sum(nil)), nil
}
