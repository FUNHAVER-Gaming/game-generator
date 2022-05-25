package main

import (
	"fmt"
	"valorant-league/pkg/service"
)

func main() {
	fmt.Println("Setting up...")
	service.Setup()
	fmt.Println("Done.")
}
