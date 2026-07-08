package go_ccoop_v2

import "errors"

// WithdrawCallback handles an incoming withdrawal callback from the payment provider.
// It validates the callback type and passes it to the provided processor function.
//
// Withdraw callback status values:
//   - PROCESSING      : Withdrawal is being processed
//   - AUTO_SUCCESS    : Withdrawal completed successfully
//   - PENDING_CONFIRM : Pending merchant confirmation (PIN required)
//   - FAILED          : Withdrawal failed
//
// Note: Callbacks for successful and unsuccessful transactions carry the same payload;
// only the status differs.
func (cli *Client) WithdrawCallback(req CCoopV2WithdrawCallbackReq, processor func(CCoopV2WithdrawCallbackReq) error) error {
	if req.Type != "WITHDRAW" {
		return errors.New("invalid callback type: expected WITHDRAW")
	}

	return processor(req)
}
