# go-ccoop-v2

Go SDK for the CCoop v2 Payment API.

**Documentation:** https://payment-doc.gitbook.io/payment-document-eng

## Features

- ✅ Create Deposit (PromptPay QR Code & Bank Transfer)
- ✅ Create Withdraw
- ✅ Get Balance
- ✅ Deposit Callback handling
- ✅ Withdraw Callback handling
- ✅ Automatic signature generation (Base64)
- ✅ Supports debug mode for HTTP logging

## Authentication

The v2 API uses three headers for authentication:

| Header          | Description                                    |
|-----------------|------------------------------------------------|
| `Authorization` | `Bearer {YOUR_SECRET_TOKEN}`                   |
| `x-api-key`     | Your API key                                   |
| `x-signature`   | Base64-encoded `{SECRET_KEY}:{order_id}:{amount}` |

For endpoints that don't use `order_id` and `amount` (e.g., Get Balance), the signature is:
`Base64("{SECRET_KEY}::")`

## Installation

```bash
go get github.com/listenfengyang/go-ccoop-v2
```

## Quick Start

### 1. Create a config file (excluded from git)

Copy `config.go.example` to `config.go` and fill in your credentials:

```go
// config.go  (this file is in .gitignore)
package go_ccoop_v2

const (
    SECRET_TOKEN = "your_secret_token"
    API_KEY      = "your_api_key"
    SECRET_KEY   = "your_secret_key"
    BASE_URL     = "https://your-api-domain.com"

    DEPOSIT_CALLBACK_URL  = "https://your-domain.com/callback/deposit"
    WITHDRAW_CALLBACK_URL = "https://your-domain.com/callback/withdraw"
)
```

### 2. Initialize the client

```go
import ccoopv2 "github.com/listenfengyang/go-ccoop-v2"

client := ccoopv2.NewClient(logger, &ccoopv2.CCoopV2InitParams{
    SecretToken:         "YOUR_SECRET_TOKEN",
    ApiKey:              "YOUR_API_KEY",
    SecretKey:           "YOUR_SECRET_KEY",
    BaseUrl:             "https://your-api-domain.com",
    DepositCallbackUrl:  "https://your-domain.com/callback/deposit",
    WithdrawCallbackUrl: "https://your-domain.com/callback/withdraw",
})
```

### 3. Create a Deposit

```go
resp, err := client.Deposit(ccoopv2.CCoopV2DepositRequest{
    OrderId:     "ORDER-001",
    Amount:      1000.00,
    RefAccount:  "1234567890",
    RefBankCode: "014",           // SCB
    RefNameTh:   "จอห์น โด",
    RefNameEn:   "John Doe",
    RefUserId:   "user123",
})
if err != nil {
    log.Fatal(err)
}

if resp.DepositChannel == ccoopv2.DepositChannelPromptPay {
    // Show QR code to user
    fmt.Println("QR:", resp.QrImageLink)
} else {
    // Show bank transfer details
    fmt.Println("Bank Acc:", resp.DestBankAccNo)
    fmt.Println("Bank Name:", resp.DestBankAccName)
}
```

### 4. Create a Withdrawal

```go
resp, err := client.Withdraw(ccoopv2.CCoopV2WithdrawRequest{
    Amount:          500.00,
    DestBankAccNo:   "9876543210",
    DestBankAccName: "John Doe",
    DestBankCode:    "004",      // KBANK
    OrderId:         "WD-001",
})
```

### 5. Get Balance

```go
balance, err := client.GetBalance()
fmt.Printf("Deposit balance:  %.2f\n", balance.CurrentDepositBalance)
fmt.Printf("Withdraw balance: %.2f\n", balance.CurrentWithdrawBalance)
```

### 6. Handle Deposit Callback

```go
// In your HTTP handler:
var req ccoopv2.CCoopV2DepositCallbackReq
json.Unmarshal(body, &req)

err := client.DepositCallback(req, func(cb ccoopv2.CCoopV2DepositCallbackReq) error {
    switch cb.Status {
    case ccoopv2.DepositStatusAutoSuccess, ccoopv2.DepositStatusSuccess:
        // Mark order as paid
    case ccoopv2.DepositStatusFailed:
        // Mark order as failed
    }
    return nil
})
```

### 7. Handle Withdraw Callback

```go
var req ccoopv2.CCoopV2WithdrawCallbackReq
json.Unmarshal(body, &req)

err := client.WithdrawCallback(req, func(cb ccoopv2.CCoopV2WithdrawCallbackReq) error {
    switch cb.Status {
    case ccoopv2.WithdrawStatusAutoSuccess:
        // Mark withdrawal as completed
    case ccoopv2.WithdrawStatusFailed:
        // Mark withdrawal as failed
    }
    return nil
})
```

## Deposit Channel

| Channel         | QR Code | Bank Account Info |
|-----------------|---------|-------------------|
| `PROMPTPAY`     | ✅       | ✅                |
| `BANK_TRANSFER` | ❌       | ✅                |

## Bank Codes

Use the built-in `BankCodeEnum`:

```go
ccoopv2.BankCodeEnum.SCB.Code    // "014"
ccoopv2.BankCodeEnum.KBANK.Code  // "004"
ccoopv2.BankCodeEnum.KTB.Code    // "006"

// Or lookup by acronym:
ccoopv2.GetBankCodeByAcronym("SCB") // "014"
```

## Callback Status Values

### Deposit
| Status         | Description                    |
|----------------|--------------------------------|
| `PROCESSING`   | Being processed                |
| `AUTO_SUCCESS` | Deposit matched automatically  |
| `SUCCESS`      | Deposit confirmed manually     |
| `FAILED`       | Deposit failed                 |

### Withdraw
| Status            | Description                               |
|-------------------|-------------------------------------------|
| `PROCESSING`      | Being processed                           |
| `AUTO_SUCCESS`    | Withdrawal completed                      |
| `PENDING_CONFIRM` | Pending PIN confirmation                  |
| `FAILED`          | Withdrawal failed                         |

## Running Tests

```bash
cd go-ccoop-v2
# Make sure config.go has your real credentials
go test -v -run TestDeposit
go test -v -run TestWithdraw
go test -v -run TestGetBalance
go test -v -run TestDepositCallback
go test -v -run TestWithdrawCallback
```

## ⚠️ Important Notes

- `config.go` is excluded from git via `.gitignore` — never commit your credentials.
- Do **not** set expiration times on withdrawal requests.
- Do **not** retry withdrawal requests.
- To re-check a withdrawal status, use the Resend Callback endpoint.
- If you manually set a deposit to SUCCESS, drop any subsequent system callbacks for that order.
