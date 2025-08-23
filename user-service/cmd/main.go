package main

import (
	"fmt"
	"log"

	"messenger-max/user-service/pkg/config"
)

func main() {
	cfg := config.LoadConfig()

	dsn := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable",
		cfg.DBUser, cfg.DBPassword, cfg.DBHost, cfg.DBPort, cfg.DBName)
	log.Println(dsn)
}
