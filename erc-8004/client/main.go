package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"

	"github.com/ethereum/go-ethereum/ethclient"

	"github.com/ryanchristo/agent-demo/erc-8004/contracts"
	"github.com/ryanchristo/agent-demo/erc-8004/contracts/identity"
	"github.com/ryanchristo/agent-demo/erc-8004/contracts/reputation"
	"github.com/ryanchristo/agent-demo/erc-8004/contracts/validation"
)

func main() {

	// Load environment variables.
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Failed to load .env file: %v", err)
	}

	// Get RPC environment variable.
	ethRpcUrl := os.Getenv("ETH_RPC_URL")
	if ethRpcUrl == "" {
		log.Fatal("ETH_RPC_URL environment variable not set")
	}

	// Log connecting to RPC endpoint.
	fmt.Println("Connecting to node RPC endpoint...")

	// Connect ethereum client to RPC endpoint.
	client, err := ethclient.Dial(ethRpcUrl)
	if err != nil {
		log.Fatalf("Failed to connect to RPC endpoint: %v", err)
	}

	// Close ethereum client when program exits.
	defer client.Close()

	// Log creating contract clients.
	fmt.Println("Creating ERC-8004 contract clients...")

	// Create identity contract client.
	identityClient, err := identity.NewClient(context.Background(), client)
	if err != nil {
		log.Fatalf("Failed to create identity client: %v", err)
	}

	// Create reputation contract client.
	reputationClient, err := reputation.NewClient(context.Background(), client)
	if err != nil {
		log.Fatalf("Failed to create reputation client: %v", err)
	}

	// Create validation contract client.
	validationClient, err := validation.NewClient(context.Background(), client)
	if err != nil {
		log.Fatalf("Failed to create validation client: %v", err)
	}

	// Log calling contract methods.
	fmt.Println("Calling ERC-8004 contract methods...")

	// Get identity contract version.
	identityVersion, err := contracts.Read(identityClient.Contract, "getVersion")
	if err != nil {
		log.Fatalf("Failed to call contract method getVersion: %v", err)
	} else {
		fmt.Printf("Version: %s\n", identityVersion)
	}

	// Get reputation contract version.
	reputationVersion, err := contracts.Read(reputationClient.Contract, "getVersion")
	if err != nil {
		log.Fatalf("Failed to call contract method getVersion: %v", err)
	} else {
		fmt.Printf("Version: %s\n", reputationVersion)
	}

	// Get validation contract version.
	validationVersion, err := contracts.Read(validationClient.Contract, "getVersion")
	if err != nil {
		log.Fatalf("Failed to call contract method getVersion: %v", err)
	} else {
		fmt.Printf("Version: %s\n", validationVersion)
	}
}
