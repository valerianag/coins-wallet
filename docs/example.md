Request
```
➜  ~ curl --request GET 'http://localhost:8080/accounts' -i
```
Response
```
HTTP/1.1 200 OK
Content-Type: application/json
Date: Mon, 08 Mar 2021 11:11:55 GMT
Content-Length: 18

{"accounts":null}
```

Request
```
➜  ~ curl --request GET 'http://localhost:8080/payments' -i
```    
Response                                                                                        
```
TTP/1.1 200 OK
Content-Type: application/json
Date: Mon, 08 Mar 2021 11:12:00 GMT
Content-Length: 18

{"payments":null}
```

Request
```
➜  ~ curl --request POST 'http://localhost:8080/accounts' --data-raw '{"account": {"name": "bob123", "balance": "1236.45", "currency": "USD"}}' -i
```
Response
```
TTP/1.1 200 OK
Content-Type: application/json
Date: Mon, 08 Mar 2021 11:18:21 GMT
Content-Length: 5

null
```

Request
```
➜  ~ curl --request POST 'http://localhost:8080/accounts' --data-raw '{"account": {"name": "max", "balance": "1236.45", "currency": "USD"}}' -i
```
Response
```
TTP/1.1 200 OK
Content-Type: application/json
Date: Mon, 08 Mar 2021 11:18:32 GMT
Content-Length: 5

null
```

Request
```
➜  ~ curl --request POST 'http://localhost:8080/accounts' --data-raw '{"account": {"name": "llew", "balance": "123.00", "currency": "USD"}}' -i
```
Response
```
TTP/1.1 200 OK
Content-Type: application/json
Date: Mon, 08 Mar 2021 11:18:48 GMT
Content-Length: 5

null
```

Request
```
➜  ~ curl --request POST 'http://localhost:8080/accounts' --data-raw '{"account": {"name": "wqrt", "balance": "1250.00", "currency": "EUR"}}' -i
```
Response
```
TTP/1.1 200 OK
Content-Type: application/json
Date: Mon, 08 Mar 2021 11:19:11 GMT
Content-Length: 5

null
```

Request
```
➜  ~ curl --request GET 'http://localhost:8080/accounts' -i
```
Response                                                                                       
```
TTP/1.1 200 OK
Content-Type: application/json
Date: Mon, 08 Mar 2021 11:19:18 GMT
Content-Length: 221

{"accounts":[{"name":"bob123","balance":"1236.45","currency":"USD"},{"name":"max","balance":"1236.45","currency":"USD"},{"name":"llew","balance":"123","currency":"USD"},{"name":"wqrt","balance":"1250","currency":"EUR"}]}
```

Request
```
➜  ~ curl --request POST 'http://localhost:8080/payments' --data-raw '{"payment": {"from_account": "bob123", "to_account": "max", "amount": "1236.00"}}' -i
```
Response
```
HTTP/1.1 200 OK
Content-Type: application/json
Date: Mon, 08 Mar 2021 11:20:03 GMT
Content-Length: 5

null
```

Request
```
➜  ~ curl --request GET 'http://localhost:8080/payments' -i
```
Response                                                                                             
```
TTP/1.1 200 OK
Content-Type: application/json
Date: Mon, 08 Mar 2021 11:20:10 GMT
Content-Length: 76

{"payments":[{"from_account":"bob123","to_account":"max","amount":"1236"}]}
```

Request
```
➜  ~ curl --request GET 'http://localhost:8080/accounts' -i
```                                                                                                
Response
```
TTP/1.1 200 OK
Content-Type: application/json
Date: Mon, 08 Mar 2021 11:20:15 GMT
Content-Length: 218

{"accounts":[{"name":"llew","balance":"123","currency":"USD"},{"name":"wqrt","balance":"1250","currency":"EUR"},{"name":"bob123","balance":"0.45","currency":"USD"},{"name":"max","balance":"2472.45","currency":"USD"}]}
```

Request
```
➜  ~ curl --request POST 'http://localhost:8080/payments' --data-raw '{"payment": {"from_account": "max", "to_account": "wqrt", "amount": "1236.00"}}' -i
```
Response
```
HTTP/1.1 400 Bad Request
Content-Type: text/plain; charset=utf-8
Date: Mon, 08 Mar 2021 11:21:09 GMT
Content-Length: 20

not equal currencies%                                                                                                                                                                                    
```

Request
```
➜  ~ curl --request POST 'http://localhost:8080/payments' --data-raw '{"payment": {"from_account": "max", "to_account": "llew", "amount": "10000.00"}}' -i
```
Response
```
HTTP/1.1 400 Bad Request
Content-Type: text/plain; charset=utf-8
Date: Mon, 08 Mar 2021 11:21:49 GMT
Content-Length: 18

not enough balance%
```       

Request                                                                                                                                                                         
```
➜  ~ curl --request POST 'http://localhost:8080/payments' --data-raw '{"payment": {"from_account": "max", "to_account": "llew", "amount": "10.003"}}' -i
``` 
Response
```
HTTP/1.1 400 Bad Request
Content-Type: text/plain; charset=utf-8
Date: Mon, 08 Mar 2021 11:21:57 GMT
Content-Length: 26

should be 2 decimal places%       
```  

Request                                                                                                                                                                     
```
➜  ~ curl --request POST 'http://localhost:8080/payments' --data-raw '{"payment": {"from_account": "max", "to_account": "llew", "amount": "103}' -i
```     
Response
```
HTTP/1.1 400 Bad Request
Content-Type: text/plain; charset=utf-8
Date: Mon, 08 Mar 2021 11:22:06 GMT
Content-Length: 14

unexpected EOF%    
```