package go_ccoop_v2

import (
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
	return &Client{
		Params:    params,
		ryClient:  resty.New(),
		debugMode: false,
		logger:    logger,
	}
}

// SetDebugMode enables or disables debug mode for HTTP logging.
func (cli *Client) SetDebugMode(debugMode bool) {
	cli.debugMode = debugMode
}
