package main

import (
	"log"
	"project/model"
	"project/router"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func main() {

	model.Db.Init()
	defer model.Db.Close()

	r := gin.Default()
	router.Router(r)
	if err := r.Run(); err != nil {
		log.Fatal(err.Error())
	}
	r.Run() // listen and serve on 0.0.0.0:8080
}
