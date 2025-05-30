# 🐜 Alice Blue Trading API SDK in Go (Unofficial) | Ant Golang Library

Welcome to the **Unofficial Go SDK for Alice Blue (Ant)**, a Go-based library created for seamless integration with Alice Blue's trading API. This project is tailored to address personal use cases and simplifies working with the **Alice Blue API** for trading-related functionalities in Go (Golang).  

**Note**: This is not an official SDK, nor is it endorsed by Alice Blue. Please refer to the [official API documentation](https://v2api.aliceblueonline.com/introduction) for complete API details.

---

## 🚨 Disclaimer  

This SDK is a community-driven, unofficial project developed for personal needs. It may lack some of the features provided by the official API. **Use this library at your own risk.**

---

## ✨ Features of Alice Blue Go SDK

- Simplified authentication and token management.  
- Easy-to-use HTTP client for API communication.  
- Reusable request functions for sending GET, POST, and other HTTP requests.  
- Customizable base URL and headers.  

---

## 📦 Installation  

To install the SDK, use the `go get` command:  

```bash
go get github.com/sureshsolanki17/ant-golang
```

---

## 🔧 How to Use Alice Blue API in Golang 

Here's a quick example to get you started:  

```go
package main

import (
	"fmt"
	"log"
	"github.com/sureshsolanki17/ant-golang/service"
)

func main() {
	// Initialize the configuration
	config := service.Config{
		BaseURL:  "https://ant.aliceblueonline.com/rest/AliceBlueAPIService/api/",
		Exchange: "NSE",
		APIKey:   "yourAPIKey",
		UserId:   "yourUserId",
	}

	// Validate the configuration
	if err := config.Validate(); err != nil {
		log.Fatalf("Error validating config: %v", err)
	}

	// Create new AliceBlue client
	app, err := service.NewConfig(config)
	if err != nil {
		log.Fatalf("Error creating client: %v", err)
	}

	// Connect to the service
	err = app.Connect()
	if err != nil {
		log.Fatalf("Error connecting: %v", err)
	}

	// Make API calls (e.g., get holdings)
	response, err := app.Holdings()
	if err != nil {
		log.Fatalf("Error fetching holdings: %v", err)
	}

	fmt.Println(response)
}
```

---

## 📖 Alice Blue API Documentation for Go

### Initialization  
Create a new configuration and initialize the client:  

```go
config := service.Config{
    BaseURL:  "https://ant.aliceblueonline.com/rest/AliceBlueAPIService/api/",
    Exchange: "NSE",
    APIKey:   "yourAPIKey",
    UserId:   "yourUserId",
}

app, err := service.NewConfig(config)
```

### Authentication  
Connect to the service using your configured credentials:  

```go
err = app.Connect()
```

### Making API Requests  
Various methods are available for different API endpoints:  

```go
// Get holdings
response, err := app.Holdings()
```

---

## 🛠️ Development  

To contribute or explore the source code:  

1. Clone the repository:
   ```bash
   git clone https://github.com/sureshsolanki17/ant-golang.git
   ```
2. Explore the code and submit PRs for improvements.

---

## 🤝 Contributing  

This project welcomes contributions!  
Feel free to submit issues, feature requests, or pull requests to improve this SDK.  

---

## Developers  

- Go Alice Blue SDK  
- Alice Blue Trading API in Golang  
- Ant API Integration using Go  
- Alice Blue Unofficial Go SDK  
- Trading API Golang library  
- Golang SDK for Alice Blue integration  

---

## 🌐 Resources  

- [Official Alice Blue API Documentation](https://v2api.aliceblueonline.com/introduction)  
- [Alice Blue Website](https://www.aliceblueonline.com/)  