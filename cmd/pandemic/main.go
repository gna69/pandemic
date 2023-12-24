package main

import (
	"fmt"
	"log"

	"github.com/caarlos0/env/v9"
	"github.com/gna69/pandemic/internal/pandemic"
)

func main() {
	var cfg pandemic.Config
	if err := env.Parse(&cfg); err != nil {
		log.Fatal(fmt.Errorf("failed read envs: %w", err))
	}
}
