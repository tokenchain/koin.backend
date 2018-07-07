package main

import (
	"github.com/kataras/iris"
	"github.com/didip/tollbooth"
	"github.com/koinkoin-io/koinkoin.backend/pkg/user"
	"github.com/koinkoin-io/koinkoin.backend/third_party"
	"github.com/koinkoin-io/koinkoin.backend/pkg/auth"
	"github.com/koinkoin-io/koinkoin.backend/pkg/bet"
	"github.com/koinkoin-io/koinkoin.backend/pkg/mail"
	"time"
	"github.com/kataras/iris/context"
)

// RouteAll route all routes from  other service.
func RouteAll(app *iris.Application) {

	app.AllowMethods(iris.MethodOptions)
	app.Use(func(ctx context.Context) {
		ctx.Header("Access-Control-Allow-Origin", "*")
		ctx.Next()
	})

	app.Get("/api/user/new", auth.MidNeedNoAuthentication, user.GetGenerateUser)
	app.Get("/api/user/", auth.MidNeedAuthentication, user.GetUser)
	app.Get("/api/user/mail", auth.MidNeedAuthentication, mail.GetSendMail)
	app.Post("/api/user/update", auth.MidNeedAuthentication, user.PostUpdateUser)
	app.Post("/api/bet", auth.MidNeedAuthentication, bet.PostBet)
	app.Get("/api/bet/stats", bet.GetStats)
	app.Get("/", func(ctx context.Context) {
		ctx.JSON(iris.Map{"online": true, "uptime": Uptime()})
	})
}

func limiter(n float64, duration time.Duration) context.Handler {
	u := tollbooth.NewLimiter(float64(n/float64(duration.Seconds())), nil)
	u.SetMessage("{\"error\": \"forbidden due to rate limiter\"}")
	u.SetMessageContentType("application/json; charset=UTF-8")
	return third.LimitHandler(u)
}
