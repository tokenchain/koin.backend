package err

import "errors"

var (
	ChanceCantBeEqOrHigher98 = errors.New("chance parameter cannot be higher than 98")
	ChanceCantBeLesser2      = errors.New("chance parameter cannot be lesser than 2")
	CoinsCantBeLesser5       = errors.New("bet cant be lesser than 5")
	NotAuthenticated         = errors.New("user need to be authenticated")
	AlreadyAuthenticated     = errors.New("user need to not be authenticated")
	NumberMalformed          = errors.New("number are malformed")
	NotEnoughCoins           = errors.New("not enough koins")
	IncorrectParameter       = errors.New("incorrect parameter")
	IncorrectName            = errors.New("name is incorrect")
	IncorrectMail            = errors.New("mail is incorrect")
)
