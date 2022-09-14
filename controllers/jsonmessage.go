package controllers

import (
	"FEINESV/spymongo"
	"fmt"
	"reflect"

	"gopkg.in/mgo.v2/bson"
)

func (c *AdminController) Getmessage() {
	if c.Ctx.Request.Method == "GET" {

		name := c.GetString("name")
		searchid := c.GetString("search")
		result := map[string]interface{}{} //标准的储存格式

		pacong := "spymongo/py/爬虫" + searchid + ".py"
		resulttemp := spymongo.GAIJosnDatapy(name, pacong)
		result, ok := resulttemp.(map[string]interface{})
		fmt.Println("type:", reflect.TypeOf(resulttemp))
		if ok {

		} else {
			fmt.Println("断言错误")
			c.Gowithmessage("断言错误这种万恶的情况出现了，快联系射老师修理！", "")
			return
		}
		c.Data["json"] = result
		ds := spymongo.NewSessionStore()
		defer ds.Close()
		col := ds.GetBufferCollection()
		for v, k := range result {
			one := map[string]interface{}{}
			one[v] = k
			selector := bson.M{v: bson.M{"$exists": true}}
			_, err := col.Upsert(selector, one)
			if err != nil {
				fmt.Println("upsert错误")
				c.Gowithmessage("更新错误", "")
			}
		}
		c.ServeJSON()
	}
}

func (c *AdminController) GetDetail() {
	if c.Ctx.Request.Method == "GET" {
		name := c.GetString("name")
		ds := spymongo.NewSessionStore()
		defer ds.Close()
		col := ds.GetStoreCollection()
		result := map[string]map[string]map[string]string{}
		err := col.Find(bson.M{name: bson.M{"$exists": true}}).One(&result)
		if err != nil {
			fmt.Println("读取失败")
		}
		newresult := map[string]map[string]string{}
		for v, k := range result {
			if v == name {
				newresult = k
				break
			}
		}
		c.Data["json"] = newresult
		c.ServeJSON()
	}
}

func (c *AdminController) GetBufferDetail() {
	if c.Ctx.Request.Method == "GET" {
		name := c.GetString("name")
		ds := spymongo.NewSessionStore()
		defer ds.Close()
		col := ds.GetBufferCollection()
		result := map[string]map[string]map[string]string{}
		err := col.Find(bson.M{name: bson.M{"$exists": true}}).One(&result)
		if err != nil {
			fmt.Println("读取失败")
		}
		newresult := map[string]map[string]string{}
		for v, k := range result {
			if v == name {
				newresult = k
				break
			}
		}
		c.Data["json"] = newresult
		c.ServeJSON()
	}
}

func (c *AdminController) GetStoreList() {
	if c.Ctx.Request.Method == "GET" {

		ds := spymongo.NewSessionStore()
		defer ds.Close()
		col := ds.GetStoreCollection()
		result := []map[string]map[string]map[string]string{}
		err := col.Find(nil).All(&result)
		if err != nil {
			fmt.Println("读取失败")
			c.Gowithmessage("加载列表失败", "")
			return
		}
		var srclist []string
		for _, k := range result {
			for name := range k {
				srclist = append(srclist, name)
			}
		}
		fmt.Println("完成------------------------>")
		c.Data["json"] = srclist
		c.ServeJSON()
	}
}

func (c *AdminController) GetBufferList() {
	if c.Ctx.Request.Method == "GET" {
		ds := spymongo.NewSessionStore()
		defer ds.Close()
		col := ds.GetBufferCollection()
		result := []map[string]map[string]map[string]string{}
		err := col.Find(nil).All(&result)
		if err != nil {
			fmt.Println("读取失败")
			c.Gowithmessage("加载列表失败", "")
			return
		}
		var srclist []string
		for _, k := range result {
			for name := range k {
				srclist = append(srclist, name)
			}
		}
		c.Data["json"] = srclist
		c.ServeJSON()
	}
}

func (c *AdminController) PushChange() {
	if c.Ctx.Request.Method == "POST" {
		name := c.GetString("name")
		pwd := c.GetString("password")
		if pwd != "72859818" {
			c.Gowithmessage("失败", "")
			return
		}
		ds := spymongo.NewSessionStore()
		defer ds.Close()
		col := ds.GetBufferCollection()
		result := map[string]map[string]map[string]string{}
		err := col.Find(bson.M{name: bson.M{"$exists": true}}).One(&result)
		if err != nil {
			fmt.Println("读取buffer失败")
			c.Gowithmessage("加载buffer失败", "")
			return
		}
		selector := bson.M{name: bson.M{"$exists": true}}
		col = ds.GetStoreCollection()
		_, err = col.Upsert(selector, result)
		if err != nil {
			fmt.Println("upsert错误")
			c.Gowithmessage("更新错误", "")
			return
		}
		c.Data["json"] = map[string]interface{}{"state": "su"}
		c.ServeJSON()
	}
}
