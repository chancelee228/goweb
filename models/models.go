package models

import (
	"time"
	"github.com/astaxie/beego/orm"
	_ "github.com/mattn/go-sqlite3"
	"github.com/Unknwon/com"
	"os"
	"path"
	"strconv"
)
const(
	_DB_NAME ="data/beeblog.db"
	_SQLITE3_DRIVER = "sqlite3"
)
type category struct {
	Id 	int64
	Title 	string
	Created time.Time `orm:"index"`
	Views 	int64	  `orm:"index"`
	TopicTime time.Time `orm:"index"`
	TopicCount int64
	TopicLastUserId int64
}
type Topic struct {
	Id int64
	Uid int64
	Title string
	Content string `orm:"size(5000)"`
	Attachment string
	Created time.Time `orm:"index"`
	Updated time.Time `orm:"index"`
	Views 	int64	  `orm:"index"`
	Author string
	ReplyTime time.Time `orm:"index"`
	ReplyCount int64
	ReplyLastUserId int64
}
func RegisterDB(){
	if !com.IsExist(_DB_NAME){
		os.MkdirAll(path.Dir(_DB_NAME),os.ModePerm)
		os.Create(_DB_NAME)
	}
	//注册模型
	orm.RegisterModel(new(category),new(Topic))
	//注册驱动
	orm.RegisterDriver(_SQLITE3_DRIVER,orm.DRSqlite)
	//注册数据库
	orm.RegisterDataBase("default",_SQLITE3_DRIVER,_DB_NAME,10)//最大连接数
}
func AddCategory(name string) error{
	o:=orm.NewOrm()
	category:=&category{Title:name,Created:time.Now(),TopicTime:time.Now()}
	qs:=o.QueryTable("category")
	err:=qs.Filter("title",name).One(category)
	if nil == err{
		return *new(error)
	}
	_,err = o.Insert(category)
	if nil != err{
		return err
	}
	return nil
}

func DelCategoryById(id string) error {
	cid,err :=strconv.ParseInt(id,10,64)
	if err != nil {
		return err
	}
	o:=orm.NewOrm()
	category:=&category{Id:cid}
	_,err=o.Delete(category)
	return err
}

func GetAllCategories() ([]*category,error){
	o:=orm.NewOrm()
	categories:=make([]*category,0)
	qs:=o.QueryTable("category")
	_,err := qs.All(&categories)
	return categories,err

}
func GetAllTopics(isDesc bool)([]*Topic,error){
	o:=orm.NewOrm()
	topics:=make([]*Topic,0)
	qs:=o.QueryTable("topic")
	var err error
	if isDesc{
		_,err = qs.OrderBy("-created").All(&topics)
	}else {
		_,err = qs.All(&topics)
	}
	return topics,err
}

func AddTopic(title,content string )error{
	o:=orm.NewOrm()
	topic:=&Topic{
		Title:title,
		Content:content,
		Created:time.Now(),
		Updated:time.Now(),
		ReplyTime:time.Now(),
	}
	_,err:=o.Insert(topic)
	return err
}

func GetTopicById(id string) (*Topic,error){
	tidNum,err:=strconv.ParseInt(id,10,64)
	if nil != err{
		return nil,err
	}
	o:=orm.NewOrm()
	topic:=new(Topic)
	qs:=o.QueryTable("topic")
	err = qs.Filter("id",tidNum).One(topic)
	if nil != err{
		return nil,err
	}
	topic.Views++
	_,err = o.Update(topic)
	return topic,err
}