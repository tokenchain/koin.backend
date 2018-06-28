package auth

import (
	"github.com/kataras/iris"
	"../err"
)

// MidNeedAuthentication is a middleware that check if the header contain a
// valid user hash in the header.
func MidNeedAuthentication(ctx iris.Context) {
	hash := ctx.GetHeader("hash")
	if New().Auth(hash) {
		ctx.Next()
	} else {
		ctx.StatusCode(iris.StatusForbidden)
		ctx.JSON(iris.Map{"error": err.NotAuthenticated.Error()})
	}
}

// MidNeedNoAuthentication is a middleware that check if an user is not logged.
func MidNeedNoAuthentication(ctx iris.Context) {
	hash := ctx.GetHeader("hash")
	if !New().Auth(hash) {
		ctx.Next()
	} else {
		ctx.StatusCode(iris.StatusForbidden)
		ctx.JSON(iris.Map{"error": err.AlreadyAuthenticated.Error()})
	}
}