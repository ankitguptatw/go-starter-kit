package query

type GetBankQuery struct {
	Code string `uri:"code" binding:"required"`
}

type GetPaymentQuery struct {
	ID uint `uri:"id" binding:"required"`
}
