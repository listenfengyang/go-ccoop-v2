package go_ccoop_v2

import (
	"fmt"
	"testing"
)

// TestGetBalance tests the get balance API
func TestGetBalance(t *testing.T) {
	cli := newTestClient()
	cli.SetDebugMode(true)

	resp, err := cli.GetBalance()
	if err != nil {
		t.Errorf("GetBalance failed: %s", err.Error())
		return
	}
	fmt.Printf("Balance response: %+v\n", resp)
	fmt.Printf("Deposit balance:  %.2f\n", resp.CurrentDepositBalance)
	fmt.Printf("Withdraw balance: %.2f\n", resp.CurrentWithdrawBalance)
}
