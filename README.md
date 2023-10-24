# Galactic Service
Backend source code for galactic application/service.

## Architecture
This backend application/service is using the microservices approach defined by Go-Kit (https://github.com/go-kit/kit).

## How to Run Locally
- Run MySQL database and change the DB configs in `.env` file.
- Make sure Air (https://github.com/cosmtrek/air) installed.
- Run `make run-server` or `go run cmd/server/main.go`.
- Explore the API(s) and have fun!

## Documentation
The API(s) documentation is generated using Swagger and available at `/swagger/index.html`.