
## nsq-tail
How to get message from NSQ ? 
You can use this service now.

## Why I create this ?
When I learn nsq for the first time, I found problem how to get the message. My friend just tell me to get nsq tail.
I don't know how to run nsq tail for the first time. 
Based on my problem, hopefully this service can help you to get the nsq message.
Please just tell your friend, "Hei you, just  clone https://github.com/tjandrayana/nsq-tail repository and then run the service."


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
