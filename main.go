package main

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	loadEnv()
}

func loadEnv() {
	err := godotenv.Load(".env")

	if err != nil {
		fmt.Printf("読み込みできませんでした: %v", err)
	}

	message := os.Getenv("SAMPLE_MESSAGE")
	fmt.Println(message)
}