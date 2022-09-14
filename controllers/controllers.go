package controllers

import (
	"fmt"
	"strings"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

type baseController struct {
	beego.Controller
	o              orm.Ormer
	controllerName string
	actionName     string
}

func (c *baseController) Prepare() {
	c.Ctx.ResponseWriter.Header().Set("Access-Control-Allow-Origin", c.Ctx.Request.Header.Get("Origin"))
	controllerName, actionName := c.GetControllerAndAction()
	c.controllerName = strings.ToLower(controllerName[0 : len(controllerName)-10])
	c.actionName = strings.ToLower(actionName)
	c.o = orm.NewOrm()
	fmt.Println(c.Ctx.GetCookie("userid"))
	if strings.ToLower(c.controllerName) == "admin" && strings.ToLower(c.actionName) != "login" {
		if c.Ctx.GetCookie("userid") == "" {
			c.Gowithmessage("请登录", "/admin/login")
		} else if c.GetSession(c.Ctx.GetCookie("userid")) == nil {
			c.Gowithmessage("请登录", "/admin/login")
		} else {
		}
	}

}

func (c *baseController) Gowithmessage(msg string, url string) {
	if url == "" {
		c.Ctx.WriteString("<script>alert('" + msg + "');window.history.go(-1);</script>")
		c.StopRun()
	} else {
		c.Redirect(url, 302)
	}
}

//获取IP地址
func (c *baseController) getIp() string {
	s := strings.Split(c.Ctx.Request.RemoteAddr, ":")
	return s[0]
}
