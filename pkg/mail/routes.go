package mail

import (
	"github.com/kataras/iris"
	"../user"
	"../err"
)



func GetSendMail(ctx iris.Context) {
	if ctx.URLParam("mail") == "" || !user.MailRegexp.MatchString(ctx.URLParam("mail")) {
		err.ThrownError(ctx, err.IncorrectMail)
		return
	}
	hash := ctx.GetHeader("hash")
	SendMail(hash, ctx.URLParam("mail"))

}