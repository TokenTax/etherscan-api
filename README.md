# etherscan-api

[![GoDoc](https://godoc.org/github.com/TokenTax/etherscan-api/v2?status.svg)](https://godoc.org/github.com/TokenTax/etherscan-api)

Golang client for the Etherscan.io v2 API with nearly full implementation(accounts, transactions, tokens, contracts, blocks, stats) and minimal dependencies.

# Usage

```bash
go get github.com/TokenTax/etherscan-api/v2
```

Create an API instance and off you go. :rocket:

```go
import (
	"github.com/nanmu42/etherscan-api/v2/pkg/client"
	"github.com/nanmu42/etherscan-api/v2/pkg/chain"
	"fmt"
)

func main() {
	// create a API client for specified ethereum net
	// there are many pre-defined network in package
	client := client.NewClient(chain.EthereumMainnet, "[your API key]")

	// or, if you are working with antoher chain
	// client := client.NewClient(chain.OpMainnet, "[your API key]")
	//
	// or more customized
	// client := etherscan.NewCustomized(etherscan.Customization{
	// Timeout:       15 * time.Second,
	// Key:           "You key here",
	// Chain:         chain.NewChain(<chain number>),
	// BaseURL:       "<whatever thid-party api provider>",
	// Verbose:       false,
	// })

	// (optional) add hooks, e.g. for rate limit
	client.BeforeRequest = func(module, action string, values url.Values) error {
		// ...
	}
	client.AfterRequest = func(module, action string, values url.Values, outcome interface{}, requestErr error) error {
		// ...
	}

	// check account balance
	balance, err := client.AccountBalance("0x281055afc982d96fab65b3a49cac8b878184cb16")
	if err != nil {
		panic(err)
	}
	// balance in wei, in *big.Int type
	fmt.Println(balance.Int())

	// check token balance
	tokenBalance, err := client.TokenBalance("contractAddress", "holderAddress")

	// check ERC20 transactions from/to a specified address
	transfers, err := client.ERC20Transfers("contractAddress", "address", startBlock, endBlock, page, offset)
}
```

You may find full method list at [GoDoc](https://godoc.org/github.com/TokenTax/etherscan-api/v2).

# Etherscan API Key

You may apply for an API key on [etherscan](https://etherscan.io/apis).

> The Etherscan Ethereum Developer APIs are provided as a community service and without warranty, so please just use what you need and no more. They support both GET/POST requests and a rate limit of 5 requests/sec (exceed and you will be blocked).

# Paperwork Things

This library is not affiliated with Etherscan.io, it's developed for internal use in TokenTax.

# License

Use of this work is governed by an MIT License.
