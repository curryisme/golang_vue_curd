package main

import (
	"fmt"
	"github.com/curryisme/golang_vue_curd/config"
	"github.com/curryisme/golang_vue_curd/handler"
	"github.com/curryisme/golang_vue_curd/model"
	"github.com/curryisme/golang_vue_curd/repository"
	"github.com/curryisme/golang_vue_curd/service"
	_ "github.com/go-sql-driver/mysql"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	_ "gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

var (
	DB              *gorm.DB
	UserHandler     handler.UserHandler
)

func initViper() {
	if err := config.Init(""); err != nil {
		panic(err)
	}
}

func initDB() {
	fmt.Println("数据库 init")
	var err error
	conf := &model.DBConf{
		Host:     viper.GetString("database.host"),
		User:     viper.GetString("database.username"),
		Password: viper.GetString("database.password"),
		DbName:   viper.GetString("database.name"),
	}

	config := fmt.Sprintf("%s:%s@tcp(%s)/%s?parseTime=true&charset=utf8&parseTime=%t&loc=%s",
		conf.User,
		conf.Password,
		conf.Host,
		conf.DbName,
		true,
		"Local")
	DB, err = gorm.Open(mysql.Open(config), &gorm.Config{})
	if err != nil {
		log.Fatalf("connect error: %v\n", err)
	}
	//DB.SingularTable(true)
	fmt.Println("数据库 init 结束...")
}

func initHandler() {
	UserHandler = handler.UserHandler{
		UserSrv: &service.UserService{
			Repo: &repository.UserRepository{
				DB: DB,
			},
		}}
}

func init() {
	initViper()
	initDB()
	initHandler()
}
