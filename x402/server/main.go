package main

import (
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/joho/godotenv"

	"github.com/mark3labs/x402-go"
	x402http "github.com/mark3labs/x402-go/http"
)

func main() {

	// Load environment variables
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Failed to load .env file: %v", err)
	}

	// Convert asset decimal to integer
	maxTimeout, err := strconv.Atoi(os.Getenv("MAX_TIMEOUT"))
	if err != nil {
		log.Fatalf("Failed to parse MAX_TIMEOUT: %v", err)
	}

	// Configure x402 middleware
	config := &x402http.Config{
		FacilitatorURL: os.Getenv("FACILITATOR_URL"),
		PaymentRequirements: []x402.PaymentRequirement{{
			Scheme:            "exact",
			Network:           os.Getenv("NETWORK_NAME"),
			MaxAmountRequired: os.Getenv("MAX_AMOUNT"),
			Asset:             os.Getenv("ASSET_ADDRESS"),
			PayTo:             os.Getenv("RECIPIENT_ADDRESS"),
			MaxTimeoutSeconds: maxTimeout,
			Extra: map[string]interface{}{
				"name":    os.Getenv("ASSET_NAME"),
				"version": os.Getenv("ASSET_VERSION"),
			},
		}},
	}

	// Create x402 middleware
	middleware := x402http.NewX402Middleware(config)

	// Create public handler
	publicHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Public content"))
	})

	// Create protected handler
	protectedHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Protected content"))
	})

	// Handle public endpoint
	http.Handle("/public", publicHandler)

	// Handle protected endpoint
	http.Handle("/protected", middleware(protectedHandler))

	// Listen and server
	http.ListenAndServe(":8080", nil)
}
