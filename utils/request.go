package utils

import (
	"net/http"
	"time"

	"github.com/go-resty/resty/v2"
)

type RestyRequest struct {
	Method   string              `json:"method"`
	Url      string              `json:"url"`
	Headers  http.Header         `json:"headers"`
	Body     interface{}         `json:"body"`
	FormData map[string][]string `json:"formData"`
}

type RestyResponse struct {
	StatusCode int         `json:"status_code"`
	Status     string      `json:"status"`
	Headers    http.Header `json:"headers"`
	Body       string      `json:"body"`
	ReceivedAt time.Time   `json:"received_at"`
	Rtt        int64       `json:"rtt"`
}

type RestyLog struct {
	Request  RestyRequest  `json:"request"`
	Response RestyResponse `json:"response"`
}

// GetRestyLog extracts request/response details for logging
func GetRestyLog(resp *resty.Response) RestyLog {
	reqHeaders := map[string][]string(resp.Request.Header)
	// Remove sensitive headers from logs
	delete(reqHeaders, "User-Agent")
	delete(reqHeaders, "Authorization")

	return RestyLog{
		Request: RestyRequest{
			Method:   resp.Request.Method,
			Url:      resp.Request.URL,
			Headers:  reqHeaders,
			Body:     resp.Request.Body,
			FormData: resp.Request.FormData,
		},
		Response: RestyResponse{
			StatusCode: resp.StatusCode(),
			Status:     resp.Status(),
			Headers:    resp.Header(),
			Body:       resp.String(),
			ReceivedAt: resp.ReceivedAt(),
			Rtt:        resp.Time().Milliseconds(),
		},
	}
}
