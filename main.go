package main

import (
	"devconfcm/pkg/router"

	_ "github.com/joho/godotenv/autoload"
)

func main() {
	router.Start()
}
