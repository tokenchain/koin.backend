package mail

import (
	"github.com/kataras/iris"
	"github.com/koinkoin-io/koinkoin.backend/pkg/user"
	"github.com/koinkoin-io/koinkoin.backend/pkg/err"
	"github.com/koinkoin-io/koinkoin.backend/pkg/worker"
)



func GetSendMail(ctx iris.Context) {
	if ctx.URLParam("mail") == "" || !user.MailRegexp.MatchString(ctx.URLParam("mail")) {
		err.ThrownError(ctx, err.IncorrectMail)
		return
	}
	hash := ctx.GetHeader("hash")

	worker.PushJob(worker.Job{
		Name:     "mail for hash " + ctx.GetHeader("hash") + " at " + ctx.URLParam("mail"),
		Runnable: func() error {
			SendMail(hash, ctx.URLParam("mail"))
			return nil
		},
	})
}