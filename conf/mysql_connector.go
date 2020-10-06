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

func InitializeDb(user string, password string, host string, name string, port int) error {
	var dbErr error
	dbURL := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local", user, password, host, port, name)
	db, err := sql.Open("mysql", dbURL)
	if err != nil {
		fmt.Printf("Cannot connect to database")
		log.Fatal("error:", err)
		return err
	}
	for i := 1; i <= 8; i++ {
		dbErr = db.Ping()
		if dbErr != nil {
			if i < 8 {
				log.Printf("db connection failed, %d retry : %v", i, dbErr)
				time.Sleep(10 * time.Second)
			}
			continue
		}

		break
	}
	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)
	fmt.Printf("We are connected to database \n")
	DB = db
	//defer DB.Close()
	return nil
}
