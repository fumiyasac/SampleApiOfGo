package database

import (
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/core"
	"github.com/go-xorm/xorm"
)

var engine *xorm.Engine

// UseEngine ... xormを利用可能にする
func UseEngine() *xorm.Engine {
	if engine == nil {
		engine = initEngine()
	}
	return engine
}

func initEngine() *xorm.Engine {

	// タイムゾーンの設定
	loc, err := time.LoadLocation("Asia/Tokyo")
	if err != nil {
		panic(err)
	}
	time.Local = loc

	// サーバー情報の設定
	dbUser := "root"
	dbPass := "fumiya0921"
	dbHost := "127.0.0.1"
	dbPort := 3306
	dbName := "sample_api"
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", dbUser, dbPass, dbHost, dbPort, dbName)
	e, err := xorm.NewEngine("mysql", dsn)
	if err != nil {
		panic(err)
	}

	// DBへの疎通確認
	e.SetMapper(core.NewCacheMapper(new(core.GonicMapper)))
	err = e.Ping()
	if err != nil {
		panic(err)
	}
	return e
}
