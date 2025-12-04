# agentic

Exploring agentic wallets and payments.

### a2a

Example code using the A2A protocol.

### erc-8004

Example code using the ERC-8004 trustless agents protocol.

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

Copy variables...

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

### x402

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

{"x402Version":1,"error":"Payment required for this resource","accepts":[{"scheme":"exact","network":"base-sepolia","maxAmountRequired":"1000","asset":"0x036CbD53842c5426634e7929541eC2318f3dCF7e","payTo":"0x49df1720937109e511c0ccd363a82620a596fe18","resource":"http://localhost:8080/protected","description":"Payment required for /protected","mimeType":"","maxTimeoutSeconds":120,"extra":{"name":"USDC","version":"2"}}]}

Protected content
```
