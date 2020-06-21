package models

import (
	"MoeBlog/conn"
	"MoeBlog/utils"
	"crypto/sha1"
	"encoding/hex"
	"log"
)

type User struct {
	ID          int    `gorm:"column:id;primary_key" json:"id"`
	NickName    string `gorm:"column:nickname" json:"nickname"`
	Telephone   string `gorm:"column:telephone" json:"telephone"`
	Password    string `gorm:"column:password" json:"password"`
	Sex         bool   `gorm:"column:sex" json:"sex"`
	Avatar      string `gorm:"column:avatar" json:"avatar"`
	Status      int    `gorm:"column:status" json:"status"`
	AccessToken string `gorm:"column:access_token" json:"access_token"`
	LastLogin   string `gorm:"column:last_login" json:"last_login"`
	CreateTime  string `gorm:"column:create_time" json:"create_time"`
	UpdateTime  string `gorm:"column:update_time" json:"update_time"`
	DeleteTime  string `gorm:"column:delete_time" json:"delete_time"`
}

//返回一个新的User
func NewUser() User {
	return User{}
}

//通过手机号获取用户
func GetUserByTelephone(telephone string) (User, error) {
	db := conn.GetInstance()
	var user User
	result := db.Where("telephone = ?", telephone).First(&user)
	return user, result.Error
}

//密码加密
func (u *User) SetPassword(password string) {
	salt := utils.GenRandomString(16)
	hash := sha1.New()
	_, err := hash.Write([]byte(password + salt))
	if err != nil {
		log.Println("[utils] Generate encrypted password error...")
	}
	bytes := hash.Sum(nil)
	u.Password = salt + ":" + hex.EncodeToString(bytes)
}

func (u *User) SetAccessToken() {
	u.AccessToken = utils.GenToken(u.Telephone)
}
