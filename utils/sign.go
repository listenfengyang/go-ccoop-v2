package utils

import (
	"encoding/base64"
	"fmt"
)

// SignV2 generates the x-signature for the v2 API.
// Format: base64("{SECRET_KEY}:{orderId}:{amount}")
// For endpoints that don't use orderId and amount, pass empty strings.
func SignV2(secretKey string, orderId string, amount string) string {
	raw := fmt.Sprintf("%s:%s:%s", secretKey, orderId, amount)
	return base64.StdEncoding.EncodeToString([]byte(raw))
}

// SignV2Empty generates the x-signature for endpoints that don't use orderId/amount.
// Format: base64("{SECRET_KEY}:"":")
func SignV2Empty(secretKey string) string {
	return SignV2(secretKey, "", "")
}
