package model

import (
	"github.com/go-ini/ini"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

type ConfigList struct {
	DBMS     string
	USER     string
	PASS     string
	PROTOCOL string
	DBNAME   string
}

var Config ConfigList

func LoadIni() {
	cfg, _ := ini.Load("config.ini")
	Config = ConfigList{
		DBMS:     cfg.Section("db").Key("dbms").String(),
		USER:     cfg.Section("db").Key("user").String(),
		PASS:     cfg.Section("db").Key("pass").String(),
		PROTOCOL: cfg.Section("db").Key("protocol").String(),
		DBNAME:   cfg.Section("db").Key("dbname").String(),
	}
}

type Todo struct {
	gorm.Model
	Title       string
	Description string
}

func ConnectDB() *gorm.DB {
	LoadIni()
	dbms := Config.DBMS
	user := Config.USER
	pass := Config.PASS
	protocol := Config.PROTOCOL
	dbname := Config.DBNAME
	connect := user + ":" + pass + "@" + protocol + "/" + dbname + "?parseTime=true&loc=Asia%2FTokyo"

	db, err := gorm.Open(dbms, connect)
	if err != nil {
		panic(err)
	}

	return db
}

func InitDB() {
	db := ConnectDB()
	defer db.Close()
	db.AutoMigrate(&Todo{})
}

func GetAll() []Todo {
	db := ConnectDB()
	defer db.Close()
	var todos []Todo
	db.Find(&todos)

	return todos
}

func GetOne(id int) Todo {
	db := ConnectDB()
	defer db.Close()
	var todo Todo
	db.First(&todo, id)

	return todo
}

func CreateTodo(title string, description string) {
	db := ConnectDB()
	defer db.Close()
	db.Create(&Todo{Title: title, Description: description})
}

func UpdateTodo(id int, title string, description string) {
	db := ConnectDB()
	defer db.Close()
	var todo Todo
	db.First(&todo, id)
	todo.Title = title
	todo.Description = description
	db.Save(&todo)
}

func DeleteTodo(id int) {
	db := ConnectDB()
	defer db.Close()
	var todo Todo
	db.First(&todo, id)
	db.Delete(&todo)
}
