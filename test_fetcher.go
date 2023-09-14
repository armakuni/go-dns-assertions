// Package dnsassertions provides some tools to perform assertions on DNS lookups.
// You can use this for writing infrastructure tests.
package dnsassertions

import (
	"github.com/armakuni/go-dns-assertions/fetcher"
	"github.com/armakuni/go-dns-assertions/miekg_dns_fetcher"
)

// NewTestFetcher creates an instance of TestFetcher using the default Fetcher implementation.
// The current default implementation uses the github.com/miekg/dns library.
func NewTestFetcher(errorable Errorable) TestFetcher {
	dnsFetcher := miekgdnsfetcher.New()
	return TestFetcher{Fetcher: dnsFetcher, Errorable: errorable}
}

type TestFetcher struct {
	Fetcher   fetcher.Fetcher
	Errorable Errorable
}

// FetchDNSRecords fetches all records of all types and returns a value which can have assertion methods called on them.
func (testFetcher TestFetcher) FetchDNSRecords(fqdn string, dnsServer string) *LookupResultWithErrorable {
	result, err := testFetcher.Fetcher.FetchDNSRecords(fqdn, dnsServer)

	if err != nil {
		testFetcher.Errorable.Errorf(err.Error())
		return nil
	}

	return &LookupResultWithErrorable{LookupResult: result, Errorable: testFetcher.Errorable}
}
