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
	"strconv"

	"github.com/pkg/errors"
	"github.com/timcki/etherscan-api/v2/pkg/response"
)

type BlockRewardParams struct {
	BlockNo int `json:"blockno"`
}

type BlockNumberParams struct {
	Timestamp int64  `json:"timestamp"`
	Closest   string `json:"closest"`
}

func (p BlockRewardParams) GetUrlValues() url.Values {
	values := url.Values{}
	values.Add("blockno", strconv.Itoa(p.BlockNo))
	return values
}

func (p BlockNumberParams) GetUrlValues() url.Values {
	values := url.Values{}
	values.Add("timestamp", strconv.FormatInt(p.Timestamp, 10))
	values.Add("closest", p.Closest)
	return values
}

// BlockReward gets block and uncle rewards by block number
func (c *Client) BlockReward(blockNum int) (response.BlockRewards, error) {
	param := BlockRewardParams{
		BlockNo: blockNum,
	}

	body, err := c.execute("block", "getblockreward", param.GetUrlValues())
	if err != nil {
		return response.BlockRewards{}, errors.Wrap(err, "executing BlockReward request")
	}
	return response.ReadResponse[response.BlockRewards](body)
}

// BlockNumber gets the closest block number by UNIX timestamp
//
// valid closest option: before, after
func (c *Client) BlockNumber(timestamp int64, closest string) (int, error) {
	param := BlockNumberParams{
		Timestamp: timestamp,
		Closest:   closest,
	}

	body, err := c.execute("block", "getblocknobytime", param.GetUrlValues())
	if err != nil {
		return 0, errors.Wrap(err, "executing BlockNumber request")
	}

	blockNumberStr, err := response.ReadResponse[string](body)
	if err != nil {
		return 0, errors.Wrap(err, "reading response")
	}

	blockNumber, err := strconv.Atoi(blockNumberStr)
	if err != nil {
		return 0, fmt.Errorf("parsing block number %q: %w", blockNumberStr, err)
	}

	return blockNumber, nil
}
