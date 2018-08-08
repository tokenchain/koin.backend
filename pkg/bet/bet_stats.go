package bet

import (
	"github.com/koin-bet/koin.backend/pkg/db"
	"math"
)

var globalStats = NewStats("global")

// Statistics represent the globals stats of a player when he bet.
type Statistics struct {
	Hash  string `json:"hash"`  // Hash is the identifier.
	Count uint64 `json:"count"` // Count is the total bet effectuated.

	AverageEarn   uint64 `json:"averageEarn"`   // AverageEarn is the average of earn.
	AverageLose   uint64 `json:"averageLose"`   // AverageLose is the average of lose.
	AverageCoins  uint64 `json:"averageCoins"`  // AverageCoins is the average of coins bet.
	AverageChance uint64 `json:"averageChance"` // AverageChance is the average of chance bet.
	AverageResult uint64 `json:"averageResult"` // AverageResult is the average of result number generated.

	MaxAmount uint64 `json:"maxAmount"` // MaxAmount is the max coins amount bet.
	MinAmount uint64 `json:"minAmount"` // MinAmount is the min coins amount bet.

	MaxChance uint64 `json:"maxChance"` // MaxChance is the highest chance bet.
	MinChance uint64 `json:"minChance"` // MinChance is the lowest chance bet.

	MaxEarn uint64 `json:"maxEarn"` // MaxEarn is the maximum bettor earn.
	MinEarn uint64 `json:"minEarn"` // MinEarn is the minimum a bettor earn.

	MaxLose uint64 `json:"maxLose"` // MaxLose is the maximum bettor lose.
	MinLose uint64 `json:"minLose"` // MinLose is the minimum a bettor lose.

	Success uint64 `json:"success"` // Success is the amount of positive bet.
	Failed  uint64 `json:"failed"`  // Failed is the amount of negative bet.

	TotalCoins  uint64 `json:"totalCoins"`  // TotalCoins is the total chance bet.
	TotalEarn   uint64 `json:"totalEarn"`   // TotalEarn is the total a bettor earn.
	TotalLose   uint64 `json:"totalLose"`   // TotalLose is the total a bettor lose.
	TotalChance uint64 `json:"totalChance"` // TotalChance is the total a bettor chance.
	TotalResult uint64 `json:"totalResult"` // TotalResult is the total of result obtained.

	Greedy  uint64 `json:"greedy"`  // Greedy represent the bettor as 'greedy'.
	Fearful uint64 `json:"fearful"` // Fearful  represent the bettor as 'fearful'.
}

// UpdateStatistics compute statistics from a bet structure.
func (s *Statistics) UpdateStatistics(bet Bet) *Statistics {
	s.Count++

	s.TotalCoins += bet.Coins
	s.TotalChance += uint64(bet.Chance)
	s.TotalResult += uint64(bet.Result)

	if bet.Win {
		s.Success++
		s.TotalEarn += bet.Earn
	} else {
		s.Failed++
		s.TotalLose += bet.Earn
	}
	s.updateMin(bet)
	s.updateMax(bet)
	s.updateAverage(bet)
	s.updateGreedy(bet)
	s.updateFearful(bet)
	return s
}

// updateMin compute only min fields.
func (s *Statistics) updateMin(bet Bet) {
	if s.MinAmount > bet.Coins {
		s.MinAmount = bet.Coins
	}
	if bet.Win && s.MinEarn > bet.Earn {
		s.MinEarn = bet.Earn
	} else if s.MinLose > bet.Earn {
		s.MinLose = bet.Earn
	}
	if s.MinChance > uint64(bet.Chance) {
		s.MinChance = uint64(bet.Chance)
	}
}

// updateMax compute only max fields.
func (s *Statistics) updateMax(bet Bet) {
	if s.MaxAmount < bet.Coins {
		s.MaxAmount = bet.Coins
	}
	if bet.Win && s.MaxEarn < bet.Earn {
		s.MaxEarn = bet.Earn
	} else if s.MaxLose < bet.Earn {
		s.MaxLose = bet.Earn
	}
	if s.MaxChance < uint64(bet.Chance) {
		s.MaxChance = uint64(bet.Chance)
	}
}

// updateAverage compute only average fields.
func (s *Statistics) updateAverage(bet Bet) {
	s.AverageCoins = uint64(float64(s.TotalCoins) / float64(s.Count))
	if s.Success != 0 {
		s.AverageEarn = uint64(float64(s.TotalEarn) / float64(s.Success))
	}
	s.AverageChance = uint64(float64(s.TotalChance) / float64(s.Count))
	s.AverageResult = uint64(float64(s.TotalResult) / float64(s.Count))
	if s.Failed != 0 {
		s.AverageLose = uint64(float64(s.TotalLose) / float64(s.Failed))
	}
}

// updateGreedy update the greedy state.
// The greedy score of a player is determined by theses factors:
// - The amount of coins bet.
// - The percentage of chance he bet.
// More the risk is high more the score will be incremented.
func (s *Statistics) updateGreedy(bet Bet) {
	if bet.TotalCoinsBefore == bet.Coins && bet.Chance < 15 {
		s.Greedy += 4
	} else if bet.TotalCoinsBefore == bet.Coins {
		s.Greedy += 3
	} else if bet.Coins/bet.TotalCoinsBefore*100 > 85 && bet.Chance < 15 {
		s.Greedy += 3
	} else if bet.Coins/bet.TotalCoinsBefore*100 > 85 {
		s.Greedy += 2
	} else if bet.Chance < 15 {
		s.Greedy += 1
	}
}

// updateFearful update the fearful state.
// The fearful score of a player is determined by theses factors:
// - The amount of coins bet.
// - The percentage of chance he bet.
// More the risk is low more the score will be incremented.
func (s *Statistics) updateFearful(bet Bet) {
	if bet.Coins/bet.TotalCoinsBefore*100 < 10 && bet.Chance > 80 {
		s.Greedy += 4
	} else if bet.Coins/bet.TotalCoinsBefore*100 < 10 {
		s.Greedy += 2
	} else if bet.Chance > 80 {
		s.Greedy += 1
	}
}

// NewStats create a new Statistics structure if no key
// was found in the database client.
// If a key was found, load it.
func NewStats(key string) *Statistics {
	stats := &Statistics{
		Hash:      key,
		MinLose:   math.MaxUint64,
		MinChance: math.MaxUint64,
		MinEarn:   math.MaxUint64,
		MinAmount: math.MaxUint64,
	}
	exist, err := db.GetDb().HKeys("stats.bet." + stats.Hash)
	if len(exist) > 0 && err == nil {
		db.StructFromKey("stats.bet."+stats.Hash, stats)
	}
	return stats
}

// save save the stats on the database.
func (s Statistics) save() {
	db.SaveStructure("stats.bet."+s.Hash, s)
}
