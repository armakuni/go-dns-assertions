package dnsassertions

import "github.com/armakuni/go-dns-assertions/fetcher"

type Errorable interface {
	Errorf(format string, args ...any)
}

type LookupResultWithErrorable struct {
	*fetcher.LookupResult
	Errorable Errorable
}
