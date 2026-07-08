package go_ccoop_v2

// CCoopV2InitParams holds the configuration for the v2 SDK client.
type CCoopV2InitParams struct {
	SecretToken string `json:"secretToken" mapstructure:"secretToken" yaml:"secretToken"` // Bearer secret token (Authorization header)
	ApiKey      string `json:"apiKey" mapstructure:"apiKey" yaml:"apiKey"`                // x-api-key header
	SecretKey   string `json:"secretKey" mapstructure:"secretKey" yaml:"secretKey"`       // used to generate x-signature

	BaseUrl string `json:"baseUrl" mapstructure:"baseUrl" yaml:"baseUrl"` // e.g. https://your-api-domain.com

	DepositCallbackUrl  string `json:"depositCallbackUrl" mapstructure:"depositCallbackUrl" yaml:"depositCallbackUrl"`    // default deposit callback URL
	WithdrawCallbackUrl string `json:"withdrawCallbackUrl" mapstructure:"withdrawCallbackUrl" yaml:"withdrawCallbackUrl"` // default withdraw callback URL
}

// ============================================================
// Deposit
// ============================================================

// CCoopV2DepositRequest is the request body for creating a deposit.
type CCoopV2DepositRequest struct {
	OrderId     string  `json:"order_id"`               // Unique order ID in your system (required)
	Amount      float64 `json:"amount"`                 // Deposit amount, e.g. 1000.00 (required)
	RefAccount  string  `json:"ref_account"`            // Customer account reference (required)
	RefBankCode string  `json:"ref_bank_code"`          // Depositor's bank code (required)
	RefNameTh   string  `json:"ref_name_th"`            // Customer Thai name (required)
	RefNameEn   string  `json:"ref_name_en,omitempty"`  // Customer English name (optional)
	RefUserId   string  `json:"ref_user_id"`            // Customer ID in your system (required)
	Ref1        string  `json:"ref1,omitempty"`         // Additional reference 1 (optional)
	Ref2        string  `json:"ref2,omitempty"`         // Additional reference 2 (optional)
	CallbackUrl string  `json:"callback_url,omitempty"` // Override callback URL (optional)
}

// CCoopV2DepositResponse is the response from the create deposit API.
type CCoopV2DepositResponse struct {
	Status          string  `json:"status"`             // e.g. "CREATED"
	RefId           string  `json:"ref_id"`             // Internal reference ID in the provider's system
	DestBankCode    string  `json:"dest_bank_code"`     // Destination bank code
	DestBankAccNo   string  `json:"dest_bank_acc_no"`   // Destination bank account number
	DestBankAccName string  `json:"dest_bank_acc_name"` // Destination bank account name
	OrderId         string  `json:"order_id"`           // Customer's deposit order ID
	Amount          float64 `json:"amount"`             // Deposit amount
	QrImage         string  `json:"qr_image"`           // Base64 QR code image (PromptPay only)
	QrImageLink     string  `json:"qr_image_link"`      // URL to QR code image (PromptPay only)
	CreateTime      string  `json:"create_time"`        // Timestamp when the transaction was created
	ExpireTime      string  `json:"expire_time"`        // Timestamp when the transaction expires
	DepositChannel  string  `json:"deposit_channel"`    // "PROMPTPAY" or "BANK_TRANSFER"
}

// DepositChannel constants
const (
	DepositChannelPromptPay    = "PROMPTPAY"
	DepositChannelBankTransfer = "BANK_TRANSFER"
)

// CCoopV2DepositCallbackReq is the payload received when a deposit callback is triggered.
type CCoopV2DepositCallbackReq struct {
	Type                string  `json:"type"`   // "DEPOSIT"
	Status              string  `json:"status"` // AUTO_SUCCESS, SUCCESS, FAILED, PROCESSING
	StatusDesc          string  `json:"status_desc"`
	Amount              float64 `json:"amount"`
	MdrAmount           float64 `json:"mdr_amount"`
	StmDesc             string  `json:"stm_desc"`
	StmRefId            string  `json:"stm_ref_id"`
	StmBankCode         string  `json:"stm_bank_code"`
	StmBankAccNo        string  `json:"stm_bank_acc_no"`
	StmBankAccName      string  `json:"stm_bank_acc_name"`
	StmTimestamp        string  `json:"stm_timestamp"`
	TxnRefId            string  `json:"txn_ref_id"`
	TxnRefUserId        string  `json:"txn_ref_user_id"`
	TxnRef1             string  `json:"txn_ref1"`
	TxnRef2             string  `json:"txn_ref2"`
	TxnRefOrderId       string  `json:"txn_ref_order_id"`
	TxnRefBankCode      string  `json:"txn_ref_bank_code"`
	TxnRefBankAccNo     string  `json:"txn_ref_bank_acc_no"`
	TxnRefBankAccNameEn string  `json:"txn_ref_bank_acc_name_en"`
	TxnRefBankAccNameTh string  `json:"txn_ref_bank_acc_name_th"`
	TxnTimestamp        string  `json:"txn_timestamp"`
}

// Deposit callback status values
const (
	DepositStatusProcessing  = "PROCESSING"
	DepositStatusAutoSuccess = "AUTO_SUCCESS"
	DepositStatusSuccess     = "SUCCESS"
	DepositStatusFailed      = "FAILED"
)

// ============================================================
// Withdraw
// ============================================================

// CCoopV2WithdrawRequest is the request body for creating a withdrawal.
type CCoopV2WithdrawRequest struct {
	Amount          float64 `json:"amount"`                  // Withdrawal amount (required)
	DestBankAccNo   string  `json:"dest_bank_acc_no"`        // Destination bank account number (required)
	DestBankAccName string  `json:"dest_bank_acc_name"`      // Destination account holder name (required)
	DestBankCode    string  `json:"dest_bank_code"`          // Destination bank code (required)
	WithdrawCode    string  `json:"withdraw_code,omitempty"` // PIN for large withdrawals (optional)
	CallbackUrl     string  `json:"callback_url,omitempty"`  // Override callback URL (optional)
	OrderId         string  `json:"order_id,omitempty"`      // Merchant order reference (optional)
}

// CCoopV2WithdrawResponse is the response from the create withdraw API.
type CCoopV2WithdrawResponse struct {
	Status          string  `json:"status"`             // e.g. "PROCESSING"
	RefId           string  `json:"ref_id"`             // Internal reference ID
	DestBankCode    string  `json:"dest_bank_code"`     // Destination bank code
	DestBankAccNo   string  `json:"dest_bank_acc_no"`   // Destination bank account number
	DestBankAccName string  `json:"dest_bank_acc_name"` // Destination account name
	Amount          float64 `json:"amount"`             // Withdrawal amount
	WithdrawTime    string  `json:"withdraw_time"`      // Timestamp
}

// CCoopV2WithdrawCallbackReq is the payload received when a withdrawal callback is triggered.
type CCoopV2WithdrawCallbackReq struct {
	Type                string  `json:"type"`   // "WITHDRAW"
	Status              string  `json:"status"` // AUTO_SUCCESS, FAILED, PROCESSING, PENDING_CONFIRM
	StatusDesc          string  `json:"status_desc"`
	Amount              float64 `json:"amount"`
	MdrAmount           float64 `json:"mdr_amount"`
	StmDesc             string  `json:"stm_desc"`
	StmRefId            string  `json:"stm_ref_id"`
	StmBankCode         string  `json:"stm_bank_code"`
	StmBankAccNo        string  `json:"stm_bank_acc_no"`
	StmBankAccName      string  `json:"stm_bank_acc_name"`
	StmTimestamp        string  `json:"stm_timestamp"`
	TxnRefId            string  `json:"txn_ref_id"`
	TxnRefUserId        string  `json:"txn_ref_user_id"`
	TxnRef1             string  `json:"txn_ref1"`
	TxnRef2             string  `json:"txn_ref2"`
	TxnRefOrderId       string  `json:"txn_ref_order_id"`
	TxnRefBankCode      string  `json:"txn_ref_bank_code"`
	TxnRefBankAccNo     string  `json:"txn_ref_bank_acc_no"`
	TxnRefBankAccNameEn string  `json:"txn_ref_bank_acc_name_en"`
	TxnRefBankAccNameTh string  `json:"txn_ref_bank_acc_name_th"`
	TxnTimestamp        string  `json:"txn_timestamp"`
}

// Withdraw callback status values
const (
	WithdrawStatusProcessing     = "PROCESSING"
	WithdrawStatusAutoSuccess    = "AUTO_SUCCESS"
	WithdrawStatusPendingConfirm = "PENDING_CONFIRM"
	WithdrawStatusFailed         = "FAILED"
)

// ============================================================
// Balance
// ============================================================

// CCoopV2BalanceResponse is the response from the get balance API.
type CCoopV2BalanceResponse struct {
	CurrentDepositBalance  float64 `json:"current_deposit_balance"`
	CurrentWithdrawBalance float64 `json:"current_withdraw_balance"`
}

// ============================================================
// Error responses
// ============================================================

// CCoopV2ErrorDetail is a validation error detail item.
type CCoopV2ErrorDetail struct {
	Loc  []interface{} `json:"loc"`
	Msg  string        `json:"msg"`
	Type string        `json:"type"`
}

// CCoopV2ValidationError is the response for HTTP 422 Validation Error.
type CCoopV2ValidationError struct {
	Detail []CCoopV2ErrorDetail `json:"detail"`
}

// CCoopV2NotFoundError is the response for HTTP 404 Not Found.
type CCoopV2NotFoundError struct {
	Detail string `json:"detail"`
}
