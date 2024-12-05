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
func (c *Client) AccountBalance(address string) (balance *types.BigInt, err error) {
	param := AccountBalanceParams{
		Tag:     "latest",
		Address: address,
	}
	balance = new(types.BigInt)
	err = c.call("account", "balance", param, balance)
	return
}

func (c *Client) MultiAccountBalance(addresses ...string) (balances []response.AccountBalance, err error) {
	param := MultiAccountBalanceParams{
		Tag:       "latest",
		Addresses: addresses,
	}
	//balances = make([]response.AccountBalance, 0, len(addresses))
	err = c.call("account", "balancemulti", param, &balances)
	return
}

func (c *Client) NormalTxByAddress(address string, startBlock *int, endBlock *int, page int, offset int, desc bool) (txs []NormalTx, err error) {
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
	err = c.call("account", "txlist", param, &txs)
	return
}

func (c *Client) InternalTxByAddress(address string, startBlock *int, endBlock *int, page int, offset int, desc bool) (txs []InternalTx, err error) {
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
	err = c.call("account", "txlistinternal", param, &txs)
	return
}

func (c *Client) ERC20Transfers(contractAddress, address *string, startBlock *int, endBlock *int, page int, offset int, desc bool) (txs []ERC20Transfer, err error) {
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
	err = c.call("account", "tokentx", param, &txs)
	return
}

func (c *Client) ERC721Transfers(contractAddress, address *string, startBlock *int, endBlock *int, page int, offset int, desc bool) (txs []ERC721Transfer, err error) {
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
	err = c.call("account", "tokennfttx", param, &txs)
	return
}

func (c *Client) ERC1155Transfers(contractAddress, address *string, startBlock *int, endBlock *int, page int, offset int, desc bool) (txs []ERC1155Transfer, err error) {
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
	err = c.call("account", "token1155tx", param, &txs)
	return
}

func (c *Client) BlocksMinedByAddress(address string, page int, offset int) (mined []MinedBlock, err error) {
	param := MinedBlockParams{
		Address:   address,
		BlockType: "blocks",
		Page:      page,
		Offset:    offset,
	}
	err = c.call("account", "getminedblocks", param, &mined)
	return
}

func (c *Client) UnclesMinedByAddress(address string, page int, offset int) (mined []MinedBlock, err error) {
	param := MinedBlockParams{
		Address:   address,
		BlockType: "uncles",
		Page:      page,
		Offset:    offset,
	}
	err = c.call("account", "getminedblocks", param, &mined)
	return
}

func (c *Client) TokenBalance(contractAddress, address string) (balance *BigInt, err error) {
	param := TokenBalanceParams{
		ContractAddress: contractAddress,
		Address:         address,
		Tag:             "latest",
	}
	err = c.call("account", "tokenbalance", param, &balance)
	return
}
