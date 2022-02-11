package main

import (
	"test-product/app/helper"
	router "test-product/app/server"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load("params/.env")

	helper.CekConnectionDB(err)

	router.Execute()
}
