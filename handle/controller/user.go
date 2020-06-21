package controller

import (
	"MoeBlog/service/resp"
	"MoeBlog/service/user"
	"github.com/gin-gonic/gin"
)

// @Summary  注册
// @Description 用户注册
// @Tags 测试
// @Accept mpfd
// @Produce json
// @Param telephone formData string true "手机号"
// @Param password formData string true "密码"
// @Success 200 {string} string "{"msg": "注册成功"}"
// @Failure 400 {string} string "{"msg": "注册失败"}"
// @Router /register [post]
func UserRegister(c *gin.Context) {
	var service user.RegisterService
	err := c.ShouldBind(&service)
	if err == nil {
		res := service.Register(c)
		c.JSON(200, res)
	} else {
		c.JSON(200, resp.Response{
			Code: resp.CodeFailed,
			Data: nil,
			Msg:  "",
			Err:  "error",
		})
	}
}

// @Summary  登录
// @Description 用户登录
// @Tags 测试
// @Accept mpfd
// @Produce json
// @Param telephone formData string true "手机号"
// @Param password formData string true "密码"
// @Success 200 {string} string "{"msg": "登录成功"}"
// @Failure 400 {string} string "{"msg": "登录失败"}"
// @Router /login [post]
func UserLogin(c *gin.Context) {
	var service user.LoginService
	err := c.ShouldBind(&service)
	if err == nil {
		res := service.Login(c)
		c.JSON(200, res)
	} else {
		c.JSON(200, resp.Response{
			Code: resp.CodeFailed,
			Data: nil,
			Msg:  "",
			Err:  "error",
		})
	}
}
