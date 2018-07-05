package err

import "github.com/kataras/iris"

type Err struct {
	Code  int    `json:"code"`
	Error string `json:"error"`
	Name  string `json:"name"`
}

var (
	ChanceCantBeEqOrHigher98 = &Err{
		0,
		"you cannot set a chance higher or equal to 98",
		"ChanceCantBeEqOrHigher98",
	}
	ChanceCantBeLesser2 = &Err{
		1,
		"you cannot set a chance lesser or equal to 2",
		"ChanceCantBeLesser2",
	}
	CoinsCantBeLesser5 = &Err{
		2,
		"your bet cant be lesser than 5",
		"CoinsCantBeLesser5",
	}
	NotAuthenticated = &Err{
		3,
		"you need to be authenticated",
		"NotAuthenticated",
	}
	AlreadyAuthenticated = &Err{
		4,
		"you need to not be authenticated to do this action",
		"AlreadyAuthenticated",
	}
	NoUserFound = &Err{
		5,
		"user not found",
		"NoUserFound",
	}
	NumberMalformed = &Err{
		6,
		"number are malformed",
		"NumberMalformed",
	}
	NotEnoughCoins = &Err{
		7,
		"not enough koins",
		"NotEnoughCoins",
	}
	IncorrectParameter = &Err{
		8,
		"incorrect parameter",
		"IncorrectParameter",
	}
	IncorrectName = &Err{
		9,
		"name is incorrect",
		"IncorrectName",
	}
	IncorrectMail = &Err{
		10,
		"mail is incorrect",
		"IncorrectMail",
	}
)

// ThrownError take a context and an error and return a json with
// the error.
func ThrownError(ctx iris.Context, err *Err) {
	ctx.StatusCode(iris.StatusBadRequest)
	ctx.JSON(err)
}
