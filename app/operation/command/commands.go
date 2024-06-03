package command

type AddBankCommand struct {
	Code string `json:"code" binding:"required"`
	URL  string `json:"url" binding:"required"`
}

type AddPaymentCommand struct {
	Amount      float64        `json:"amount" binding:"required"`
	Beneficiary AccountDetails `json:"beneficiary" binding:"required"`
	Payee       AccountDetails `json:"payee" binding:"required"`
}

type AccountDetails struct {
	Name          string `json:"name" binding:"required"`
	AccountNumber int64  `json:"accountNumber" binding:"required"`
	Code          string `json:"code" binding:"required"`
}
