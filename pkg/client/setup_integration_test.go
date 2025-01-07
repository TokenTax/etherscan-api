//go:build integration
// +build integration

/*
 * Copyright (c) 2018 LI Zhennan
 *
 * Use of this work is governed by a MIT License.
 * You may find a license copy in project root.
 */

package client

import (
	"fmt"
	"net/url"
	"os"
	"time"

	"github.com/TokenTax/etherscan-api/v2/pkg/chain"
)

const apiKeyEnvName = "ETHERSCAN_API_KEY"

var (
	// api test client for many test cases
	api *Client
	// bucket default rate limiter
	bucket *Bucket
	// apiKey etherscan API key
	apiKey string
	ok     bool
)

func init() {
	if apiKey, ok = os.LookupEnv(apiKeyEnvName); !ok {
		panic(fmt.Sprintf("API key is empty, set env variable %q with a valid API key to proceed.", apiKeyEnvName))
	}
	bucket = NewBucket(500 * time.Millisecond)

	api = NewClient(chain.EthereumMainnet, apiKey)
	//api.Verbose = true
	api.BeforeRequest = func(module string, action string, values url.Values) error {
		bucket.Take()
		return nil
	}
}

// Bucket is a simple and easy rate limiter
// Use NewBucket() to construct one.
type Bucket struct {
	bucket     chan bool
	refillTime time.Duration
}

// NewBucket factory of Bucket
func NewBucket(refillTime time.Duration) (b *Bucket) {
	b = &Bucket{
		bucket:     make(chan bool),
		refillTime: refillTime,
	}

	go b.fillRoutine()

	return
}

// Take a action token from bucket,
// blocks if there is currently no token left.
func (b *Bucket) Take() {
	<-b.bucket
}

// fill a action token into bucket,
// no-op if the bucket is currently full
func (b *Bucket) fill() {
	b.bucket <- true
}

func (b *Bucket) fillRoutine() {
	ticker := time.NewTicker(b.refillTime)

	for range ticker.C {
		b.fill()
	}
}
