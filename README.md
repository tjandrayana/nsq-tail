
## nsq-tail
How to get message from NSQ ? 
You can use this service now.

## How To Use ???
```
1. Clone this repository
2. Run the service 
3. Do request use postman / curl
```

## Example 
```
- Request

Method : POST
URL : http://localhost:8005/get-nsq-messages
Body : 
{
	"lookup_url" : "10.2.0.3",
	"lookup_port":"4161",
	"topic":"User_Info",
	"total_messages":1
}
-----------------------------------------------------

- Response

{
	"user_id":12345,
	"age":20,
	"email":"\"try@gmail.com\""
}

```
