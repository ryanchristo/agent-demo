package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"os"

	"github.com/ipfs/boxo/files"
	"github.com/ipfs/kubo/client/rpc"
)

func main() {
	ctx := context.Background()

	// Connect to IPFS deamon.
	api, err := rpc.NewLocalApi()
	if err != nil {
		log.Fatal(err)
	}

	// Open agent registration file.
	reader1, err := os.Open("examples/registration.json")
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	defer reader1.Close()

	// Read agent registration file.
	file1 := files.NewReaderFile(reader1)

	// Add agent registration file to IPFS.
	cid1, err := api.Unixfs().Add(ctx, file1)
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	// Print result.
	fmt.Println(cid1)

	// Get agent registration data from IPFS.
	node1, err := api.Unixfs().Get(ctx, cid1)
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	// Convert node to file.
	retrieved1 := files.ToFile(node1)

	// Read agent registration data.
	bytes1, err := io.ReadAll(retrieved1)
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	// Print result.
	fmt.Println(string(bytes1))

	// Open agent feedback file.
	reader2, err := os.Open("examples/feedback.json")
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	defer reader2.Close()

	// Read agent feedback file.
	file2 := files.NewReaderFile(reader2)

	// Add agent feedback file to IPFS.
	cid2, err := api.Unixfs().Add(ctx, file2)
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	// Print result.
	fmt.Println(cid2)

	// Get agent feeback data from IPFS.
	node2, err := api.Unixfs().Get(ctx, cid2)
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	// Convert node to file.
	retrieved2 := files.ToFile(node2)

	// Read agent feedback data.
	bytes2, err := io.ReadAll(retrieved2)
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	// Print result.
	fmt.Println(string(bytes2))
}
