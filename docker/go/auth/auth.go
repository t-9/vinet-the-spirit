package auth

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/hmac"
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
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

// GetAccessKey gets an AccessKey.
func GetAccessKey() string {
	return accessKey
}

// GenerateTimestamp generates a timestamp.
func GenerateTimestamp() string {
	return strconv.FormatInt(time.Now().Unix(), 10)
}

// GenerateSign generates a sign.
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

// GetRequest requests a get repuest.
func GetRequest(path string, params map[string]string) ([]byte, error) {
	return sendRequest(http.MethodGet, path, params, nil)
}

// PostRequest requests a post repuest.
func PostRequest(path string, b interface{}) ([]byte, error) {
	return sendRequest(http.MethodPost, path, map[string]string{}, b)
}

func sendRequest(method, path string, params map[string]string, b interface{}) ([]byte, error) {
	url := "https://api.bitflyer.com" + path

	reqBody := []byte{}
	var reqBodyReader io.Reader
	if b != nil {
		var err error
		reqBody, err = json.Marshal(b)
		if err != nil {
			return []byte{}, err
		}
		reqBodyReader = bytes.NewReader(reqBody)
	}

	req, err := http.NewRequest(method, url, reqBodyReader)
	if err != nil {
		return []byte{}, err
	}

	timestamp := GenerateTimestamp()
	sign, err := GenerateSign(timestamp, method, path, string(reqBody))
	if err != nil {
		return []byte{}, err
	}

	req.Header.Set("ACCESS-KEY", GetAccessKey())
	req.Header.Set("ACCESS-TIMESTAMP", timestamp)
	req.Header.Set("ACCESS-SIGN", sign)
	req.Header.Set("Content-Type", "application/json")

	values := req.URL.Query()
	for k, v := range params {
		values.Add(k, v)
	}
	req.URL.RawQuery = values.Encode()

	timeout := time.Duration(5 * time.Second)
	client := &http.Client{
		Timeout: timeout,
	}

	resp, err := client.Do(req)
	if err != nil {
		return []byte{}, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return []byte{}, err
	}

	if resp.StatusCode != 200 {
		return []byte{}, fmt.Errorf(string(body))
	}

	return body, nil
}
