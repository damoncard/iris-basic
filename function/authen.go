package function

import (
	"gopkg.in/kataras/iris.v6"
)

func Login(ctx *iris.Context) {
	session := ctx.Session()
 
	user, level := CheckUser(ctx.PostValue("username"), ctx.PostValue("password"))

	if user {
		session.Set("authenticated", true)
		ctx.Redirect("/forum/" + level)
		return
	}
	
	ctx.Redirect("/")
}

func Logout(ctx *iris.Context) {
	session := ctx.Session()
	session.Set("authenticated", false)
	ctx.Redirect("/")
}