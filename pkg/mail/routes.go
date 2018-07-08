package mail

import (
	"github.com/kataras/iris"
	"github.com/koinkoin-io/koinkoin.backend/pkg/user"
	"github.com/koinkoin-io/koinkoin.backend/pkg/err"
)



func GetSendMail(ctx iris.Context) {
	if ctx.URLParam("mail") == "" || !user.MailRegexp.MatchString(ctx.URLParam("mail")) {
		err.ThrownError(ctx, err.IncorrectMail)
		return
	}
	hash := ctx.GetHeader("hash")
	go SendMail(hash, ctx.URLParam("mail"))
}