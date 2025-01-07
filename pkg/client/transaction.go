/*
 * Copyright (c) 2018 LI Zhennan
 *
 * Use of this work is governed by a MIT License.
 * You may find a license copy in project root.
 */

package client

import (
	"net/url"

	"github.com/TokenTax/etherscan-api/v2/pkg/response"
	"github.com/pkg/errors"
)

// ErrPreByzantiumTx transaction before 4,370,000 does not support receipt status check
var ErrPreByzantiumTx = errors.New("pre-byzantium transaction does not support receipt status check")

type TransactionParams struct {
	TxHash string `json:"txhash"`
}

func (p TransactionParams) GetUrlValues() url.Values {
	values := url.Values{}
	if p.TxHash != "" {
		values.Add("txhash", p.TxHash)
	}
	return values
}

// ExecutionStatus checks contract execution status
func (c *Client) ExecutionStatus(txHash string) (response.ExecutionStatus, error) {
	param := TransactionParams{TxHash: txHash}

	body, err := c.execute("transaction", "getstatus", param.GetUrlValues())
	if err != nil {
		return response.ExecutionStatus{}, errors.Wrap(err, "executing ExecutionStatus request")
	}
	return response.ReadResponse[response.ExecutionStatus](body)
}

// ReceiptStatus checks transaction receipt status
func (c *Client) ReceiptStatus(txHash string) (int, error) {
	param := TransactionParams{TxHash: txHash}
	body, err := c.execute("transaction", "gettxreceiptstatus", param.GetUrlValues())
	if err != nil {
		return 0, errors.Wrap(err, "executing ReceiptStatus request")
	}

	rawStatus, err := response.ReadResponse[response.StatusReponse](body)
	if err != nil {
		return 0, err
	}

	switch rawStatus.Status {
	case "0":
		return 0, nil
	case "1":
		return 1, nil
	default:
		return -1, ErrPreByzantiumTx
	}
}
