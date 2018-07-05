package user

import (
	"github.com/kataras/iris"
	"../err"
	"../auth"
)

// GenerateUser generate a new user and return the json of the User struct.
// Need to not be authenticated.
func GetGenerateUser(ctx iris.Context) {
	u := New()
	u.Save()
	ctx.JSON(u)
}

// PostUpdateUser update Name of an user or Email.
// Regular expression are in user.go.
// If an error occur return a text with the error description.
// Else update the state of user and safe.
//
// Note: you can update name and fail with mail.
func PostUpdateUser(ctx iris.Context) {
	if !updateValue(ctx, "name", func(str string) bool {
		return NameRegexp.MatchString(str)
	}, func(user *User, str string) {
		user.Name = str
	}, err.IncorrectName) {
		return
	}

	if !updateValue(ctx, "mail", func(str string) bool {
		return MailRegexp.MatchString(str)
	}, func(user *User, str string) {
		user.Mail = str
	}, err.IncorrectMail) {
		return
	}
	ctx.JSON(iris.Map{"success": true})
}

func GetUser(ctx iris.Context) {
	if auth.New().Auth(ctx.URLParam("hash")) {
		ctx.JSON(Get(ctx.URLParam("hash")))
		return
	}
	err.ThrownError(ctx, err.NoUserFound)
}

// updateValue is just boilerplate because go.
func updateValue(ctx iris.Context, name string, validation func(str string) bool,
	then func(user *User, str string), er *err.Err) bool {
	if x := ctx.PostValue(name); x != "" {
		if validation(x) {
			u := Get(ctx.GetHeader("hash"))
			then(u, x)
			u.Save()
			return true
		} else {
			err.ThrownError(ctx, er)
			return false
		}
	}
	return true
}


