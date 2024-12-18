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
	"github.com/timcki/etherscan-api/pkg/response"
)

 type ContractParams struct {
	Address string `json:"address"`
 }

 func (p ContractParams) GetUrlValues() url.Values {
	values := url.Values{}
	if p.Address != "" {
		values.Add("address", p.Address)
	}
	return values
 }

 // ContractABI gets contract abi for verified contract source codes
 func (c *Client) ContractABI(address string) (string, error) {
	param := ContractParams{
		Address: address,
	}

	body, err := c.execute("contract", "getabi", param.GetUrlValues())
	if err != nil {
		return "", errors.Wrap(err, "executing ContractABI request")
	}
	return response.ReadResponse[string](body)
 }

 // ContractSource gets contract source code for verified contract source codes
 func (c *Client) ContractSource(address string) ([]response.ContractSource, error) {
	param := ContractParams{
		Address: address,
	}

	body, err := c.execute("contract", "getsourcecode", param.GetUrlValues())
	if err != nil {
		return nil, errors.Wrap(err, "executing ContractSource request")
	}
	return response.ReadResponse[[]response.ContractSource](body)
 }
