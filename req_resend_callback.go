package go_ccoop_v2

import (
	"fmt"

	jsoniter "github.com/json-iterator/go"
	"github.com/listenfengyang/go-ccoop-v2/utils"
)

// CCoopV2ResendCallbackRequest is the request body for resending a callback for a transaction.
type CCoopV2ResendCallbackRequest struct {
	RefId string `json:"ref_id"` // The internal reference ID of the transaction (required)
}

// CCoopV2ResendCallbackResponse is the response from the resend callback API.
type CCoopV2ResendCallbackResponse struct {
	Detail string `json:"detail"` // e.g. "resend callback success"
}

// ResendCallback manually triggers a callback resend for a transaction by its ref_id.
//
// API: POST /api/v1/client/resend_callback_for_transaction
//
// The signature is generated with empty orderId and amount:
// base64("{SECRET_KEY}:\"\":\"\"")
//
// Use this when your callback endpoint missed a notification and you need to
// re-trigger it without creating a new transaction.
func (cli *Client) ResendCallback(req CCoopV2ResendCallbackRequest) (*CCoopV2ResendCallbackResponse, error) {
	rawURL := cli.Params.BaseUrl + EndpointResendCallbackForTxn

	// For resend_callback, orderId and amount are empty strings in the signature
	headers := cli.getHeadersEmpty()

	body := map[string]interface{}{
		"ref_id": req.RefId,
	}

	var result CCoopV2ResendCallbackResponse

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
	cli.logger.Infof("PSPResty#ccoopv2#resendCallback->%+v", string(restLog))

	if err != nil {
		return nil, err
	}

	if resp.StatusCode() != 200 {
		var errBody struct {
			Detail string `json:"detail"`
		}
		_ = jsoniter.ConfigCompatibleWithStandardLibrary.Unmarshal(resp.Body(), &errBody)
		if errBody.Detail != "" {
			return nil, fmt.Errorf("resend callback failed: %s", errBody.Detail)
		}
		return nil, fmt.Errorf("resend callback failed, status code: %d, body: %s", resp.StatusCode(), resp.Body())
	}

	return &result, nil
}
