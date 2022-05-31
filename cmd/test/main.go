package main

import (
	"fmt"
	"github.com/otiai10/gosseract/v2"
)

func main() {
	pathToFile := "C:\\Users\\tda\\Pictures\\valorant.png"
	client := gosseract.NewClient()
	defer client.Close()
	client.SetImage(pathToFile)
	text, _ := client.Text()
	fmt.Println(text)
}
