# agent-demo

Exploring agent protocols and frameworks.

### a2a

Example code using the A2A protocol.

### erc-8004

Example code using the ERC-8004 trustless agents protocol.

### mcp

Example code using the Model Context Protocol.

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
