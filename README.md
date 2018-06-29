## Koinkoin.io

Koinkoin.io is a simple game that allows you to bet a certain amount of koins with a certain percentage of success. If it's win you win a commission of your bet, otherwise you lose it.

## API

#### GET {host}/api/user/new <br>
__Description__: Will create a user with a random hash and 100 coins. <br>
__Rate limiter__: 1 request per hour <br>
__Response__: <br>

| Field name    | Type         |
| ------------- |:------------:|
| `hash`       | string        |
| `coins`      | uint64        |
| `email`      | string        |
| `name`       | string        |

#### POST {host}/api/bet <br>
__Form__: <br>

| Name          | Type          | Description     |
| ------------- |:-------------:|:----------------|
| `coins`       | uint64        | coins to bet    |
| `chance`      | int           | chance to win   | 
__Description__: bet a amount with a percentage <br>
__Rate limiter__: 1 request per second <br>
__Response__:

| Name          | Type          | Description                 |
| ------------- |:-------------:|:----------------------------|
| `earn`        | uint64        | what user win or lose       |
| `win`         | uint64        | if player win or lose       | 
| `result`      | int           | the number generated        | 
| `coins`       | uint64        | the actual number of coins  | 
