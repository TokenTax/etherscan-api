/*
 * Copyright (c) 2018 LI Zhennan
 *
 * Use of this work is governed by a MIT License.
 * You may find a license copy in project root.
 */

package client

import (
	"net/url"
	"strconv"
	"strings"

	"github.com/pkg/errors"
	"github.com/timcki/etherscan-api/internal/types"
	"github.com/timcki/etherscan-api/pkg/response"
)

type AccountBalanceParams struct {
	Tag     string `json:"tag"`
	Address string `json:"address"`
}

type MultiAccountBalanceParams struct {
	Tag       string   `json:"tag"`
	Addresses []string `json:"address"`
}

type TxListParams struct {
	Address    string `json:"address"`
	StartBlock *int   `json:"startblock,omitempty"`
	EndBlock   *int   `json:"endblock,omitempty"`
	Page       int    `json:"page"`
	Offset     int    `json:"offset"`
	Sort       string `json:"sort"`
}

type TokenTransferParams struct {
	ContractAddress *string `json:"contractaddress,omitempty"`
	Address         *string `json:"address,omitempty"`
	StartBlock      *int    `json:"startblock,omitempty"`
	EndBlock        *int    `json:"endblock,omitempty"`
	Page            int     `json:"page"`
	Offset          int     `json:"offset"`
	Sort            string  `json:"sort"`
}

type MinedBlockParams struct {
	Address   string `json:"address"`
	BlockType string `json:"blocktype"`
	Page      int    `json:"page"`
	Offset    int    `json:"offset"`
}

type TokenBalanceParams struct {
	ContractAddress string `json:"contractaddress"`
	Address         string `json:"address"`
	Tag             string `json:"tag"`
}

func (p AccountBalanceParams) GetUrlValues() url.Values {
	values := url.Values{}
	if p.Tag != "" {
		values.Add("tag", p.Tag)
	}
	if p.Address != "" {
		values.Add("address", p.Address)
	}
	return values
}

func (p MultiAccountBalanceParams) GetUrlValues() url.Values {
	values := url.Values{}
	if p.Tag != "" {
		values.Add("tag", p.Tag)
	}
	if len(p.Addresses) > 0 {
		values.Add("address", strings.Join(p.Addresses, ","))
	}
	return values
}

func (p TxListParams) GetUrlValues() url.Values {
	values := url.Values{}
	if p.Address != "" {
		values.Add("address", p.Address)
	}
	if p.StartBlock != nil {
		values.Add("startblock", strconv.Itoa(*p.StartBlock))
	}
	if p.EndBlock != nil {
		values.Add("endblock", strconv.Itoa(*p.EndBlock))
	}
	values.Add("page", strconv.Itoa(p.Page))
	values.Add("offset", strconv.Itoa(p.Offset))
	if p.Sort != "" {
		values.Add("sort", p.Sort)
	}
	return values
}

func (p TokenTransferParams) GetUrlValues() url.Values {
	values := url.Values{}
	if p.ContractAddress != nil {
		values.Add("contractaddress", *p.ContractAddress)
	}
	if p.Address != nil {
		values.Add("address", *p.Address)
	}
	if p.StartBlock != nil {
		values.Add("startblock", strconv.Itoa(*p.StartBlock))
	}
	if p.EndBlock != nil {
		values.Add("endblock", strconv.Itoa(*p.EndBlock))
	}
	values.Add("page", strconv.Itoa(p.Page))
	values.Add("offset", strconv.Itoa(p.Offset))
	if p.Sort != "" {
		values.Add("sort", p.Sort)
	}
	return values
}

func (p MinedBlockParams) GetUrlValues() url.Values {
	values := url.Values{}
	if p.Address != "" {
		values.Add("address", p.Address)
	}
	if p.BlockType != "" {
		values.Add("blocktype", p.BlockType)
	}
	values.Add("page", strconv.Itoa(p.Page))
	values.Add("offset", strconv.Itoa(p.Offset))
	return values
}

func (p TokenBalanceParams) GetUrlValues() url.Values {
	values := url.Values{}
	if p.ContractAddress != "" {
		values.Add("contractaddress", p.ContractAddress)
	}
	if p.Address != "" {
		values.Add("address", p.Address)
	}
	if p.Tag != "" {
		values.Add("tag", p.Tag)
	}
	return values
}

// Refactored methods
func (c *Client) AccountBalance(address string) (response.AccountBalance, error) {
	param := AccountBalanceParams{
		Tag:     "latest",
		Address: address,
	}
	body, err := c.execute("account", "balance", param.GetUrlValues())
	if err != nil {
		return response.AccountBalance{}, errors.Wrap(err, "executing AccountBalance request")
	}
	return response.ReadResponse[response.AccountBalance](body)
}

func (c *Client) MultiAccountBalance(addresses ...string) ([]response.AccountBalance, error) {
	param := MultiAccountBalanceParams{
		Tag:       "latest",
		Addresses: addresses,
	}
	body, err := c.execute("account", "balancemulti", param.GetUrlValues())
	if err != nil {
		return []response.AccountBalance{}, errors.Wrap(err, "executing MultiAccountBalance request")
	}
	return response.ReadResponse[[]response.AccountBalance](body)
}

func (c *Client) NormalTxByAddress(address string, startBlock *int, endBlock *int, page int, offset int, desc bool) ([]response.NormalTx, error) {
	param := TxListParams{
		Address:    address,
		StartBlock: startBlock,
		EndBlock:   endBlock,
		Page:       page,
		Offset:     offset,
		Sort:       "asc",
	}
	if desc {
		param.Sort = "desc"
	}
	body, err := c.execute("account", "txlist", param.GetUrlValues())
	if err != nil {
		return []response.NormalTx{}, errors.Wrap(err, "executing NormalTxByAddress request")
	}
	return response.ReadResponse[[]response.NormalTx](body)
}

func (c *Client) InternalTxByAddress(address string, startBlock *int, endBlock *int, page int, offset int, desc bool) ([]response.InternalTx, error) {
	param := TxListParams{
		Address:    address,
		StartBlock: startBlock,
		EndBlock:   endBlock,
		Page:       page,
		Offset:     offset,
		Sort:       "asc",
	}
	if desc {
		param.Sort = "desc"
	}

	body, err := c.execute("account", "txlistinternal", param.GetUrlValues())
	if err != nil {
		return []response.InternalTx{}, errors.Wrap(err, "executing InternalTxByAddress request")
	}
	return response.ReadResponse[[]response.InternalTx](body)
}

func (c *Client) ERC20Transfers(contractAddress, address *string, startBlock *int, endBlock *int, page int, offset int, desc bool) ([]response.ERC20Transfer, error) {
	param := TokenTransferParams{
		ContractAddress: contractAddress,
		Address:         address,
		StartBlock:      startBlock,
		EndBlock:        endBlock,
		Page:            page,
		Offset:          offset,
		Sort:            "asc",
	}
	if desc {
		param.Sort = "desc"
	}
	body, err := c.execute("account", "tokentx", param.GetUrlValues())
	if err != nil {
		return []response.ERC20Transfer{}, errors.Wrap(err, "executing ERC20Transfers request")
	}
	return response.ReadResponse[[]response.ERC20Transfer](body)
}

func (c *Client) ERC721Transfers(contractAddress, address *string, startBlock *int, endBlock *int, page int, offset int, desc bool) ([]response.ERC721Transfer, error) {
	param := TokenTransferParams{
		ContractAddress: contractAddress,
		Address:         address,
		StartBlock:      startBlock,
		EndBlock:        endBlock,
		Page:            page,
		Offset:          offset,
		Sort:            "asc",
	}
	if desc {
		param.Sort = "desc"
	}
	body, err := c.execute("account", "tokennfttx", param.GetUrlValues())
	if err != nil {
		return []response.ERC721Transfer{}, errors.Wrap(err, "executing ERC721Transfers request")
	}
	return response.ReadResponse[[]response.ERC721Transfer](body)
}

func (c *Client) ERC1155Transfers(contractAddress, address *string, startBlock *int, endBlock *int, page int, offset int, desc bool) ([]response.ERC1155Transfer, error) {
	param := TokenTransferParams{
		ContractAddress: contractAddress,
		Address:         address,
		StartBlock:      startBlock,
		EndBlock:        endBlock,
		Page:            page,
		Offset:          offset,
		Sort:            "asc",
	}
	if desc {
		param.Sort = "desc"
	}
	body, err := c.execute("account", "token1155tx", param.GetUrlValues())
	if err != nil {
		return []response.ERC1155Transfer{}, errors.Wrap(err, "executing ERC1155Transfers request")
	}
	return response.ReadResponse[[]response.ERC1155Transfer](body)
}

func (c *Client) BlocksMinedByAddress(address string, page int, offset int) ([]response.MinedBlock, error) {
	param := MinedBlockParams{
		Address:   address,
		BlockType: "blocks",
		Page:      page,
		Offset:    offset,
	}
	body, err := c.execute("account", "getminedblocks", param.GetUrlValues())
	if err != nil {
		return []response.MinedBlock{}, errors.Wrap(err, "executing BlocksMinedByAddress request")
	}
	return response.ReadResponse[[]response.MinedBlock](body)
}

func (c *Client) UnclesMinedByAddress(address string, page int, offset int) ([]response.MinedBlock, error) {
	param := MinedBlockParams{
		Address:   address,
		BlockType: "uncles",
		Page:      page,
		Offset:    offset,
	}
	body, err := c.execute("account", "getminedblocks", param.GetUrlValues())
	if err != nil {
		return []response.MinedBlock{}, errors.Wrap(err, "executing UnclesMinedByAddress request")
	}
	return response.ReadResponse[[]response.MinedBlock](body)
}

func (c *Client) TokenBalance(contractAddress, address string) (types.BigInt, error) {
	param := TokenBalanceParams{
		ContractAddress: contractAddress,
		Address:         address,
		Tag:             "latest",
	}
	body, err := c.execute("account", "tokenbalance", param.GetUrlValues())
	if err != nil {
		return types.BigInt{}, errors.Wrap(err, "executing TokenBalance request")
	}
	return response.ReadResponse[types.BigInt](body)
}
