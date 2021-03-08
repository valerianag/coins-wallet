## Wallet http API description
### Payments

#### Get all payments
GET /payments
##### Request example
```curl --request GET 'http://localhost:8080/payments' -i```
##### Response example
```json
{"payments":[
  {"from_account":"bob123","to_account":"lers143","amount":"23.55"},
  {"from_account":"lers143","to_account":"max","amount":"123.56"}
]}
```

#### Create new payment
POST /payments
##### Json body example
```json
{
  "payment": {
    "from_account": "bob123",
    "to_account": "lers143",
    "amount": 2345.90
  }
} 
```
##### Request example
```curl --request POST 'http://localhost:8080/payments' --data-raw '{"payment": {"from_account": "bob123", "to_account": "lers143", "amount": 2345.90}}' -i```

### Accounts

#### Get all accounts
GET /accounts
##### Request example
```curl --request GET 'http://localhost:8080/accounts' -i```
##### Response example
```json
{"accounts":[
  {"name":"bob123","balance":"978","currency":"USD"},
  {"name":"lers143","balance":"1371.34","currency":"USD"},
  {"name":"max","balance":"247.01","currency":"USD"},
  {"name":"John Wick","balance":"100000","currency":"RUB"}
]}
```

#### Create new account
POST /accounts
##### Json body example
```json
{
  "account": {
    "name": "bob123", 
    "balance": "1236.45", 
    "currency": "USD"
  }
}
```
##### Request example
```curl --request POST 'http://localhost:8080/accounts' --data-raw '{"account": {"name": "bob123", "balance": "1236.45", "currency": "USD"}}' -i```
