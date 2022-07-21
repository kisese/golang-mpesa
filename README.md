# Setup
😀

Demonstrating MPESA STK Push server side setup in Golang. Gin HTTP Framework.
I went a bit overboard with the logging, I know.

## Installation
Requirements.
- Go 1.13 and above
- Rabitmq

Install the needed dependencies.

```bash
go mod tidy
```

### Local Setup

Start your local rabbitmq server or run rabbitmq via docker
```bash
docker run -d --hostname my-rabbit --name some-rabbit -p 15672:15672 -p 5672:5672 rabbitmq:3-management
```

Start the main API and the queue consumer
```bash
cd cmd
go run main.go server
go run main.go worker
```
Once this is done. You can access the API as below.

For an STK Request

```bash
http://localhost:8080/stk/request

POST
{
    "msisdn":"254700000000",
    "amount":"100"
}
```

If everything is ok you can check your phone for the M-PESA payment pop-up.

Mpesa callbacks will be received on
```bash
http://localhost:8080/stk/callback
```

For env setup, you can find an example env file in the pkg folder,  create your copies and fill as necessary
```bash
cp .env.example .env
```
There is an additional helper for the env file, on utils/getpath where you can specify your working directory name
### Logging
To view debug logs navigate to the following directory.
```bash
/pkg/storage/log/debug.log
```

### Docker

```bash

```

## Contributing
Pull requests are welcome. For major changes, please open an issue first to discuss what you would like to change.

Please make sure to update tests as appropriate.

## License
[MIT](https://choosealicense.com/licenses/mit/)