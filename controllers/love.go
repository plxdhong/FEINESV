package controllers

import (
	"FEINESV/spymongo"
	"fmt"

	"gopkg.in/mgo.v2/bson"
)

type LovePerson struct {
	Userid   string
	LoveName []string
}

func (c *AdminController) GetLoveList() {
	if c.Ctx.Request.Method == "GET" {

		if c.GetSession(c.Ctx.GetCookie("userid")) == nil {
			c.Gowithmessage("获取用户失败请重新登陆", "/admin/login")
			return
		}
		uid := c.Ctx.GetCookie("userid")
		ds := spymongo.NewSessionStore()
		defer ds.Close()
		col := ds.GetLoveCollection()
		result := LovePerson{uid, []string{}}
		err := col.Find(bson.M{"userid": uid}).One(&result)
		if err != nil {
			c.Gowithmessage("获取失败", "")
			return
		}
		c.Data["json"] = result
		c.ServeJSON()
	}
}

func (c *AdminController) AddLove() {
	if c.Ctx.Request.Method == "GET" {
		name := c.GetString("name")
		if c.GetSession(c.Ctx.GetCookie("userid")) == nil {
			c.Gowithmessage("获取用户失败请重新登陆", "/admin/login")
			return
		}
		uid := c.Ctx.GetCookie("userid")
		ds := spymongo.NewSessionStore()
		defer ds.Close()
		col := ds.GetLoveCollection()
		result := LovePerson{uid, []string{}}
		err := col.Find(bson.M{"userid": uid}).One(&result)
		if err != nil {
			c.Gowithmessage("获取用户喜爱失败", "")
		}
		result.LoveName = append(result.LoveName, name)
		_, err = col.Upsert(bson.M{"userid": uid}, result)
		if err != nil {
			fmt.Println("喜欢失败")
			c.Gowithmessage("喜欢失败", "")
			return
		}
		c.Data["json"] = result
		c.ServeJSON()
	}
}
