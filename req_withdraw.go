package go_ccoop_v2

import (
	"fmt"
	"strconv"

	jsoniter "github.com/json-iterator/go"
	"github.com/listenfengyang/go-ccoop-v2/utils"
)

// Withdraw creates a withdrawal transaction.
//
// API: POST /api/v1/client/create_withdraw
//
// The signature is generated as: base64("{SECRET_KEY}:{order_id}:{amount}")
//
// ⚠️  Critical rules:
//   - Do NOT set an expiration time on withdrawal requests.
//   - To re-check status, use the Resend Callback endpoint only.
//   - Do NOT send retries under any circumstances.
func (cli *Client) Withdraw(req CCoopV2WithdrawRequest) (*CCoopV2WithdrawResponse, error) {
	// Use WithdrawBaseUrl if configured, otherwise fall back to BaseUrl
	baseUrl := cli.Params.WithdrawBaseUrl
	if baseUrl == "" {
		baseUrl = cli.Params.BaseUrl
	}
	rawURL := baseUrl + EndpointCreateWithdraw

	// Fill in the default callback URL if not specified in the request
	if req.CallbackUrl == "" && cli.Params.WithdrawCallbackUrl != "" {
		req.CallbackUrl = cli.Params.WithdrawCallbackUrl
	}

	// Build request body
	body := map[string]interface{}{
		"amount":             req.Amount,
		"dest_bank_acc_no":   req.DestBankAccNo,
		"dest_bank_acc_name": req.DestBankAccName,
		"dest_bank_code":     req.DestBankCode,
		"withdraw_code":      req.WithdrawCode,
		"callback_url":       req.CallbackUrl,
		"order_id":           req.OrderId,
	}

	// Signature uses order_id and amount
	amountStr := strconv.FormatFloat(req.Amount, 'f', 2, 64)
	headers := cli.getHeaders(req.OrderId, amountStr)

	var result CCoopV2WithdrawResponse

	resp, err := cli.ryClient.
		SetCloseConnection(true).
		R().
		SetBody(body).
		SetHeaders(headers).
		SetDebug(cli.debugMode).
		SetLogger(cli.logger).
		SetResult(&result).
		SetError(&result).
		Post(rawURL)

	restLog, _ := jsoniter.ConfigCompatibleWithStandardLibrary.Marshal(utils.GetRestyLog(resp))
	cli.logger.Infof("PSPResty#ccoopv2#withdraw->%+v", string(restLog))

	if err != nil {
		return nil, err
	}

	if resp.StatusCode() != 200 {
		var errBody struct {
			Detail string `json:"detail"`
		}
		_ = jsoniter.ConfigCompatibleWithStandardLibrary.Unmarshal(resp.Body(), &errBody)
		if errBody.Detail != "" {
			return nil, fmt.Errorf("withdraw failed: %s", errBody.Detail)
		}
		return nil, fmt.Errorf("withdraw failed, status code: %d, body: %s", resp.StatusCode(), resp.Body())
	}

	return &result, nil
}
