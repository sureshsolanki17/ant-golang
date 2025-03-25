package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/sureshsolanki17/ant-golang/service"
)

var (
	ApiKey string
	UserId string
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Printf("Warning: .env file not found")
	}
	ApiKey = os.Getenv("ANT_API_KEY")
	if ApiKey == "" {
		log.Fatal("ANT_API_KEY environment variable is required")
	}

	UserId = os.Getenv("ANT_USER_ID")
	if UserId == "" {
		log.Fatal("ANT_USER_ID environment variable is required")
	}
}

func main() {
	config := service.Config{
		BaseURL:  "https://ant.aliceblueonline.com/rest/AliceBlueAPIService/api/",
		Exchange: "NSE",
		APIKey:   ApiKey,
		UserId:   UserId,
	}

	if err := config.Validate(); err != nil {
		log.Fatalf("Error validating config: %v", err)
	}

	app, err := service.NewConfig(config)
	if err != nil {
		panic(err)
	}

	err = app.Connect()
	if err != nil {
		panic(err)
	}

	response, err := app.GetScripQuoteDetails("11536")
	if err != nil {
		log.Fatalf("Error making request: %v", err)
	}

	file, _ := os.Create("response.json")
	defer file.Close()

	jsonData, err := json.Marshal(response)
	if err != nil {
		log.Fatalf("Error marshaling response: %v", err)
	}

	if _, err := file.Write(jsonData); err != nil {
		log.Fatalf("Error writing file: %v", err)
	}

	fmt.Printf("%+v\n", response)
}
