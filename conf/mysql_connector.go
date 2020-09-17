package conf

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"time"
)

type DbConfig struct {
	DbUser     string `env:"DB_USER"`
	DbPassword string `env:"DB_PASSWORD"`
	DbPort     int    `env:"DB_PORT" envDefault:"3306"`
	DbHost     string `env:"DB_HOST"`
	DbName     string `env:"DB_NAME"`
}

var DB *sql.DB

func InitializeDb(user string, password string, host string, name string, port int) {
	var err error
	dbURL := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local", user, password, host, port, name)
	DB, err = sql.Open("mysql", dbURL)
	if err != nil {
		fmt.Printf("Cannot connect to database")
		log.Fatal("error:", err)
	}
	if err = DB.Ping(); err != nil {
		fmt.Println("DB dead")
	}
	DB.SetConnMaxLifetime(time.Minute * 3)
	DB.SetMaxOpenConns(10)
	DB.SetMaxIdleConns(10)
	fmt.Printf("We are connected to database \n")

	//defer DB.Close()
}
