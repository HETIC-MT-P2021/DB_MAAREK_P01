package main

import (
	"fmt"
	"github.com/JackMaarek/DS/conf"
	routes "github.com/JackMaarek/DS/router"
	"github.com/caarlos0/env"
	"github.com/gin-gonic/gin"
	"log"
)

func main() {
	cfg := conf.DbConfig{}
	if err := env.Parse(&cfg); err != nil {
		log.Fatal(err)
	}
	fmt.Println(cfg)
	conf.InitializeDb(cfg.DbUser, cfg.DbPassword, cfg.DbHost, cfg.DbName, cfg.DbPort)
	router := gin.Default()
	routes.SetupRouter(router)

	fmt.Println(conf.DB)

	log.Fatal(router.Run(":8080"))
}
