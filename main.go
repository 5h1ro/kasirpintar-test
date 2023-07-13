package main

import (
	"fmt"
	"log"
	"os"

	"kaspin/docs"
	"kaspin/server"

	"github.com/joho/godotenv"
	"github.com/rs/zerolog"
)

// @title Kasir Pintar x NicePay
// @version 1.0
// @description This is a test task to join Kasir Pintar.

// @contact.name Nurhakiki Romadhony Ikhwandany
// @contact.url https://http://hqq.seovdetech.com//
// @contact.email nurhakiki.ri@gmail.com

// @BasePath /
func main() {
	reqFile, err := os.OpenFile(
		"log/request.log",
		os.O_APPEND|os.O_CREATE|os.O_WRONLY,
		0664,
	)
	defer reqFile.Close()
	if err != nil {
		panic(err)
	}
	resFile, err := os.OpenFile(
		"log/response.log",
		os.O_APPEND|os.O_CREATE|os.O_WRONLY,
		0664,
	)
	defer resFile.Close()
	if err != nil {
		panic(err)
	}
	errFile, err := os.OpenFile(
		"log/error.log",
		os.O_APPEND|os.O_CREATE|os.O_WRONLY,
		0664,
	)
	defer errFile.Close()
	if err != nil {
		panic(err)
	}
	reqLog := zerolog.New(reqFile).With().Timestamp().Logger()
	resLog := zerolog.New(resFile).With().Timestamp().Logger()
	errLog := zerolog.New(errFile).With().Timestamp().Logger()
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	if os.Getenv("MODE") == "production" {
		docs.SwaggerInfo.Host = fmt.Sprintf("%s:%s", os.Getenv("HOST"), os.Getenv("EXPOSE_PORT"))
	} else {
		docs.SwaggerInfo.Host = fmt.Sprintf("%s:%s", os.Getenv("HOST"), os.Getenv("PORT"))
	}

	app := server.NewServer()
	server.ConfigureRoutes(app, reqLog, resLog, errLog)

	if err := app.Run(os.Getenv("PORT")); err != nil {
		log.Fatal(err)
	}
}
