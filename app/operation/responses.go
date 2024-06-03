package operation

type AddBankResponse struct {
	ID uint `json:"id"`
}
type BankResponse struct {
	ID   uint   `json:"id"`
	Code string `json:"code"`
	URL  string `json:"url"`
}
type BanksResponse struct {
	Banks []BankResponse `json:"banks"`
}

type AddPaymentResponse struct {
	ID uint `json:"id"`
}

type Account struct {
	Name          string `json:"name"`
	AccountNumber int64  `json:"accountNumber"`
	Code          string `json:"Code"`
}

type PaymentResponse struct {
	ID          uint    `json:"id"`
	Amount      float64 `json:"amount"`
	Beneficiary Account `json:"beneficiary"`
	Payee       Account `json:"payee"`
}
