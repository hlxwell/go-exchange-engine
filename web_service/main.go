package main

import "github.com/kataras/iris"

func main() {
	app := iris.Default()

	app.Handle("GET", "/", func(ctx iris.Context) {
		ctx.HTML("Hello world!")
	})

	app.Get("/ping", func(ctx iris.Context) {
		ctx.WriteString("pong")
	})

	app.Get("/hello", func(ctx iris.Context) {
		ctx.JSON(iris.Map{"message": "Hello iris web framework."})
	})

	app.Run(iris.Addr(":8080"))
}
