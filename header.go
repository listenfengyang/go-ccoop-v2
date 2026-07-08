package go_ccoop_v2

import "github.com/listenfengyang/go-ccoop-v2/utils"

// getHeaders returns the required HTTP headers for the v2 API.
// x-signature is generated based on secretKey, orderId, and amount.
func (cli *Client) getHeaders(orderId string, amount string) map[string]string {
	sig := utils.SignV2(cli.Params.SecretKey, orderId, amount)
	return map[string]string{
		"Content-Type":  "application/json",
		"Authorization": "Bearer " + cli.Params.SecretToken,
		"x-api-key":     cli.Params.ApiKey,
		"x-signature":   sig,
	}
}

// getHeadersEmpty returns headers for endpoints that don't use orderId/amount in the signature.
func (cli *Client) getHeadersEmpty() map[string]string {
	sig := utils.SignV2Empty(cli.Params.SecretKey)
	return map[string]string{
		"Content-Type":  "application/json",
		"Authorization": "Bearer " + cli.Params.SecretToken,
		"x-api-key":     cli.Params.ApiKey,
		"x-signature":   sig,
	}
}
