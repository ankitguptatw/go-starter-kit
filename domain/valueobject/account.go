package valueobject

type Account struct {
	name          string
	accountNumber int64
	code          string
}

func NewAccount(name string, accountNumber int64, code string) Account {
	return Account{name: name, accountNumber: accountNumber, code: code}
}
func (a Account) Name() string {
	return a.name
}
func (a Account) AccountNumber() int64 {
	return a.accountNumber
}
func (a Account) Code() string {
	return a.code
}
