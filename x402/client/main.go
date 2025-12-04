package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"

	"github.com/mark3labs/x402-go/http"
	x402http "github.com/mark3labs/x402-go/http"
	"github.com/mark3labs/x402-go/signers/evm"
)

func main() {

	// Load environment variables
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Failed to load .env file: %v", err)
	}

	// Convert asset decimal to integer
	assetDecimal, err := strconv.Atoi(os.Getenv("ASSET_DECIMAL"))
	if err != nil {
		log.Fatalf("Failed to parse ASSET_DECIMAL: %v", err)
	}

	// Create EVM signer
	signer, _ := evm.NewSigner(
		evm.WithPrivateKey(os.Getenv("SENDER_PRIVATE_KEY")),
		evm.WithNetwork(os.Getenv("NETWORK_NAME")),
		evm.WithToken(os.Getenv("ASSET_ADDRESS"), os.Getenv("ASSET_NAME"), assetDecimal),
	)

	// Create http client (without x402)
	client1, _ := http.NewClient()

	// Create http client (with x402 enabled)
	client2, _ := x402http.NewClient(x402http.WithSigner(signer))

	// Make request (public, returns 200)
	resp1, _ := client1.Get("http://localhost:8080/public")
	defer resp1.Body.Close()

	body1, _ := io.ReadAll(resp1.Body)
	fmt.Println(string(body1))

	// Make request (protected, returns 402)
	resp2, _ := client1.Get("http://localhost:8080/protected")
	defer resp2.Body.Close()

	body2, _ := io.ReadAll(resp2.Body)
	fmt.Println(string(body2))

	// Make request (protected, returns 200)
	resp3, _ := client2.Get("http://localhost:8080/protected")
	defer resp3.Body.Close()

	body3, _ := io.ReadAll(resp3.Body)
	fmt.Println(string(body3))
}
