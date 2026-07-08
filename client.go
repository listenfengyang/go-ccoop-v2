package go_ccoop_v2

import (
	"context"
	"crypto/tls"
	"net"
	"net/http"
	"strings"
	"time"

	"github.com/go-resty/resty/v2"
	"github.com/listenfengyang/go-ccoop-v2/utils"
)

// Client is the v2 SDK client for the CCoop payment API.
type Client struct {
	Params *CCoopV2InitParams

	ryClient  *resty.Client
	debugMode bool
	logger    utils.Logger
}

// NewClient creates a new CCoop v2 SDK client.
func NewClient(logger utils.Logger, params *CCoopV2InitParams) *Client {
	// Custom dialer: resolve sandbox.goingpays.com to known Cloudflare IPs
	// since the subdomain has no public DNS record.
	customDialer := &net.Dialer{
		Timeout:   30 * time.Second,
		KeepAlive: 30 * time.Second,
	}
	transport := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		DialContext: func(ctx context.Context, network, addr string) (net.Conn, error) {
			host, port, err := net.SplitHostPort(addr)
			if err != nil {
				return nil, err
			}
			// If the host cannot be resolved via DNS, fall back to the known IP
			if strings.HasSuffix(host, "goingpays.com") {
				if _, err := net.DefaultResolver.LookupHost(ctx, host); err != nil {
					// Use the Cloudflare IP that goingpays.com resolves to
					addr = net.JoinHostPort("172.67.161.51", port)
				}
			}
			return customDialer.DialContext(ctx, network, addr)
		},
	}

	rc := resty.New()
	rc.SetTransport(transport)

	return &Client{
		Params:    params,
		ryClient:  rc,
		debugMode: false,
		logger:    logger,
	}
}

// SetDebugMode enables or disables debug mode for HTTP logging.
func (cli *Client) SetDebugMode(debugMode bool) {
	cli.debugMode = debugMode
}
