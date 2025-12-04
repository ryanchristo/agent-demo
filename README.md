# agent-demo

Exploring agent protocols and frameworks.

### a2a

Example code using the A2A protocol.

### erc-8004

Example code using the ERC-8004 protocol.

### ipfs

Example code using the IPFS protocol.

### mcp

Example code using the model context protocol.

### x402

Example code using the x402 protocol.

## Get Started

### a2a

Run example server...

```sh
go run server/main.go
```

Run example client...

```sh
go run client/main.go
```

View response...

```
2025/12/02 12:28:13 Server responded with: &{ID:019ae0c0-5907-785e-9011-fb5050a6cfc8 ContextID: Extensions:[] Metadata:map[] Parts:[{Text:Hello, world! Metadata:map[]}] ReferenceTasks:[] Role:agent TaskID:}
```

### erc-8004

Set environment...

```sh
cp .env-example .env
```

Run example client...

```sh
go run client/main.go
```

View response...

```
Connecting to node RPC endpoint...
Creating ERC-8004 contract clients...
Calling ERC-8004 contract methods...
Version: [1.0.0]
Version: [1.0.0]
Version: [1.0.0]
```

### ipfs

Start ipfs daemon...

```sh
ipfs daemon
```

Run example client...

```sh
go run client/main.go
```

View response...

```
/ipfs/QmfFZkpJGawhzVU6GupCUmxvV41JuSF5AUzQvRGqzfZoWv
{
  "type": "https://eips.ethereum.org/EIPS/eip-8004#registration-v1",
  "name": "myAgentName",
  "description": "A natural language description of the Agent",
  "image": "https://example.com/agent-image.png",
  "endpoints": [
    {
      "name": "agentWallet",
      "endpoint": "eip155:1:0x742d35Cc6634C0532925a3b844Bc9e7595f0bEb7"
    },
    {
      "name": "DID",
      "endpoint": "did:method:foobar",
      "version": "v1"
    },
    {
      "name": "ENS",
      "endpoint": "vitalik.eth",
      "version": "v1"
    },
    {
      "name": "A2A",
      "endpoint": "https://api.example.com/a2a",
      "version": "0.30",
      "a2aSkills": [
        "skill_1",
        "skill_2",
        "skill_3"
      ]
    },
    {
      "name": "MCP",
      "endpoint": "https://api.example.com/mcp",
      "version": "2025-06-18",
      "mcpTools": [
        "data_analysis",
        "chart_generation",
        "report_creation"
      ],
      "mcpPrompts": [
        "prompt_1",
        "prompt_2",
        "prompt_3"
      ],
      "mcpResources": [
        "resource_1",
        "resource_2",
        "resource_3"
      ]
    },
    {
      "name": "OASF",
      "endpoint": "https://github.com/agntcy/oasf/",
      "version": "v0.8.0",
      "skills": [
        "advanced_reasoning_planning/strategic_planning",
        "data_engineering/data_transformation_pipeline"
      ],
      "domains": [
        "finance_and_business/investment_services",
        "technology/data_science/data_visualization"
      ]
    }
  ],
  "registrations": [
    {
      "agentId": 241,
      "agentRegistry": "eip155:11155111:0x8004a6090Cd10A7288092483047B097295Fb8847"
    }
  ],
  "supportedTrusts": [
    "reputation",
    "crypto-economic",
    "tee-attestation"
  ],
  "active": true,
  "x402support": true
}
/ipfs/Qma4TVqn8pAf7xSeT9sVGdGTfZTCc2LkZzVnfUguJTDpkH
{
  "agentRegistry": "eip155:1:{identityRegistry}",
  "agentId": 22,
  "clientAddress": "eip155:1:{clientAddress}",
  "createdAt": "2025-09-23T12:00:00Z",
  "feedbackAuth": "...",
  "score": 70,
  "tag1": "foo",
  "tag2": "bar",
  "skill": "as-defined-by-A2A",
  "context": "as-defined-by-A2A",
  "task": "as-defined-by-A2A",
  "capability": "tools",
  "name": "Put the name of the MCP tool you liked!",
  "proofOfPayment": {
    "fromAddress": "0x00...",
    "toAddress": "0x00...",
    "chainId": "1",
    "txHash": "0x00..."
  }
}
```

### mcp

Run example server...

```sh
go run server/main.go
```

Run example client...

```sh
go run client/main.go
```

View response...

```
2025/12/03 21:32:59 Connecting to MCP server at http://localhost:8000
2025/12/03 21:32:59 Connected to server (session ID: Y3QVIILO37ZLOOUKCG2P5C57YZ)
2025/12/03 21:32:59 Listing available tools...
2025/12/03 21:32:59   - cityTime: Get the current time in NYC, San Francisco, or Boston
2025/12/03 21:32:59 Getting time for each city...
2025/12/03 21:32:59   The current time in New York City is 2025-12-04T00:32:59-05:00
2025/12/03 21:32:59   The current time in San Francisco is 2025-12-03T21:32:59-08:00
2025/12/03 21:32:59   The current time in Boston is 2025-12-04T00:32:59-05:00
2025/12/03 21:32:59 Client completed successfully
```

### x402

Set environment...

```sh
cp .env-example .env
```

Run example server...

```sh
go run server/main.go
```

Run example client...

```sh
go run client/main.go
```

View response...

```
Public content

{"x402Version":1,"error":"Payment required for this resource","accepts":[{"scheme":"exact","network":"base-sepolia","maxAmountRequired":"1000","asset":"0x036CbD53842c5426634e7929541eC2318f3dCF7e","payTo":"0x0","resource":"http://localhost:8080/protected","description":"Payment required for /protected","mimeType":"","maxTimeoutSeconds":120,"extra":{"name":"USDC","version":"2"}}]}

Protected content
```
