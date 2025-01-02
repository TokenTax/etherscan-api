/*
 * Copyright (c) 2018 LI Zhennan
 *
 * Use of this work is governed by a MIT License.
 * You may find a license copy in project root.
 */

package client

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"net/http/httputil"
	"net/url"
	"strconv"
	"time"

	"github.com/pkg/errors"
	"github.com/timcki/etherscan-api/internal/chain"
)

type (
	// Client etherscan API client
	// Clients are safe for concurrent use by multiple goroutines.
	Client struct {
		conn    *http.Client
		key     string
		baseURL string
		chain   chain.Chain

		// Verbose when true, talks a lot
		Verbose bool

		// BeforeRequest runs before every client request, in the same goroutine.
		// May be used in rate limit.
		// Request will be aborted, if BeforeRequest returns non-nil err.
		BeforeRequest func(module, action string, values url.Values) error

		// AfterRequest runs after every client request, even when there is an error.
		AfterRequest func(module, action string, values url.Values, outcome interface{}, requestErr error) error
	}

	// Customization is used in NewCustomized()
	Customization struct {
		// Timeout for API call
		Timeout time.Duration
		// API key applied from Etherscan
		Key string
		// Base URL like `https://api.etherscan.io/api?`
		BaseURL string
		// When true, talks a lot
		Verbose bool
		// ChainID to be used
		Chain chain.Chain
		// HTTP Client to be used. Specifying this value will ignore the Timeout value set
		// Set your own timeout.
		Client *http.Client

		// BeforeRequest runs before every client request, in the same goroutine.
		// May be used in rate limit.
		// Request will be aborted, if BeforeRequest returns non-nil err.
		BeforeRequest func(module, action string, values url.Values) error

		// AfterRequest runs after every client request, even when there is an error.
		AfterRequest func(module, action string, values url.Values, outcome interface{}, requestErr error) error
	}
)

// NewClient initialize a new etherscan API client
// please use pre-defined network value
func NewClient(chain chain.Chain, APIKey string) *Client {
	return NewCustomized(Customization{
		Timeout: 30 * time.Second,
		Key:     APIKey,
		Chain:   chain,
		BaseURL: `https://api.etherscan.io/v2/api`,
	})
}

// NewCustomized initialize a customized API client,
// useful when calling against etherscan-family API like BscScan.
func NewCustomized(config Customization) *Client {
	var httpClient *http.Client
	if config.Client != nil {
		httpClient = config.Client
	} else {
		httpClient = &http.Client{Timeout: config.Timeout}
	}
	return &Client{
		conn:          httpClient,
		key:           config.Key,
		chain:         config.Chain,
		baseURL:       config.BaseURL,
		Verbose:       config.Verbose,
		BeforeRequest: config.BeforeRequest,
		AfterRequest:  config.AfterRequest,
	}
}

func (c *Client) execute(module, action string, values url.Values) (bytes.Buffer, error) {
	var content = bytes.Buffer{}

	req, err := http.NewRequest(http.MethodGet, c.craftURL(module, action, values), http.NoBody)
	if err != nil {
		return content, errors.Wrap(err, "creating request")
	}
	req.Header.Set("User-Agent", "etherscan-api(Go)")
	req.Header.Set("Content-Type", "application/json; charset=utf-8")

	if c.Verbose {
		reqDump, err := httputil.DumpRequestOut(req, false)
		if err != nil {
			return content, errors.Wrap(err, "verbose mode: dumping request")
		}

		fmt.Printf("\n%s\n", reqDump)

		defer func() {
			if err != nil {
				fmt.Printf("[Error] %v\n", err)
			}
		}()
	}

	res, err := c.conn.Do(req)
	if err != nil {
		return content, errors.Wrap(err, "sending request")
	}
	defer res.Body.Close()

	if c.Verbose {
		resDump, err := httputil.DumpResponse(res, true)
		if err != nil {
			return content, errors.Wrap(err, "verbose mode:dumping response")
		}

		fmt.Printf("%s\n", resDump)
	}

	if _, err = io.Copy(&content, res.Body); err != nil {
		return content, errors.Wrap(err, "reading response")
	}

	if res.StatusCode != http.StatusOK {
		return content, errors.Errorf("got non-200 status code; status: %v, status text: %s, response body: %s", res.StatusCode, res.Status, content.String())
	}

	return content, nil
}

/*
func (c *Client) innerCall(module, action string, values url.Values, outcome any) error {
	req, err := http.NewRequest(http.MethodGet, c.craftURL(module, action, values), http.NoBody)
	if err != nil {
		return errors.Wrap(err, "creating request")
	}
	req.Header.Set("User-Agent", "etherscan-api(Go)")
	req.Header.Set("Content-Type", "application/json; charset=utf-8")

	if c.Verbose {
		reqDump, err := httputil.DumpRequestOut(req, false)
		if err != nil {
			return errors.Wrap(err, "verbose mode: dumping request")
		}

		fmt.Printf("\n%s\n", reqDump)

		defer func() {
			if err != nil {
				fmt.Printf("[Error] %v\n", err)
			}
		}()
	}

	res, err := c.conn.Do(req)
	if err != nil {
		return errors.Wrap(err, "sending request")
	}
	defer res.Body.Close()

	if c.Verbose {
		resDump, err := httputil.DumpResponse(res, true)
		if err != nil {
			return errors.Wrap(err, "verbose mode:dumping response")
		}

		fmt.Printf("%s\n", resDump)
	}

	resp := response.ReadResponse(res)

	if res.StatusCode != http.StatusOK {
		return errors.Errorf("got non-200 status code; status: %v, status text: %s, response body: %s", res.StatusCode, res.Status, content.String())
	}

	var content bytes.Buffer
	if _, err = io.Copy(&content, res.Body); err != nil {
		return errors.Wrap(err, "reading response")
	}

	var envelope Envelope
	if err = json.Unmarshal(content.Bytes(), &envelope); err != nil {
		return errors.Wrap(err, "unmarshaling response")
	}
	if envelope.Status != 1 {
		return errors.Errorf("etherscan server: %s", envelope.Message)
	}

	// workaround for missing tokenDecimal for some tokentx calls
	if action == "tokentx" {
		err = json.Unmarshal(bytes.Replace(envelope.Result, []byte(`"tokenDecimal":""`), []byte(`"tokenDecimal":"0"`), -1), outcome)
	} else {
		err = json.Unmarshal(envelope.Result, outcome)
	}

	return errors.Wrap(err, "unmarshaling result")
}

// call executes the Before/AfterRequest hooks and executes innerCall in between
func (c *Client) call(module, action string, values url.Values, outcome interface{}) error {
	// fire hooks if in need
	if c.BeforeRequest != nil {
		if err := c.BeforeRequest(module, action, values); err != nil {
			return errors.Wrap(err, "running beforeRequest")
		}
	}

	err := c.innerCall(module, action, values, outcome)
	if c.AfterRequest != nil {
		if err := c.AfterRequest(module, action, param, outcome, err); err != nil {
			return errors.Wrapf(err, "running afterRequest with request Error: %v", err)
		}
	}
	return err
}
*/

// craftURL returns desired URL via param provided
func (c *Client) craftURL(module, action string, values url.Values) string {
	values.Add("module", module)
	values.Add("action", action)
	values.Add("apikey", c.key)
	values.Add("chainid", strconv.Itoa(c.chain.ChainID()))

	return fmt.Sprintf("%s?%s", c.baseURL, values.Encode())
}
