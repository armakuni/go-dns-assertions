package dnsassertions_test

import (
	"fmt"
	dnsassertions "github.com/armakuni/go-dns-assertions"
	"github.com/armakuni/go-dns-assertions/fetcher"
	"github.com/stretchr/testify/assert"
	"testing"
)

type testDnsFetcher struct {
	willReturn       *fetcher.LookupResult
	fetchedFqdn      *string
	fetchedDnsServer *string
}

func (t *testDnsFetcher) FetchDNSRecords(fqdn string, dnsServer string) (*fetcher.LookupResult, error) {
	if t.willReturn == nil {
		return nil, fmt.Errorf("no return value mocked")
	}

	t.fetchedFqdn = &fqdn
	t.fetchedDnsServer = &dnsServer

	return t.willReturn, nil
}

func TestFetchDNSRecordsTriggersErrorWhenFetcherErrors(t *testing.T) {
	errorable := testErrorable{t: t}
	dnsFetcher := &testDnsFetcher{willReturn: nil}
	testFetcher := dnsassertions.TestFetcher{Fetcher: dnsFetcher, Errorable: &errorable}

	testFetcher.FetchDNSRecords("example.com", "8.8.8.8")

	errorable.AssertRaisedError("no return value mocked")
}

func TestFetchDNSRecordsDoesNotRaiseAnErrorWhenFetchingIsSuccessful(t *testing.T) {
	errorable := testErrorable{t: t}
	dnsFetcher := &testDnsFetcher{willReturn: &fetcher.LookupResult{}}
	testFetcher := dnsassertions.TestFetcher{Fetcher: dnsFetcher, Errorable: &errorable}

	testFetcher.FetchDNSRecords("example.com", "8.8.8.8")

	errorable.AssertNoErrorRaised()
}

func TestFetchDNSRecordsReturnsTheResultWhenSuccessful(t *testing.T) {
	errorable := testErrorable{t: t}
	expected := fetcher.LookupResult{FQDN: "returned-example.com"}
	dnsFetcher := &testDnsFetcher{willReturn: &expected}
	testFetcher := dnsassertions.TestFetcher{Fetcher: dnsFetcher, Errorable: &errorable}

	actual := testFetcher.FetchDNSRecords("example.com", "8.8.8.8")

	assert.Equal(t, "example.com", *dnsFetcher.fetchedFqdn)
	assert.Equal(t, "8.8.8.8", *dnsFetcher.fetchedDnsServer)
	assert.Equal(t,
		dnsassertions.LookupResultWithErrorable{LookupResult: &expected, Errorable: &errorable},
		*actual,
	)
}
