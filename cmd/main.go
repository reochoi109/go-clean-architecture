package main

import (
	"fmt"
	"go-clean-architecture/config"
)

func main() {
	cfg := config.Load()
	fmt.Printf("Loaded Config: %s\n", cfg)
}
