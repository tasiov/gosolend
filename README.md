# gosolend

Go client for interacting with the Solend protocol on Solana.

## Project Structure

```
├── cmd/                  # Command-line tools and examples
├── internal/            # Private application and library code
├── pkg/                 # Public API code
│   ├── client/         # Main client interface
│   ├── models/         # Data models
│   └── service/        # Service implementations
├── scripts/            # Build and maintenance scripts
└── README.md
```

## Installation

```bash
go get github.com/tasiov/gosolend
```

## Usage

```go
package main

import (
    "context"
    "log"

    "github.com/tasiov/gosolend/pkg/client"
)

func main() {
    c := client.NewClient("https://api.mainnet-beta.solana.com")

    reserve, err := c.GetReserve(context.Background(), "reserveAddress")
    if err != nil {
        log.Fatal(err)
    }

    log.Printf("Borrow rate: %f", reserve.BorrowRate)
}
```

## Development

1. Clone the repository
2. Run tests: `go test ./...`
3. Generate models: `go run scripts/generate_models.go`

## License

MIT
