package config

import (
	"core/system"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
)

type Database struct {
	Driver       string `yaml:"driver"`
	Url          string `yaml:"url"`
	Username     string `yaml:"username"`
	Password     string `yaml:"password"`
	MaxIdleConns int    `yaml:"maxIdleConns"`
	MaxOpenConns int    `yaml:"maxOpenConns"`
	ShowSQL      bool   `yaml:"showSQL"`
}

func (e *Database) Init() {
	if e == nil {
		return
	}
	engine, err := xorm.NewEngine(e.Driver, fmt.Sprintf("%s:%s@%s", e.Username, e.Password, e.Url))
	if err != nil {
		fmt.Println(err.Error())
	}
	err = engine.Ping()
	if err != nil {
		fmt.Printf("connect ping failed: %v", err)
		return
	}
	//最大空闲连接数
	engine.SetMaxIdleConns(e.MaxIdleConns)
	//最大连接数
	engine.SetMaxOpenConns(e.MaxOpenConns)
	//是否打印SQL
	engine.ShowSQL(e.ShowSQL)
	system.SetXormEngine(engine)
}
