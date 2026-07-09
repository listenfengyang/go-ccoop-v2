package go_ccoop_v2

// API endpoint paths
const (
	EndpointCreateDeposit        = "/api/v1/client/create_deposit"
	EndpointCreateWithdraw       = "/api/v1/client/create_withdraw"
	EndpointGetBalance           = "/api/v1/client/get_balance"
	EndpointResendCallbackForTxn = "/api/v1/client/resend_callback_for_transaction"
)

// BankCode represents a Thai bank code entry
type BankCode struct {
	Code string
	Name string
	Desc string
}

// BankCodeEnum contains all supported Thai bank codes
var BankCodeEnum = struct {
	BBL   BankCode
	KBANK BankCode
	KTB   BankCode
	TTB   BankCode
	SCB   BankCode
	BAY   BankCode
	KKP   BankCode
	CIMBT BankCode
	TISCO BankCode
	UOBT  BankCode
	TCD   BankCode
	LHFG  BankCode
	ICBCT BankCode
	SME   BankCode
	BAAC  BankCode
	EXIM  BankCode
	GSB   BankCode
	GHB   BankCode
	ISBT  BankCode
}{
	BBL:   BankCode{"002", "BBL", "Bangkok Bank"},
	KBANK: BankCode{"004", "KBANK", "Kasikorn Bank"},
	KTB:   BankCode{"006", "KTB", "Krungthai Bank"},
	TTB:   BankCode{"011", "TTB", "TMBThanachart Bank"},
	SCB:   BankCode{"014", "SCB", "Siam Commercial Bank"},
	BAY:   BankCode{"025", "BAY", "Bank of Ayudhya"},
	KKP:   BankCode{"069", "KKP", "Kiatnakin Phatra Bank"},
	CIMBT: BankCode{"022", "CIMBT", "CIMB Thai Bank"},
	TISCO: BankCode{"067", "TISCO", "Tisco Bank"},
	UOBT:  BankCode{"024", "UOBT", "United Overseas Bank"},
	TCD:   BankCode{"071", "TCD", "Thai Credit Retail Bank"},
	LHFG:  BankCode{"073", "LHFG", "Land and Houses Bank"},
	ICBCT: BankCode{"070", "ICBCT", "ICBC (Thai)"},
	SME:   BankCode{"098", "SME", "SME Development Bank"},
	BAAC:  BankCode{"034", "BAAC", "Bank for Agriculture and Agricultural Cooperatives"},
	EXIM:  BankCode{"035", "EXIM", "Export-Import Bank of Thailand"},
	GSB:   BankCode{"030", "GSB", "Government Savings Bank"},
	GHB:   BankCode{"033", "GHB", "Government Housing Bank"},
	ISBT:  BankCode{"066", "ISBT", "Islamic Bank of Thailand"},
}

// GetBankCodeByAcronym returns the numeric bank code for the given acronym
func GetBankCodeByAcronym(acronym string) string {
	codes := map[string]string{
		"BBL":   "002",
		"KBANK": "004",
		"KTB":   "006",
		"TTB":   "011",
		"SCB":   "014",
		"BAY":   "025",
		"KKP":   "069",
		"CIMBT": "022",
		"TISCO": "067",
		"UOBT":  "024",
		"TCD":   "071",
		"LHFG":  "073",
		"ICBCT": "070",
		"SME":   "098",
		"BAAC":  "034",
		"EXIM":  "035",
		"GSB":   "030",
		"GHB":   "033",
		"ISBT":  "066",
	}
	return codes[acronym]
}

// Eq checks if the bank code matches the given value
func (b BankCode) Eq(value string) bool {
	return b.Code == value
}
