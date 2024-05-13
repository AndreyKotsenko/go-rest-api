package main

import (
	"fmt"
	"rest-api/internal/config"
)

func main() {
	cfg := config.MustLoad()

	fmt.Println(cfg)
}
