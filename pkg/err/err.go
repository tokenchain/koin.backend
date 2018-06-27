package err

import "errors"

var (
	ChanceCantBeEqOrHigher98 = errors.New("chance parameter cannot be higher than 98")
	ChanceCantBeLesser2      = errors.New("chance parameter cannot be lesser than 2")
	CoinsCantBeLesser5       = errors.New("bet cant be lesser than 5")
)
