package bet

import (
	"github.com/kataras/iris"
	"strconv"
	"../err"
	"../user"
	"../auth"
)

// Bet is a http handler that invoke New().Bet method with full error check.
func PostBet(ctx iris.Context) {
	coins, errCoins := strconv.ParseUint(ctx.FormValue("coins"), 10, 64)
	chance, errChance := strconv.Atoi(ctx.FormValue("chance"))

	// Check if conversion int are okay
	if errCoins != nil  || errChance != nil {
		no(ctx, err.NumberMalformed)
		return
	}

	us := user.Get(ctx.GetHeader("hash"))

	// Check if us has enough coins
	if !us.HasEnoughCoin(coins) {
		no(ctx, err.NotEnoughCoins)
		return
	}

	bet := New(coins, chance, us.Coins)
	er := bet.Bet()
	bet.BeforeCoins = us.Coins
	// Check if bet didn't throw an error.
	if er != nil {
		no(ctx, er)
		return
	}

	// If win add coins else remove coins
	if bet.Win {
		us.Coins += bet.Earn
	} else {
		us.Coins -= bet.Earn
	}


	us.Save()
	bet.AfterCoins = us.Coins
	globalStats.UpdateStatistics(*bet).save()
	NewStats(us.Hash).UpdateStatistics(*bet).save()
	ctx.JSON(bet)
}

func GetStats(ctx iris.Context) {
	hash := ctx.URLParam("hash")
	if hash != "" && auth.New().Auth(hash) || hash == "global" {
		ctx.JSON(NewStats(hash))
		return
	}
	no(ctx, err.IncorrectParameter)
}

// no just set statuscode and json for an error.
func no(ctx iris.Context, err error) {
	ctx.StatusCode(iris.StatusBadRequest)
	ctx.JSON(iris.Map{"error": err.Error()})
}