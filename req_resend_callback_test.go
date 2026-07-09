package go_ccoop_v2

import (
	"fmt"
	"testing"
)

// TestResendCallback tests the resend callback API.
//
// Before running this test, ensure you have a valid ref_id from a real or sandbox
// transaction. You can obtain ref_id from the Deposit or Withdraw response (resp.RefId).
//
// Run with:
//
//	go test -v -run TestResendCallback
func TestResendCallback(t *testing.T) {
	cli := newTestClient()
	cli.SetDebugMode(true)

	// Replace with a real ref_id from your sandbox / production environment
	refId := "OZOAQQ-1783580526949-20260709140206-a518d2"

	resp, err := cli.ResendCallback(CCoopV2ResendCallbackRequest{
		RefId: refId,
	})
	if err != nil {
		t.Errorf("ResendCallback failed: %s", err.Error())
		return
	}

	fmt.Printf("ResendCallback response: %+v\n", resp)
	fmt.Printf("Detail: %s\n", resp.Detail)
}

// TestResendCallbackAfterDeposit creates a deposit first, then immediately
// triggers a callback resend using the ref_id from the deposit response.
//
// This is an end-to-end test that exercises both the Deposit and ResendCallback APIs.
//
// Run with:
//
//	go test -v -run TestResendCallbackAfterDeposit
func TestResendCallbackAfterDeposit(t *testing.T) {
	cli := newTestClient()
	cli.SetDebugMode(true)

	// Step 1: Create a deposit to get a valid ref_id
	t.Log("Step 1: Creating deposit to obtain ref_id...")
	depositResp, err := cli.Deposit(genDepositRequestDemo())
	if err != nil {
		t.Fatalf("Deposit failed: %s", err.Error())
	}
	fmt.Printf("Deposit created — RefId: %s, Status: %s\n", depositResp.RefId, depositResp.Status)

	if depositResp.RefId == "" {
		t.Fatal("Deposit response did not contain a ref_id; cannot proceed with resend callback test")
	}

	// Step 2: Resend the callback using the ref_id obtained above
	t.Logf("Step 2: Resending callback for ref_id=%s ...", depositResp.RefId)
	resendResp, err := cli.ResendCallback(CCoopV2ResendCallbackRequest{
		RefId: depositResp.RefId,
	})
	if err != nil {
		t.Errorf("ResendCallback failed: %s", err.Error())
		return
	}

	fmt.Printf("ResendCallback response: %+v\n", resendResp)
	fmt.Printf("Detail: %s\n", resendResp.Detail)
}
