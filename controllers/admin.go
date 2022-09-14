package controllers

import (
	"FEINESV/models"
	"FEINESV/util"
	"crypto/md5"
	"encoding/hex"
	"strings"
)

//AdminController继承基础控制器
type AdminController struct {
	baseController
}

//Login in
func (c *AdminController) Login() {
	if c.Ctx.Request.Method == "POST" {
		username := c.GetString("username")
		password := c.GetString("password")
		user := models.User{Username: username}
		c.o.Read(&user, "username")

		if user.Password == "" {
			c.Gowithmessage("账号不存在", "")
		}
		if util.Md5(password) != strings.Trim(user.Password, "") {
			c.Gowithmessage("密码错误", "")
		}
		user.LastIp = c.getIp()
		user.LoginCount = user.LoginCount + 1
		if _, err := c.o.Update(&user); err != nil {
			c.Gowithmessage("登陆失败", "")
		} else {
			encodeId := md5.New()
			encodeId.Write([]byte(username))
			md5encodeId := encodeId.Sum(nil)
			md5StringID := hex.EncodeToString(md5encodeId)
			c.Ctx.SetCookie("userid", md5StringID)
			c.Gowithmessage("登陆成功", "/admin/main")
			c.SetSession(md5StringID, int(user.Id))
		}
	}
	c.TplName = c.controllerName + "/login.html"
}

//Main context
func (c *AdminController) Main() {
	c.TplName = c.controllerName + "/index.html"
}

//Logout and destroy the session
func (c *AdminController) Logout() {
	sessionName := c.Ctx.GetCookie("userid")
	c.DelSession(sessionName)
	c.Gowithmessage("退出登陆", "/admin/login.html")
}
