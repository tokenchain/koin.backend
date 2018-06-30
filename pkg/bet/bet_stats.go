package bet

import "../db"

var globalStats = NewStats("global")

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

	TotalCoins  uint64 `json:"totalCoins"` // TotalCoins is the total chance bet.
	TotalEarn   uint64 `json:"totalEarn"`   // TotalEarn is the total a bettor earn.
	TotalLose   uint64 `json:"totalLose"`   // TotalLose is the total a bettor lose.
	TotalChance uint64 `json:"totalChance"` // TotalChance is the total a bettor chance.
	TotalResult uint64 `json:"totalResult"` // TotalResult is the total of result obtained.

	Greedy  uint64 `json:"isGreedy"`  // IsGreedy represent the bettor as 'greedy'.
	Fearful uint64 `json:"isFearful"` // IsFearful represent the bettor as 'fearful'.
}

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
	//todo update Greedy and Fearful ? How to determine it ?
	return s
}

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
func (s *Statistics) updateAverage(bet Bet) {
	s.AverageCoins = uint64(float64(s.TotalCoins) / float64(s.Count))
	s.AverageEarn = uint64(float64(s.TotalEarn) / float64(s.Success))
	s.AverageChance = uint64(float64(s.TotalChance) / float64(s.Count))
	s.AverageResult = uint64(float64(s.TotalResult) / float64(s.Count))
	s.AverageLose = uint64(float64(s.TotalLose) / float64(s.Failed))
}

func (s *Statistics) updateGreedy(bet Bet) {
	if bet.TotalCoinsBefore == bet.Coins && bet.Chance < 15 {
		s.Greedy += 4
	} else if bet.TotalCoinsBefore == bet.Coins {
		s.Greedy += 3
	} else if bet.Coins / bet.TotalCoinsBefore * 100 > 85 && bet.Chance < 15 {
		s.Greedy += 3
	} else if bet.Coins / bet.TotalCoinsBefore * 100 > 85 {
		s.Greedy += 2
	} else if bet.Chance < 15 {
		s.Greedy += 1
	}
}

func (s *Statistics) updateFearful(bet Bet) {
	if bet.Coins / bet.TotalCoinsBefore * 100 < 10 && bet.Chance > 80 {
		s.Greedy += 4
	} else if bet.Coins / bet.TotalCoinsBefore * 100 < 10 {
		s.Greedy += 2
	} else if bet.Chance > 80 {
		s.Greedy += 1
	}
}

func NewStats(key string) *Statistics {
	stats := &Statistics{Hash: key}
	exist, err := db.GetDb().HKeys("stats.bet." + stats.Hash)
	if len(exist) > 0 && err == nil {
		db.StructFromKey("stats.bet."+stats.Hash, stats)
	}
	return stats
}

func (s Statistics) save() {
	db.SaveStructure("stats.bet."+s.Hash, s)
}
