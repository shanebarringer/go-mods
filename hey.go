package hey

import (
	"rsc.io/quote/v3"
)

func Hi() string {
	return quote.HelloV3()
}

func Proverb() string {
	return quote.Concurrency()
}
