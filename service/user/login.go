package user

import (
	"MoeBlog/models"
	"MoeBlog/service/resp"
	"MoeBlog/service/verify"
	"errors"
	"github.com/gin-gonic/gin"
	"time"
)

type LoginService struct {
	Telephone string `form:"telephone" json:"telephone"`
	Password  string `form:"password" json:"password"`
}

func (ls *LoginService) Login(ctx *gin.Context) resp.Response {
	if verify.IsValidTelephone(ls.Telephone) == false {
		return resp.ParamErr("手机号格式错误", errors.New("telephone format error"))
	}

	if verify.IsValidPassword(ls.Password) == false {
		return resp.ParamErr("密码格式错误", errors.New("password format error"))
	}
	user, err := models.GetUserByTelephone(ls.Telephone)
	if err != nil {
		return resp.Response{
			Code: resp.CodeFailed,
			Msg:  "this telephone has not registered",
			Err:  err.Error(),
		}
	}
	if result := verify.PasswordVerify(ls.Password, user.Password); result == false {
		return resp.Response{
			Code: resp.CodeFailed,
			Msg:  "password is wrong try again",
		}
	}
	u := models.NewUser()
	u.Telephone = user.Telephone
	u.NickName = user.NickName
	u.AccessToken = user.AccessToken
	u.Sex = user.Sex
	u.Avatar = user.Avatar
	u.LastLogin = time.Now().Format("2006-01-02 3:04:05.000 PM Mon Jan")
	return resp.Response{
		Code: resp.CodeSuccess,
		Data: user,
		Msg:  "login successful",
	}
}
