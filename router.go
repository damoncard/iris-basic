package main

import (
	"./function"
	"gopkg.in/kataras/iris.v6"
	"gopkg.in/kataras/iris.v6/adaptors/httprouter"
	"gopkg.in/kataras/iris.v6/adaptors/sessions"
	"gopkg.in/kataras/iris.v6/adaptors/view"
)

func main() {
	app := iris.New()
	app.Adapt(
		httprouter.New(),
		view.HTML("./view", ".html").Reload(true),
		sessions.New(sessions.Config{Cookie: "yummy"}),
	)
	app.StaticWeb("/static", "./assets")

	app.Get("/", func(ctx *iris.Context) {
		ctx.Render("index.html", nil)
	})

	app.Get("/register", func(ctx *iris.Context) {
		ctx.Render("register.html", nil)
	})

	app.Post("/register", func(ctx *iris.Context) {
		function.AddUser(ctx.PostValue("username"), ctx.PostValue("password"))
		ctx.Redirect("/")
	})
	
	app.Post("/login", function.Login)
	app.Get("/logout", function.Logout)

	app.Get("/forum/:name", func(ctx *iris.Context) {
		if checkAuthen(ctx) {
			ctx.Redirect("/")
			return
		}

		admin := false

		if ctx.Param("name") == "admin" {
			admin = true
		}

		data := []function.Data { function.ReadJSON(admin) }
		ctx.Render("todo.html",  struct { Datas []function.Data }{data})
	})

	app.Get("/edit/:n", func(ctx *iris.Context) {
		if checkAuthen(ctx) {
			ctx.Redirect("/")
			return
		}

		data := []function.Task { function.ReadJSONbyID(ctx.Param("n")) }
		ctx.Render("editing.html",  struct { Tasks []function.Task }{data})
	})

	app.Post("/add", func(ctx *iris.Context) {
		if checkAuthen(ctx) {
			ctx.Redirect("/")
			return
		}

		t:= setTask(ctx)
		function.AddJSON(t)
		ctx.Redirect("/forum/admin")
	})

	app.Post("/edit", func(ctx *iris.Context) {
		if checkAuthen(ctx) {
			ctx.Redirect("/")
			return
		}

		t:= setTask(ctx)
		function.EditJSON(t)
		ctx.Redirect("/forum/admin")
	})

	app.Post("/delete", func(ctx *iris.Context) {
		if checkAuthen(ctx) {
			ctx.Redirect("/")
			return
		}

		function.DeleteTask(ctx.PostValue("delete"))
		ctx.Redirect("/forum/admin")
	})

	println("Server is running at :7000")
	app.Listen(":7000")
}

func checkAuthen(ctx *iris.Context) bool {
	if auth, _ := ctx.Session().GetBoolean("authenticated"); !auth {
		return true
	} else {
		return false
	}
}

func setTask(ctx *iris.Context) function.Task {
	name := ctx.PostValue("title")
	date := ctx.PostValue("date")
	status := ctx.PostValue("status")
	t := function.Task {
		Name: name,
		Date: date,
		Done: status,
	}
	return t
}