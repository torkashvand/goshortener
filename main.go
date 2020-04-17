package main

import (
	"github.com/torkashvand/goshortener/cmd"
	"github.com/torkashvand/goshortener/models"
	"github.com/torkashvand/goshortener/routers"
)

func main() {

	mysql := &models.MysqlDB{}
	mysql.Open()
	mysql.AutoMigrate()
	defer mysql.Close()

	router := routers.SetupRouter(mysql)
	router.Run()

	cmd.Execute()
}
