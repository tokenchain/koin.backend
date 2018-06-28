package bet

import (
	"math/rand"
	"time"
	"../err"
)

// BetService declare the function Bet that is a function to bet coins.
type BetService interface {
	Bet(coins uint64, chance int) (result int, earn uint64, win bool, error error)
}

// Bet is stateless structure to implement BetService.
type Bet struct{}

// New return a new structure of type Bet that implement BetService.
func New() Bet {
	return Bet{}
}

// random generate a random number between min and max.
// seed is defined with time.Now().UTC().UnixNano().
func random(min, max int) int {
	rand.Seed(time.Now().UTC().UnixNano())
	return rand.Intn(max-min) + min
}

// Bet check:
//  - if chance is >= to 98 -> Error
//  - if chance is <= 2 -> Error
//  - if coins < 5 -> Error
// If all conditions are passed, generate a random number, if this number is lesser than chance, the bettor gain
// (1-(chance/100)) * coins.
// Else the bettor lose all coins.
func (c Bet) Bet(coins uint64, chance int) (result int, earn uint64, win bool, error error) {
	if chance >= 98 {
		return 0, 0, false, err.ChanceCantBeEqOrHigher98
	} else if chance <= 2 {
		return 0, 0, false, err.ChanceCantBeLesser2
	} else if coins < 5 {
		return 0, 0, false, err.CoinsCantBeLesser5
	}
	random := random(0, 100)
	if random <= chance {
		return random, uint64(float64(coins)*float64(1.-(float64(chance)/100.))), true, nil
	}
	return random, coins, false, nil
}
