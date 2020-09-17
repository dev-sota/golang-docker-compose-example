package main

import (
	"log"
	"net/http"
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	db := NewDB(30)
	defer db.Close()

	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/", hello)

	e.Logger.Fatal(e.Start(":9000"))
}

func hello(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}

func NewDB(count uint) *gorm.DB {
	dsn := "root:@tcp(mysql:3306)/example_db"
	db, err := gorm.Open("mysql", dsn)
	if err != nil {
		if count == 0 {
			log.Fatalln(err)
		}
		time.Sleep(time.Second)
		count--
		log.Printf("%v %v", "Retry connecting...", count)
		return NewDB(count)
	}
	db.LogMode(true)
	return db
}
