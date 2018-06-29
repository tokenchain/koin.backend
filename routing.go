package main

import (
	"github.com/kataras/iris"
	"github.com/didip/tollbooth"
	"./pkg/user"
	"./third_party"
	"./pkg/auth"
	"./pkg/bet"
	"time"
	"github.com/kataras/iris/context"
)

// RouteAll route all routes from  other service.
func RouteAll(app *iris.Application) {

	app.Get("/api/user/new", limiter(1, time.Second), auth.MidNeedNoAuthentication, user.GetGenerateUser)
	app.Post("/api/bet", limiter(2, time.Second), auth.MidNeedAuthentication, bet.PostBet)
}

func limiter(n float64, duration time.Duration) context.Handler {
	u := tollbooth.NewLimiter(float64(n/float64(duration.Seconds())), nil)
	u.SetMessage("{\"error\": \"forbidden due to rate limiter\"}")
	u.SetMessageContentType("application/json; charset=UTF-8")
	return third.LimitHandler(u)
}