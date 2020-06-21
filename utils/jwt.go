package utils

import (
	"crypto"
	"encoding/hex"
	"fmt"
	"log"
	"time"
)

func GenToken(telephone string) string {
	ts := fmt.Sprintf("%x", time.Now().Unix())
	hash := crypto.MD5.New()
	_, err := hash.Write([]byte(telephone + ts))
	if err != nil {
		log.Println("[utils] generate token error,", err)
		return ""
	}
	return hex.EncodeToString(hash.Sum(nil)) + ":" + ts
}
