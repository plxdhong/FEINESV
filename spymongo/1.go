package spymongo

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"os/exec"
	"time"

	"github.com/astaxie/beego"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

var session *mgo.Session

func Init() {
	session = ConnectToDB()
}

//储存主session的对象
type SessionStore struct {
	session *mgo.Session
}

type People struct {
	Name  string
	Phone string
}

func ConnectToDB() *mgo.Session {
	mongoDialInfo := &mgo.DialInfo{
		Addrs:    []string{"127.0.0.1:27017"},
		Timeout:  60 * time.Second,
		Database: "fv",
		Username: "fv",
		Password: "whyisyyw",
	}
	session, _ := mgo.DialWithInfo(mongoDialInfo)
	return session
}

//建立新session
func NewSessionStore() *SessionStore {
	ns := &SessionStore{
		session: session.Copy(),
	}
	return ns
}

//SessionStore的析函数
func (s *SessionStore) Close() {
	s.session.Close()
}

func (s *SessionStore) GetStoreCollection() *mgo.Collection {
	mdbname := beego.AppConfig.String("mdbname")
	mdbstore := beego.AppConfig.String("mdbstore")
	//mdbbuffer := beego.AppConfig.String("mdbbuffer")
	c := s.session.DB(mdbname).C(mdbstore)
	return c
}

func (s *SessionStore) GetBufferCollection() *mgo.Collection {
	mdbname := beego.AppConfig.String("mdbname")
	//mdbstore := beego.AppConfig.String("mdbstore")
	mdbbuffer := beego.AppConfig.String("mdbbuffer")
	c := s.session.DB(mdbname).C(mdbbuffer)
	return c
}

func (s *SessionStore) GetLoveCollection() *mgo.Collection {
	mdbname := beego.AppConfig.String("mdbname")
	mdbstore := beego.AppConfig.String("mdblove")
	//mdbbuffer := beego.AppConfig.String("mdbbuffer")
	c := s.session.DB(mdbname).C(mdbstore)
	return c
}

//func ConnectToDB() *mgo.Collection {
//	mdbhost := beego.AppConfig.String("mdbhost")
//	mdbname := beego.AppConfig.String("mdbname")
//	//mdbstore := beego.AppConfig.String("mdbstore")
//	mdbbuffer := beego.AppConfig.String("mdbbuffer")
//	session, err := mgo.Dial(mdbhost)
//	if err != nil {
//		panic(err)
//	}
//	//defer session.Close()
//	session.SetMode(mgo.Monotonic, true)
//	c := session.DB(mdbname).C(mdbbuffer)
//	return c
//}

//func InsertToDB(jsondata interface{}) {
//	c := ConnectToDB()
//	err := c.Insert(&jsondata)
//	if err != nil {
//		fmt.Println("插入数据库错误")
//		log.Fatal(err)
//	}
//}

func InsertToDB(jsondata interface{}) {
	ds := NewSessionStore()
	defer ds.Close()
	c := ds.GetBufferCollection()
	err := c.Insert(&jsondata)
	if err != nil {
		fmt.Println("插入数据库错误")
		log.Fatal(err)
	}
}

func StoreToDB(jsondata interface{}) {
	ds := NewSessionStore()
	defer ds.Close()
	c := ds.GetStoreCollection()
	err := c.Insert(&jsondata)
	if err != nil {
		fmt.Println("插入数据库错误")
		log.Fatal(err)
	}
}

/*
func GetPeopleInfo() {
	c := ConnectToDB()
	result := People{}
	err := c.Find(bson.M{"name": "kkk"}).One(&result)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("people", result)
	peoples := make([]People, 20)
	err = c.Find(nil).All(&peoples)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(peoples)
}
*/

//GAIJosnDatapy 得到参数一：剧名、参数二：python文件所在地址./py/....该操作将依次将结果插入DB
func GAIJosnDatapy(string_a string, pyfilepath string) interface{} {
	Josndata := map[string]map[string]map[string]string{} //标准的储存格式
	cmd := exec.Command("python", pyfilepath, string_a)
	message, _ := cmd.Output()
	tempdata := map[string]interface{}{}
	if message != nil {
		err := json.Unmarshal(message, &Josndata)
		//fmt.Println(message)
		if err != nil {
			fmt.Println("解码错误")
		}
		ds := NewSessionStore()
		defer ds.Close()
		c := ds.GetBufferCollection()
		for v, k := range Josndata {
			tempdata[v] = k
			selector := bson.M{v: bson.M{"$exists": true}}
			_, err := c.Upsert(selector, tempdata)
			if err != nil {
				fmt.Println("upsert错误")
			}
		}
	}
	fmt.Println(tempdata)
	return tempdata
}

//GAIJosnData 从message.json中提取并依次插入DB
func GAIJosnData() {
	Josndata := map[string]map[string]map[string]string{} //标准的储存格式
	messagefile, err := os.Open("./message.json")
	if err != nil {
		fmt.Println("Failed to open config file")
		log.Fatal(err)
	}
	fi, _ := messagefile.Stat()
	if fi.Size() == 0 {
		fmt.Println("config file (%q) is empty, skipping")
	}
	message := make([]byte, fi.Size())
	_, err = messagefile.Read(message)
	if err != nil {
		fmt.Println("传递错误")
	}
	err = json.Unmarshal(message, &Josndata)
	if err != nil {
		fmt.Println("解码错误")
	}
	for v, k := range Josndata {
		tempdata := map[string]interface{}{}
		tempdata[v] = k
		InsertToDB(tempdata)
	}
}

func GetFromDB() interface{} {
	ds := NewSessionStore()
	defer ds.Close()
	c := ds.GetStoreCollection()
	size, err := c.Count()
	if err != nil {
		fmt.Println("读取总数失败")
	}
	alldata := make([]map[string]map[string]map[string]string, size)
	c.Find(nil).All(&alldata)
	fmt.Println(alldata)
	fmt.Println("----------------------------------------------------------------------------------------")
	for _, onedata := range alldata {
		fmt.Println(onedata)
	}
	return alldata
}
