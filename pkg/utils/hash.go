package utils

import (
	"crypto/hmac"
	"crypto/rand"
	"crypto/sha512"
	"encoding/base64"
)

const SALTSIZE = 16

func generateRandomSalt(saltSize int) (string, error) {
	bytes := make([]byte, saltSize)
	_, err := rand.Read(bytes)
	if err != nil {
		return "", err
	}
	return base64.StdEncoding.EncodeToString(bytes)[:saltSize], nil
}

func hashPassword(pwd, salt string) string {
	h := hmac.New(sha512.New, []byte(salt))
	h.Write([]byte(pwd))
	return base64.StdEncoding.EncodeToString(h.Sum(nil))
}

func Match(currPwd, hashedPwd, salt string) bool {
	currPwdHash := hashPassword(currPwd, salt)
	return currPwdHash == hashedPwd
}

func HashPassword(password string) (string, string, error) {
	salt, err := generateRandomSalt(SALTSIZE)
	if err != nil {
		return "", "", err
	}
	pwd := hashPassword(password, salt)
	return pwd, salt, nil
}
