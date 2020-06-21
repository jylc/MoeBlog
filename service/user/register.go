package user

import (
	"MoeBlog/conn"
	"MoeBlog/models"
	"MoeBlog/service/resp"
	"MoeBlog/service/verify"
	"MoeBlog/utils"
	"errors"
	"github.com/gin-gonic/gin"
	"time"
)

type RegisterService struct {
	Telephone string `form:"telephone" json:"telephone"`
	Password  string `form:"password" json:"password"`
}

func (rs *RegisterService) Register(ctx *gin.Context) resp.Response {
	if verify.IsValidTelephone(rs.Telephone) == false {
		return resp.ParamErr("手机号格式错误", errors.New("telephone format error"))
	}

	if verify.IsValidPassword(rs.Password) == false {
		return resp.ParamErr("密码格式错误", errors.New("password format error"))
	}

	db := conn.GetInstance()
	count := 0
	db.Model(&models.User{}).Where("telephone = ?", rs.Telephone).Count(&count)
	if count != 0 {
		return resp.Response{
			Code: resp.CodeFailed,
			Msg:  "register filed",
			Err:  errors.New("this telephone has been registered").Error(),
		}
	}

	user := models.NewUser()
	user.NickName = GenDefaultNickName()
	user.Telephone = rs.Telephone
	user.SetPassword(rs.Password)
	user.SetAccessToken()
	user.Sex = true
	user.CreateTime = time.Now().Format("2006-01-02 3:04:05.000 PM Mon Jan")
	err := db.Create(&user).Error
	if err != nil {
		return resp.Response{
			Code: resp.CodeFailed,
			Msg:  "register filed",
			Err:  err.Error(),
		}
	}
	return resp.Response{
		Code: resp.CodeSuccess,
		Data: user,
		Msg:  "register success",
	}
}

func GenDefaultNickName() string {
	return utils.GenRandomString(6)
}
