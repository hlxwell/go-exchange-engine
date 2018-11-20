package main

import (
	"fmt"
	"math/rand"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/kataras/iris"
	// 	"github.com/labstack/echo"
)

type User struct {
	gorm.Model
	ID    int `gorm:"AUTO_INCREMENT"`
	Email string
}

type Account struct {
	gorm.Model
	ID               int `gorm:"AUTO_INCREMENT"`
	Currency         string
	Balance          float64
	AvailableBalance float64
}

func main() {
	app := iris.Default()

	// create db connection
	db, err := gorm.Open("postgres", "host=127.0.0.1 port=5432 user=helixing password= dbname='go-exchange-engine' sslmode=disable")
	defer db.Close()
	// db.Model(user).Related(&account)

	if err != nil {
		fmt.Println(err)
		panic("Failed to connect to database.")
	}
	db.AutoMigrate(&User{})

	app.Handle("GET", "/", func(ctx iris.Context) {
		email := fmt.Sprintf("hlxwell%d@gmail.com", rand.Intn(100))
		db.Create(&User{Email: email})

		var users []User
		db.Where("email = ?", "hlxwell1@gmail.com").Find(&users)
		// db.Delete(users)

		fmt.Println(email)
		for _, user := range users {
			fmt.Println(user.Email)
		}

		defer ctx.HTML("Hello world!")
	})

	app.Get("/ping", func(ctx iris.Context) {
		ctx.WriteString("pong")
	})

	app.Get("/hello", func(ctx iris.Context) {
		ctx.JSON(iris.Map{"message": "Hello iris web framework."})
	})

	app.Run(iris.Addr(":8000"))
}
