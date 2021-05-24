package main

import (
	"log"
	"project/router"
	"project/model"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func main() {

	r := gin.Default()

	model.Db.Init()
	defer model.Db.Close()

	router.Router(r)
	if err := r.Run(); err != nil {
		log.Fatal(err.Error())
	}
}
