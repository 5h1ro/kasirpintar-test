# Overview

It's an API Kasir Pintar with Nicepay payment gateway based on Gin framework. I also added swagger for API documentation and dockerize this code.

## Endpoint:

- Registration (/register)
- Check Status (/status)
- Payment (/payment)
- Swagger documentation (/swagger/index.html)

## Usage

1.  Copy .env.example to .env and set the environment variables:

    `cp .env.example .env`

2.  Run unit test using this command, or you can skip this step:

    `go test ./server/handler/ -v`

3.  Before run the application, please change mode to **production** in .env file:

4.  Run your application using the command in the terminal:

    `docker-compose up -d`

5.  Browse to {HOST}:{EXPOSE_PORT}/swagger/index.html. You will see swagger documentation.
