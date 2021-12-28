package models

import (
	"database/sql"
	"fmt"
	"github.com/gin-blog/pkg/setting"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

var db *sql.DB


type Model struct {
	ID int `gorm:"primary_key" json:"id"`
	CreatedOn int `json:"created_on"`
	ModifiedOn int `json:"modified_on"`
}


func init()  {
	sec,err := setting.Cfg.GetSection("database")
	if err != nil{
		log.Fatal(2, "Fail to get section 'database': %v", err)
	}

	dbType := sec.Key("TYPE").String()
	dbName := sec.Key("NAME").String()
	user := sec.Key("USER").String()
	password := sec.Key("PASSWORD").String()
	host := sec.Key("HOST").String()
	//tablePrefix := sec.Key("TABLE_PREFIX").String()

	// fmt.Sprintf 将变量格式化为str输出
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local",
			user,
			password,
			host,
			dbName,
			)
	// 定义db
	db,err = sql.Open(dbType,dsn)
	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()
	db.SetMaxIdleConns(10)
	db.SetMaxOpenConns(100)



}
