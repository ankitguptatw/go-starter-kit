package data

import "myapp/persistence/dao"

// TODO: use https://github.com/brianvoe/gofakeit

var Banks = []dao.Bank{
	{URL: "http://example.com/foo", Code: "FOO"},
	{URL: "http://example.com/bar", Code: "BAR"},
	{URL: "http://example.com/baz", Code: "BAZ"},
}
