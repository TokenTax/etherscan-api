/*
 * Copyright (c) 2018 LI Zhennan
 *
 * Use of this work is governed by a MIT License.
 * You may find a license copy in project root.
 */

package client

import (
	"net/url"

	"github.com/pkg/errors"
	"github.com/timcki/etherscan-api/v2/internal/types"
	"github.com/timcki/etherscan-api/v2/pkg/response"
)

type TokenTotalSupplyParams struct {
	ContractAddress string `json:"contractaddress"`
}

func (p TokenTotalSupplyParams) GetUrlValues() url.Values {
	values := url.Values{}
	if p.ContractAddress != "" {
		values.Add("contractaddress", p.ContractAddress)
	}
	return values
}

// EtherTotalSupply gets total supply of ether
func (c *Client) EtherTotalSupply() (totalSupply types.BigInt, err error) {
	body, err := c.execute("stats", "ethsupply", nil)
	if err != nil {
		return types.BigInt{}, errors.Wrap(err, "executing ExecutionStatus request")
	}
	return response.ReadResponse[types.BigInt](body)
}

// EtherLatestPrice gets the latest ether price, in BTC and USD
func (c *Client) EtherLatestPrice() (price response.LatestPrice, err error) {
	body, err := c.execute("stats", "ethprice", nil)
	if err != nil {
		return response.LatestPrice{}, errors.Wrap(err, "executing LatestPrice request")
	}
	return response.ReadResponse[response.LatestPrice](body)
}

// TokenTotalSupply gets total supply of token on specified contract address
func (c *Client) TokenTotalSupply(contractAddress string) (types.BigInt, error) {
	values := TokenTotalSupplyParams{ContractAddress: contractAddress}

	body, err := c.execute("stats", "tokensupply", values.GetUrlValues())
	if err != nil {
		return types.BigInt{}, errors.Wrap(err, "executing TokenTotalSupply request")
	}
	return response.ReadResponse[types.BigInt](body)
}
