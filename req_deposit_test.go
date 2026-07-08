package go_ccoop_v2

import (
	"fmt"
	"testing"
	"time"
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

	// Sandbox environment does not return qr_image / qr_image_link.
	// Mock these values so the PromptPay branch can be exercised in tests.
	if resp.QrImage == "" {
		resp.QrImage = "iVBORw0KGgoAAAANSUhEUgAAAXIAAAFyAQAAAADAX2yk..."
	}
	if resp.QrImageLink == "" {
		resp.QrImageLink = "https://storage.googleapis.com/corp-richpay_richman_qr_data/006/CWUMQG-1775190305360-20260403112503-3719ad.png"
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

// genDepositRequestDemo builds a sample deposit request for testing.
// order_id and ref_account use a timestamp suffix to avoid duplicate/pending transaction errors.
func genDepositRequestDemo() CCoopV2DepositRequest {
	ts := time.Now().UnixMilli()
	return CCoopV2DepositRequest{
		OrderId:     fmt.Sprintf("TEST-DEP-%d", ts),              // unique order ID per run
		Amount:      1000.00,                                     // deposit amount in THB
		RefAccount:  fmt.Sprintf("%d", ts%9000000000+1000000000), // unique 10-digit account per run
		RefBankCode: BankCodeEnum.SCB.Code,                       // "014" - SCB
		RefNameTh:   "จอห์น โด",                                  // customer Thai name
		RefNameEn:   "John Doe",                                  // customer English name
		RefUserId:   fmt.Sprintf("user-%d", ts),                  // unique user ID per run
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
