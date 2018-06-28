package bet

import (
	"github.com/kataras/iris"
	"strconv"
	"../err"
	"../user"
	"fmt"
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
	fmt.Println(us.HasEnoughCoin(coins))
	if !us.HasEnoughCoin(coins) {
		no(ctx, err.NotEnoughCoins)
		return
	}
	res, coins, win, er := New().Bet(coins, chance)

	// Check if bet didn't throw an error.
	if er != nil {
		no(ctx, er)
		return
	}

	// If win add coins else remove coins
	if win {
		us.Coins += coins
	} else {
		us.Coins -= coins
	}

	us.Save()
	ctx.JSON(iris.Map{"result": res, "earn": coins, "win": win, "coins": us.Coins})
}

// no just set statuscode and json for an error.
func no(ctx iris.Context, err error) {
	ctx.StatusCode(iris.StatusBadRequest)
	ctx.JSON(iris.Map{"error": err.Error()})
}