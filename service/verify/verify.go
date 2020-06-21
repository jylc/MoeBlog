package verify

import (
	"crypto/sha1"
	"encoding/hex"
	"log"
	"regexp"
	"strings"
)

func PasswordVerify(pwd, realPwd string) bool {
	temp := strings.Split(realPwd, ":")
	if len(temp) != 2 {
		return false
	}
	salt := temp[0]
	hash := sha1.New()
	_, err := hash.Write([]byte(pwd+salt))
	if err != nil {
		log.Println("[verify] Generate pwd error, ", err)
		return false
	}
	s := hex.EncodeToString(hash.Sum(nil))
	if strings.Compare(s, temp[1]) != 0 {
		log.Println("[verify] Password error")
		return false
	}
	return true
}

func IsValidTelephone(telephone string) bool {
	regular := "^((13[0-9])|(14[5,7])|(15[0-3,5-9])|(17[0,3,5-8])|(18[0-9])|166|198|199|(147))\\d{8}$"
	compile, err := regexp.Compile(regular)
	if err != nil {
		return false
	}
	if match := compile.MatchString(telephone); !match {
		return false
	}
	return true
}

func IsValidPassword(password string) bool {
	if len(password) < 6 {
		log.Println("[verify] Password length must large than 6")
		return false
	}
	return true
}
