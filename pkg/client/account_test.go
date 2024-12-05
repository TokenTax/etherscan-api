package client

import (
	"net/url"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAccountBalanceParams_GetUrlValues(t *testing.T) {
	tests := []struct {
		name     string
		params   AccountBalanceParams
		expected url.Values
	}{
		{
			name:     "empty params",
			params:   AccountBalanceParams{},
			expected: url.Values{},
		},
		{
			name: "full params",
			params: AccountBalanceParams{
				Tag:     "latest",
				Address: "0x742d35Cc6634C0532925a3b844Bc454e4438f44e",
			},
			expected: url.Values{
				"tag":     []string{"latest"},
				"address": []string{"0x742d35Cc6634C0532925a3b844Bc454e4438f44e"},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.params.GetUrlValues()
			assert.Equal(t, tt.expected, result)
		})
	}
}

func TestMultiAccountBalanceParams_GetUrlValues(t *testing.T) {
	tests := []struct {
		name     string
		params   MultiAccountBalanceParams
		expected url.Values
	}{
		{
			name:     "empty params",
			params:   MultiAccountBalanceParams{},
			expected: url.Values{},
		},
		{
			name: "full params",
			params: MultiAccountBalanceParams{
				Tag: "latest",
				Addresses: []string{
					"0x742d35Cc6634C0532925a3b844Bc454e4438f44e",
					"0x742d35Cc6634C0532925a3b844Bc454e4438f44f",
				},
			},
			expected: url.Values{
				"tag":     []string{"latest"},
				"address": []string{"0x742d35Cc6634C0532925a3b844Bc454e4438f44e,0x742d35Cc6634C0532925a3b844Bc454e4438f44f"},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.params.GetUrlValues()
			assert.Equal(t, tt.expected, result)
		})
	}
}

func TestTxListParams_GetUrlValues(t *testing.T) {
	startBlock := 12345
	endBlock := 12346

	tests := []struct {
		name     string
		params   TxListParams
		expected url.Values
	}{
		{
			name: "minimal params",
			params: TxListParams{
				Address: "0x742d35Cc6634C0532925a3b844Bc454e4438f44e",
				Page:    1,
				Offset:  10,
			},
			expected: url.Values{
				"address": []string{"0x742d35Cc6634C0532925a3b844Bc454e4438f44e"},
				"page":    []string{"1"},
				"offset":  []string{"10"},
			},
		},
		{
			name: "full params",
			params: TxListParams{
				Address:    "0x742d35Cc6634C0532925a3b844Bc454e4438f44e",
				StartBlock: &startBlock,
				EndBlock:   &endBlock,
				Page:       1,
				Offset:     10,
				Sort:       "asc",
			},
			expected: url.Values{
				"address":    []string{"0x742d35Cc6634C0532925a3b844Bc454e4438f44e"},
				"startblock": []string{"12345"},
				"endblock":   []string{"12346"},
				"page":       []string{"1"},
				"offset":     []string{"10"},
				"sort":       []string{"asc"},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.params.GetUrlValues()
			assert.Equal(t, tt.expected, result)
		})
	}
}

func TestTokenTransferParams_GetUrlValues(t *testing.T) {
	startBlock := 12345
	endBlock := 12346
	contractAddr := "0x742d35Cc6634C0532925a3b844Bc454e4438f44e"
	address := "0x742d35Cc6634C0532925a3b844Bc454e4438f44f"

	tests := []struct {
		name     string
		params   TokenTransferParams
		expected url.Values
	}{
		{
			name: "minimal params",
			params: TokenTransferParams{
				Page:   1,
				Offset: 10,
			},
			expected: url.Values{
				"page":   []string{"1"},
				"offset": []string{"10"},
			},
		},
		{
			name: "full params",
			params: TokenTransferParams{
				ContractAddress: &contractAddr,
				Address:         &address,
				StartBlock:      &startBlock,
				EndBlock:        &endBlock,
				Page:            1,
				Offset:          10,
				Sort:            "asc",
			},
			expected: url.Values{
				"contractaddress": []string{"0x742d35Cc6634C0532925a3b844Bc454e4438f44e"},
				"address":         []string{"0x742d35Cc6634C0532925a3b844Bc454e4438f44f"},
				"startblock":      []string{"12345"},
				"endblock":        []string{"12346"},
				"page":            []string{"1"},
				"offset":          []string{"10"},
				"sort":            []string{"asc"},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.params.GetUrlValues()
			assert.Equal(t, tt.expected, result)
		})
	}
}

func TestMinedBlockParams_GetUrlValues(t *testing.T) {
	tests := []struct {
		name     string
		params   MinedBlockParams
		expected url.Values
	}{
		{
			name: "minimal params",
			params: MinedBlockParams{
				Page:   1,
				Offset: 10,
			},
			expected: url.Values{
				"page":   []string{"1"},
				"offset": []string{"10"},
			},
		},
		{
			name: "full params",
			params: MinedBlockParams{
				Address:   "0x742d35Cc6634C0532925a3b844Bc454e4438f44e",
				BlockType: "blocks",
				Page:      1,
				Offset:    10,
			},
			expected: url.Values{
				"address":   []string{"0x742d35Cc6634C0532925a3b844Bc454e4438f44e"},
				"blocktype": []string{"blocks"},
				"page":      []string{"1"},
				"offset":    []string{"10"},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.params.GetUrlValues()
			assert.Equal(t, tt.expected, result)
		})
	}
}

func TestTokenBalanceParams_GetUrlValues(t *testing.T) {
	tests := []struct {
		name     string
		params   TokenBalanceParams
		expected url.Values
	}{
		{
			name:     "empty params",
			params:   TokenBalanceParams{},
			expected: url.Values{},
		},
		{
			name: "full params",
			params: TokenBalanceParams{
				ContractAddress: "0x742d35Cc6634C0532925a3b844Bc454e4438f44e",
				Address:         "0x742d35Cc6634C0532925a3b844Bc454e4438f44f",
				Tag:             "latest",
			},
			expected: url.Values{
				"contractaddress": []string{"0x742d35Cc6634C0532925a3b844Bc454e4438f44e"},
				"address":         []string{"0x742d35Cc6634C0532925a3b844Bc454e4438f44f"},
				"tag":             []string{"latest"},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.params.GetUrlValues()
			assert.Equal(t, tt.expected, result)
		})
	}
}
