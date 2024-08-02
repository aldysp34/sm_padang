package util

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"math/rand"
	"reflect"
	"time"

	"golang.org/x/crypto/bcrypt"
)

func ComparePassword(hashedPwd, plainPwd string) bool {
	byteHash := []byte(hashedPwd)
	password := []byte(plainPwd)
	err := bcrypt.CompareHashAndPassword(byteHash, password)
	return err == nil
}

func Contains(slice []string, item string) bool {
	for _, s := range slice {
		if s == item {
			return true
		}
	}
	return false
}

func InterfaceToString(i any) string {
	if i == nil {
		return ""
	}
	// token, err := base64.StdEncoding.DecodeString(i.(string))
	// if err != nil {
	// 	return ""
	// }
	// return string(token)
	v := reflect.ValueOf(i)
	switch v.Kind() {
	case reflect.String:
		return v.String()
	default:
		return fmt.Sprintf("%v", i)
	}
}

func GenerateCode() string {
	rand.Seed(time.Now().UnixNano())

	randomString := func(length int) string {
		charset := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890"
		b := make([]byte, length)
		for i := range b {
			b[i] = charset[rand.Intn(len(charset))]
		}
		return string(b)
	}

	hasher := sha256.New()
	hasher.Write([]byte(randomString(10)))
	hash := hex.EncodeToString(hasher.Sum(nil))

	return hash
}
