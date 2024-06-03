package core

type Bank struct {
	//nolint:unused
	id   int64
	code string
	url  string
}

func (bank *Bank) URL() string {
	return bank.url
}

func (bank *Bank) Code() string {
	return bank.code
}
