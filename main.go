package main

import (
	"github.com/JackMaarek/DS/conf"
	routes "github.com/JackMaarek/DS/router"
	"github.com/caarlos0/env"
	"github.com/gin-gonic/gin"
	"log"
	"os"
)

func main() {
	cfg := conf.DbConfig{}
	if err := env.Parse(&cfg); err != nil {
		log.Fatal(err)
	}
	err := conf.InitializeDb(cfg.DbUser, cfg.DbPassword, cfg.DbHost, cfg.DbName, cfg.DbPort)
	if err != nil {
		log.Println("Cannot connect to database project shutting")
		os.Exit(1)
	}
	router := gin.Default()
	routes.SetupRouter(router)

	log.Fatal(router.Run(":8080"))
}
