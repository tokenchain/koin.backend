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
	"github.com/iris-contrib/middleware/cors"
)

// RouteAll route all routes from  other service.
func RouteAll(app *iris.Application) {
	crs := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"}, // allows everything, use that to change the hosts.
		AllowCredentials: true,
	})

	a := app.Party("/api", crs).AllowMethods(iris.MethodOptions)
	a.Get("/user/new", auth.MidNeedNoAuthentication, user.GetGenerateUser)
	a.Get("/user/", auth.MidNeedAuthentication, user.GetUser)
	a.Get("/user/mail", auth.MidNeedAuthentication, mail.GetSendMail)
	a.Post("/user/update", auth.MidNeedAuthentication, user.PostUpdateUser)
	a.Post("/bet", auth.MidNeedAuthentication, bet.PostBet)
	a.Get("/bet/stats", bet.GetStats)
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
