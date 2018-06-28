## Koinkoin.io

Koinkoin.io is a simple game that allows you to bet a certain amount of koins with a certain percentage of success. If it's win you win a commission of your bet, otherwise you lose it.

## API

- GET {host}/api/user/new <br>
__description__: Will create a user with a random hash and 100 coins. <br>
__rate limiter__: 1 request per hour <br>
__response__:
```
{
  "Hash": string,
  "Coins": uint64,
  "Email": string,
  "Name": string
}
```

- POST {host}/api/bet <br>
__form__: coins uint64, chance int <br>
__description__: bet a amount with a percentage <br>
__rate limiter__: 1 request per second <br>
__response__:
```
{
  "earn": uint64, //what user win or lose
  "win": uint64, //if player win or lose
  "result": int, //the number generated
  "coins": uint64 //the actual number of coins
}
```