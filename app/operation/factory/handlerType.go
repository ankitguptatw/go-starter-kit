package factory

type CommandHandlers int
type QueryHandlers int

const (
	AddBankCommandHandler CommandHandlers = iota
	AddPaymentCommandHandler
)
const (
	GetaAllBanksQueryHandler QueryHandlers = iota
	GetBankQueryHandler
	GetPaymentQueryHandler
)
