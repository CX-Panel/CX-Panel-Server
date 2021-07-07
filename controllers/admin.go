package controllers

import (
	"CX-Panel-Server/models"
	//"encoding/json"

	"github.com/astaxie/beego"
)

// 关于管理员的操作
type AdminController struct {
	beego.Controller
}

// @Title 获得信息
// @Description 获得管理员uid
// @Param	uid		path 	string	true		"uid"
// @Success 200 {object} models.User
// @Failure 403 :uid 不存在
// @router /:uid [get]
func (this *AdminController) Get() {
	uid := this.GetString(":uid")
	if uid != "" {
		user, err := models.GetUser(uid)
		if err != nil {
			this.Data["json"] = err.Error()
		} else {
			this.Data["json"] = user
		}
	}
	this.ServeJSON()
}

// @Title 登录
// @Description 管理员登录系统
// @Param	username		query 	string	true		"用户名"
// @Param	password		query 	string	true		"密码"
// @Success 200 {string} 登录成功
// @Failure 403 登录失败
// @router /login [get]
func (this *AdminController) Login() {
	username := this.GetString("username")
	password := this.GetString("password")
	if models.Login(username, password) {
		this.Data["json"] = "登录成功"
	} else {
		this.Data["json"] = "登录失败"
	}
	this.ServeJSON()
}

// @Title 登出
// @Description 注销当前管理员的会话
// @Success 200 {string} 注销成功
// @router /logout [get]
func (this *AdminController) Logout() {
	this.Data["json"] = "注销成功"
	this.ServeJSON()
}
