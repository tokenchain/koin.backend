## Koin.bet

Koin.bet is a simple game that allows you to bet a certain amount of koins with a certain percentage of success. If it's win you win a commission of your bet, otherwise you lose it.

## File: datas.env 

Need to contains this fields:
```
redis_host=redis
redis_port=6379

mail=koin@gmail.com
mail_pwd=password_here
mail_host=smtp.gmail.com
mail_port=465

```


## API

#### Note on error

On each endpoint check if an error field is present, if it's the case the error message would be like this:

| Field name    | Type         |
| ------------- |:------------:|
| `code`       | int        |
| `error`      | string        |
| `name`      | string        |

The `code` is the identifier of the error. 
The `error` field is the message of the error.

### User

#### GET {{host}}/api/user/new
__Description__: Will create a user with a random hash and 100 coins. <br>
__Rate limiter__: 1 request per hour <br>
__Response__: <br>

| Field name    | Type         |
| ------------- |:------------:|
| `hash`       | string        |
| `coins`      | uint64        |
| `email`      | string        |
| `name`       | string        |

#### GET {{host}}/api/bet/mail/?mail={{hash}}

__Description__: Send a mail with the hash of the player. <br>
__Rate limiter__: 1 request per hour <br>
__Response__: <br>

| Name          | Type          |
| ------------- |:-------------:|
| `success`       | bool        |


#### GET {{host}}/api/user/update

__Description__: Update the name or the mail of an user <br>
__Rate limiter__: 1 request per second <br>
__Header__: `hash: string` <br>
__Form__: <br>

| Name          | Type          |
| ------------- |:-------------:|
| `name?`       | string        |
| `mail?`      | string        |

__Note__: `?` mean that is __optional__. <br>
__Response__:

| Name          | Type          |
| ------------- |:-------------:|
| `success`       | bool        |

#### GET {{host}}/api/user/?hash={{hash}}

__Description__: Get information about an user. <br>
__Rate limiter__: 20 request per seconds <br>
__Response__: <br>

| Field name    | Type         |
| ------------- |:------------:|
| `hash`       | string        |
| `coins`      | uint64        |
| `email`      | string        |
| `name`       | string        |

### Bet

#### POST {{host}}/api/bet <br>

__Description__: bet a amount with a percentage <br>
__Header__: `hash: string` <br>
__Rate limiter__: 2 request per second <br>
__Form__: <br>

| Name          | Type          | Description     |
| ------------- |:-------------:|:----------------|
| `coins`       | uint64        | coins to bet.   |
| `chance`      | int           | chance to win.   |

__Response__:

| Name          | Type          | Description                 |
| ------------- |:-------------:|:----------------------------|
| `result`        | int        | is the number generated by the Bet method       |
| `earn`         | uint64        | is the number that the user win or lose.     | 
| `win`          | bool           | is a boolean to say if yes or no the user has win.        | 
| `chance`       | int        | is a boolean to say if yes or no the user has win. | 
| `coins`       | uint64        | is the amount of coins that the player bet. | 
| `afterCoins`       | uint64        | is the amount of new coins after add or remove the gain. | 
| `beforeCoins`       | uint64        | is the amount of old coins before bet.| 

#### POST {{host}}/api/bet/stats/?hash={{hash}}

__Response__:

| Name          | Description                 |
| ------------- |:----------------------------|
| `hash` | is the identifier. |
| `count` | is the total bet effectuated. |
| `averageEarn` | is the average of earn. |
| `averageLose` | is the average of lose. |
| `averageCoins` | is the average of coins bet. |
| `averageChance` | is the average of chance bet. |
| `averageResult` | is the average of result number generated. |
| `maxAmount` | is the max coins amount bet. |
| `minAmount` | is the min coins amount bet. |
| `maxChance` | is the highest chance bet. |
| `minChance` | is the lowest chance bet. |
| `maxEarn` | is the maximum bettor earn. |
| `minEarn` | is the minimum a bettor earn. |
| `maxLose` | is the maximum bettor lose. |
| `minLose` | is the minimum a bettor lose. |
| `success` | is the amount of positive bet. |
| `failed` | is the amount of negative bet. |
| `totalCoins` | is the total chance bet. |
| `totalEarn` | is the total a bettor earn. |
| `totalLose` | is the total a bettor lose. |
| `totalChance` | is the total a bettor chance. |
| `totalResult` | is the total of result obtained. |
| `greedy` | represent the bettor as 'greedy'. |
| `fearful` | represent the bettor as 'fearful'. |

__Note__: all of theses types are __uint64__ <br>
__Note 2__: use `global` as a hash to get global statistics over all bettors.

