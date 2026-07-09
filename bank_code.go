package go_ccoop_v2

// THBBankCode represents a THB bank mapping entry between the client platform
// bank code (acronym) and the ccoop numeric bank code.
type THBBankCode struct {
	ClientCode string // Client platform bank code (e.g. "SCB", "KBANK")
	CcoopCode  string // CCoop numeric bank code (e.g. "014", "004")
	Name       string // Bank full name
}

// THBBankCodes is the mapping table between the client platform's THB bank codes
// and the CCoop v2 numeric bank codes used in deposit/withdraw requests.
//
// Usage:
//
//	ccoopCode := GetCCoopBankCode("SCB")  // returns "014"
//	clientCode := GetClientBankCode("014") // returns "SCB"
var THBBankCodes = []THBBankCode{
	{ClientCode: "BBL", CcoopCode: "002", Name: "Bangkok Bank"},
	{ClientCode: "KBANK", CcoopCode: "004", Name: "KASIKORNBANK PUBLIC COMPANY LIMITED"},
	{ClientCode: "KKR", CcoopCode: "004", Name: "KasiKorn Bank"}, // alias for KBANK
	{ClientCode: "KTB", CcoopCode: "006", Name: "KTB Net Bank"},
	{ClientCode: "TTB", CcoopCode: "011", Name: "TMBTHANACHART BANK PUBLIC COMPANY LIMITED"},
	{ClientCode: "TMB", CcoopCode: "011", Name: "TMBThananachart Bank (TTB)"}, // alias for TTB
	{ClientCode: "SCB", CcoopCode: "014", Name: "Siam Commercial Bank"},
	{ClientCode: "CIMBT", CcoopCode: "022", Name: "CIMB Thai"},
	{ClientCode: "UOBT", CcoopCode: "024", Name: "UNITED OVERSEAS BANK(THAl)PUBLIC COMPANY LIMITED"},
	{ClientCode: "BAY", CcoopCode: "025", Name: "BANK OF AYUDHYA PUBLIC COMPANY LTD"},
	{ClientCode: "BOA", CcoopCode: "025", Name: "Bank Of Ayudhya"}, // alias for BAY
	{ClientCode: "GSB", CcoopCode: "030", Name: "GOVERNMENT SAVINGS BANK"},
	{ClientCode: "GHB", CcoopCode: "033", Name: "THE GOVERNMENT HOUSING BANK"},
	{ClientCode: "BAAC", CcoopCode: "034", Name: "BANK FOR AGRICULTURE AND AGRICULTURAL COOPERATIVES"},
	{ClientCode: "EXIM", CcoopCode: "035", Name: "EXPORT-IMPORT BANK OF THAILAND"},
	{ClientCode: "ISBT", CcoopCode: "066", Name: "ISLAMIC BANK OF THAILAND"},
	{ClientCode: "TISCO", CcoopCode: "067", Name: "TISCO BANK PUBLIC COMPANY LIMITED"},
	{ClientCode: "ICBCT", CcoopCode: "070", Name: "INDUSTRIAL AND COMMERCIAL BANK OF CHINA (THAl)PUBLIC COMPANY LIMITED"},
	{ClientCode: "TCD", CcoopCode: "071", Name: "THE THAI CREDIT RETAIL BANK PUBLIC COMPANY LIMITED"},
	{ClientCode: "LHFG", CcoopCode: "073", Name: "LAND AND HOUSES BANK PUBLIC COMPANY LIMITED"},
	{ClientCode: "KKP", CcoopCode: "069", Name: "KIATNAKIN PHATRA BANK PUBLIC COMPANY LIMITED"},
	{ClientCode: "KNK", CcoopCode: "069", Name: "Kiatnakin Bank"}, // alias for KKP
	{ClientCode: "SME", CcoopCode: "098", Name: "SMALL ANDMEDIUM ENTERPRISE DEVELOPMENT BANK OF THAILAND"},
	// PPTP (Promptpay) maps to ccoop PromptPay deposit channel, not a numeric bank code
	{ClientCode: "PPTP", CcoopCode: "", Name: "Promptpay (use DepositChannelPromptPay)"},
}

// clientToCcoop maps client platform bank code → ccoop numeric bank code
var clientToCcoop map[string]string

// ccoopToClient maps ccoop numeric bank code → client platform bank code (primary only)
var ccoopToClient map[string]string

func init() {
	clientToCcoop = make(map[string]string, len(THBBankCodes))
	ccoopToClient = make(map[string]string)
	for _, b := range THBBankCodes {
		if b.CcoopCode != "" {
			clientToCcoop[b.ClientCode] = b.CcoopCode
			// Only register the first mapping to avoid alias overwrite
			if _, exists := ccoopToClient[b.CcoopCode]; !exists {
				ccoopToClient[b.CcoopCode] = b.ClientCode
			}
		}
	}
}

// GetCCoopBankCode converts a client platform bank code (e.g. "SCB") to
// the CCoop numeric bank code (e.g. "014").
// Returns empty string if not found.
func GetCCoopBankCode(clientCode string) string {
	return clientToCcoop[clientCode]
}

// GetClientBankCode converts a CCoop numeric bank code (e.g. "014") to
// the client platform bank code (e.g. "SCB").
// Returns empty string if not found.
func GetClientBankCode(ccoopCode string) string {
	return ccoopToClient[ccoopCode]
}
