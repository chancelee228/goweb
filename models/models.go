package models

import (
	"time"
	"github.com/astaxie/beego/orm"
	_ "github.com/mattn/go-sqlite3"
	"github.com/Unknwon/com"
	"os"
	"path"
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