package go_ccoop_v2

import (
	"fmt"
	"testing"
)

// TestWithdraw tests the withdraw API (create_withdraw)
func TestWithdraw(t *testing.T) {
	cli := newTestClient()
	cli.SetDebugMode(true)

	resp, err := cli.Withdraw(genWithdrawRequestDemo())
	if err != nil {
		t.Errorf("Withdraw failed: %s", err.Error())
		return
	}
	fmt.Printf("Withdraw response: %+v\n", resp)
	fmt.Printf("Status: %s\n", resp.Status)
	fmt.Printf("Ref ID: %s\n", resp.RefId)
	fmt.Printf("Amount: %.2f\n", resp.Amount)
}

// genWithdrawRequestDemo builds a sample withdraw request for testing
func genWithdrawRequestDemo() CCoopV2WithdrawRequest {
	return CCoopV2WithdrawRequest{
		Amount:          500.00,                  // withdrawal amount in THB
		DestBankAccNo:   "9876543210",            // destination bank account number
		DestBankAccName: "John Doe",              // destination account holder name
		DestBankCode:    BankCodeEnum.KBANK.Code, // "004" - KBANK
		OrderId:         "TEST-WD-003",           // merchant order reference (optional)
		// WithdrawCode: "635124",                 // PIN for large withdrawals (if required)
	}
}

// TestWithdrawCallback tests the withdraw callback processing
func TestWithdrawCallback(t *testing.T) {
	cli := newTestClient()

	// Simulate a callback payload from the payment provider
	mockCallback := CCoopV2WithdrawCallbackReq{
		Type:          "WITHDRAW",
		Status:        WithdrawStatusAutoSuccess,
		StatusDesc:    "Withdrawal completed successfully",
		Amount:        500.00,
		TxnRefOrderId: "TEST-WD-001",
		TxnRefId:      "TXN-WD-XYZ789",
	}

	err := cli.WithdrawCallback(mockCallback, func(req CCoopV2WithdrawCallbackReq) error {
		fmt.Printf("Processing withdraw callback: order=%s status=%s amount=%.2f\n",
			req.TxnRefOrderId, req.Status, req.Amount)
		return nil
	})

	if err != nil {
		t.Errorf("WithdrawCallback failed: %s", err.Error())
	}
}
