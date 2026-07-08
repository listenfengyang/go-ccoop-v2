package go_ccoop_v2

import (
	"fmt"
	"testing"
)

// VLog is a simple logger implementation for testing
type VLog struct{}

func (l VLog) Debugf(format string, args ...interface{}) {
	fmt.Printf("[DEBUG] "+format+"\n", args...)
}
func (l VLog) Infof(format string, args ...interface{}) {
	fmt.Printf("[INFO]  "+format+"\n", args...)
}
func (l VLog) Warnf(format string, args ...interface{}) {
	fmt.Printf("[WARN]  "+format+"\n", args...)
}
func (l VLog) Errorf(format string, args ...interface{}) {
	fmt.Printf("[ERROR] "+format+"\n", args...)
}

// newTestClient creates a test client using credentials from config.go (gitignored)
func newTestClient() *Client {
	return NewClient(VLog{}, &CCoopV2InitParams{
		SecretToken:         SECRET_TOKEN,
		ApiKey:              API_KEY,
		SecretKey:           SECRET_KEY,
		BaseUrl:             BASE_URL,
		DepositCallbackUrl:  DEPOSIT_CALLBACK_URL,
		WithdrawCallbackUrl: WITHDRAW_CALLBACK_URL,
	})
}

// TestDeposit tests the deposit API (create_deposit)
func TestDeposit(t *testing.T) {
	cli := newTestClient()
	cli.SetDebugMode(true)

	resp, err := cli.Deposit(genDepositRequestDemo())
	if err != nil {
		t.Errorf("Deposit failed: %s", err.Error())
		return
	}
	fmt.Printf("Deposit response: %+v\n", resp)
	fmt.Printf("Status: %s\n", resp.Status)
	fmt.Printf("Deposit channel: %s\n", resp.DepositChannel)
	if resp.DepositChannel == DepositChannelPromptPay {
		fmt.Printf("QR Image Link: %s\n", resp.QrImageLink)
	} else {
		fmt.Printf("Dest Bank Acc No: %s\n", resp.DestBankAccNo)
		fmt.Printf("Dest Bank Acc Name: %s\n", resp.DestBankAccName)
	}
}

// genDepositRequestDemo builds a sample deposit request for testing
func genDepositRequestDemo() CCoopV2DepositRequest {
	return CCoopV2DepositRequest{
		OrderId:     "TEST-DEP-001",        // unique order ID in your system
		Amount:      250.00,                // deposit amount in THB
		RefAccount:  "1234567890",          // customer bank account number
		RefBankCode: BankCodeEnum.SCB.Code, // "014" - SCB
		RefNameTh:   "จอห์น โด",            // customer Thai name
		RefNameEn:   "John Doe",            // customer English name
		RefUserId:   "user123",             // user ID in your system
		Ref1:        "ref-data-1",
		Ref2:        "ref-data-2",
	}
}

// TestDepositCallback tests the deposit callback processing
func TestDepositCallback(t *testing.T) {
	cli := newTestClient()

	// Simulate a callback payload from the payment provider
	mockCallback := CCoopV2DepositCallbackReq{
		Type:          "DEPOSIT",
		Status:        DepositStatusAutoSuccess,
		StatusDesc:    "Deposit completed successfully",
		Amount:        250.00,
		TxnRefOrderId: "TEST-DEP-001",
		TxnRefId:      "TXN-REF-ABC123",
		TxnRefUserId:  "user123",
	}

	err := cli.DepositCallback(mockCallback, func(req CCoopV2DepositCallbackReq) error {
		fmt.Printf("Processing deposit callback: order=%s status=%s amount=%.2f\n",
			req.TxnRefOrderId, req.Status, req.Amount)
		return nil
	})

	if err != nil {
		t.Errorf("DepositCallback failed: %s", err.Error())
	}
}
