package auth

import (
	"github.com/kataras/iris"
	"github.com/koinkoin-io/koinkoin.backend/pkg/err"
	"fmt"
)

// MidNeedAuthentication is a middleware that check if the header contain a
// valid user hash in the header.
func MidNeedAuthentication(ctx iris.Context) {
	hash := ctx.GetHeader("hash")
	if New().Auth(hash) {
		ctx.Next()
	} else {
		ctx.StatusCode(iris.StatusForbidden)
		err.ThrownError(ctx, err.NotAuthenticated)
	}
}

// MidNeedNoAuthentication is a middleware that check if an user is not logged.
func MidNeedNoAuthentication(ctx iris.Context) {
	hash := ctx.GetHeader("hash")
	fmt.Printf("HOOOOO YA PAS DE HASH FDP :" + ctx.GetHeader("hash"))
	if hash == "" || !New().Auth(hash) {
		ctx.Next()
	} else {
		ctx.StatusCode(iris.StatusForbidden)
		err.ThrownError(ctx, err.AlreadyAuthenticated)
	}
}