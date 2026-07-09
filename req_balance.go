package go_ccoop_v2

import (
	"fmt"

	jsoniter "github.com/json-iterator/go"
	"github.com/listenfengyang/go-ccoop-v2/utils"
)

// GetBalance retrieves the current deposit and withdraw balances.
//
// API: GET /api/v1/client/get_balance
//
// For this endpoint, the x-signature is generated with empty orderId and amount:
// base64("{SECRET_KEY}:\"\":\"\"")
func (cli *Client) GetBalance() (*CCoopV2BalanceResponse, error) {
	rawURL := cli.Params.BaseUrl + EndpointGetBalance

	// For get_balance, orderId and amount are empty strings in the signature
	headers := cli.getHeadersEmpty()

	var result CCoopV2BalanceResponse

	resp, err := cli.ryClient.
		SetCloseConnection(true).
		R().
		SetHeaders(headers).
		SetDebug(cli.debugMode).
		SetLogger(cli.logger).
		SetResult(&result).
		SetError(&result).
		Get(rawURL)

	restLog, _ := jsoniter.ConfigCompatibleWithStandardLibrary.Marshal(utils.GetRestyLog(resp))
	cli.logger.Infof("PSPResty#ccoopv2#getBalance->%+v", string(restLog))

	if err != nil {
		return nil, err
	}

	if resp.StatusCode() != 200 {
		var errBody struct {
			Detail string `json:"detail"`
		}
		_ = jsoniter.ConfigCompatibleWithStandardLibrary.Unmarshal(resp.Body(), &errBody)
		if errBody.Detail != "" {
			return nil, fmt.Errorf("get balance failed: %s", errBody.Detail)
		}
		return nil, fmt.Errorf("get balance failed, status code: %d, body: %s", resp.StatusCode(), resp.Body())
	}

	return &result, nil
}
