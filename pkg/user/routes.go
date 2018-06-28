package user

import "github.com/kataras/iris"

// GenerateUser generate a new user and return the json of the User struct.
// Need to not be authenticated.
func GetGenerateUser(ctx iris.Context) {
	u := New()
	u.Save()
	ctx.JSON(u)
}