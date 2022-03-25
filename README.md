# Setup
😀

Demonstrating MPESA STK Push server side setup in Golang. Gin HTTP Framework.
I went a bit overboard with the logging, I know.

## Installation

Use the package manager [go get](https://pkg.go.dev/) to install dependancies for the two separate apps (mpesa-request, mpesa-consumer).

```bash
go get
```

## Usage
You can use a local RabbitMQ setup or use docker compose.

### Docker Compose
We have a docker compose setup that will handle the initial setup.

```bash
 docker-compose up --build
```

Once this is done. You can access the apps as below.

mpesa-request app
```bash
http://localhost:8000/stk-request
```
mpesa-listener app
```bash
http://localhost:8001/stk-callback
```

For env setup, you can find example env files for each of the respective foder,  create your copies and fill as necessary
```bash
cp .env.example .env
```

For startup without docker-compose.

Start your local rabbitmq server

Start the Mpesa Consumer App
```bash
cd mpesa-consumer
go run main.go 
```

Start the Mpesa Request App
```bash
cd mpesa-request
go run main.go 
```

## Contributing
Pull requests are welcome. For major changes, please open an issue first to discuss what you would like to change.

Please make sure to update tests as appropriate.

## License
[MIT](https://choosealicense.com/licenses/mit/)