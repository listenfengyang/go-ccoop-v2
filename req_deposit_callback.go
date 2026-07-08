package go_ccoop_v2

import "errors"

// DepositCallback handles an incoming deposit callback from the payment provider.
// It validates the callback type and passes it to the provided processor function.
//
// Deposit callback status values:
//   - PROCESSING    : Transaction is being processed
//   - AUTO_SUCCESS  : Deposit succeeded automatically (system matched)
//   - SUCCESS       : Deposit succeeded (manually confirmed)
//   - FAILED        : Deposit failed
//
// Note: If a transaction was manually set to SUCCESS by the merchant,
// ignore/drop any subsequent system callback for that transaction to avoid duplicate processing.
func (cli *Client) DepositCallback(req CCoopV2DepositCallbackReq, processor func(CCoopV2DepositCallbackReq) error) error {
	if req.Type != "DEPOSIT" {
		return errors.New("invalid callback type: expected DEPOSIT")
	}

	return processor(req)
}
