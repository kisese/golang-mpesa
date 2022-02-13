# Setup

Demonstrating MPESA STK Push server side setup in Golang. Gin HTTP Framework.
I went a bit overboard with the logging, I know.

## Installation

Use the package manager [go get](https://pkg.go.dev/) to install foobar.

```bash
go get
```

## Usage
You should have RabbitMQ setup already

Startup the project
```python
cp .env.example .env

go run main.go
```
Start the Queue Listener
```python
go run QueueListener.go 
```

## Contributing
Pull requests are welcome. For major changes, please open an issue first to discuss what you would like to change.

Please make sure to update tests as appropriate.

## License
[MIT](https://choosealicense.com/licenses/mit/)