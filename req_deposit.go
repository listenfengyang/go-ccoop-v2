package go_ccoop_v2

import (
	"fmt"
	"strconv"

	jsoniter "github.com/json-iterator/go"
	"github.com/listenfengyang/go-ccoop-v2/utils"
)

// Deposit creates a deposit transaction and returns the deposit details including QR code info.
//
// API: POST /api/v1/client/create_deposit
//
// The signature is generated as: base64("{SECRET_KEY}:{order_id}:{amount}")
// The response includes either a PromptPay QR code or bank transfer details depending on the channel.
func (cli *Client) Deposit(req CCoopV2DepositRequest) (*CCoopV2DepositResponse, error) {
	rawURL := cli.Params.BaseUrl + EndpointCreateDeposit

	// Fill in the default callback URL if not specified in the request
	if req.CallbackUrl == "" && cli.Params.DepositCallbackUrl != "" {
		req.CallbackUrl = cli.Params.DepositCallbackUrl
	}

	// Build request body
	body := map[string]interface{}{
		"order_id":      req.OrderId,
		"amount":        req.Amount,
		"ref_account":   req.RefAccount,
		"ref_bank_code": req.RefBankCode,
		"ref_name_th":   req.RefNameTh,
		"ref_name_en":   req.RefNameEn,
		"ref_user_id":   req.RefUserId,
		"ref1":          req.Ref1,
		"ref2":          req.Ref2,
		"callback_url":  req.CallbackUrl,
	}

	// Signature uses order_id and amount
	amountStr := strconv.FormatFloat(req.Amount, 'f', 2, 64)
	headers := cli.getHeaders(req.OrderId, amountStr)

	var result CCoopV2DepositResponse

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
	cli.logger.Infof("PSPResty#ccoopv2#deposit->%+v", string(restLog))

	if err != nil {
		return nil, err
	}

	if resp.StatusCode() != 200 {
		var errBody struct {
			Detail string `json:"detail"`
		}
		_ = jsoniter.ConfigCompatibleWithStandardLibrary.Unmarshal(resp.Body(), &errBody)
		if errBody.Detail != "" {
			return nil, fmt.Errorf("deposit failed: %s", errBody.Detail)
		}
		return nil, fmt.Errorf("deposit failed, status code: %d, body: %s", resp.StatusCode(), resp.Body())
	}

	return &result, nil
}
