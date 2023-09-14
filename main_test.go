package dnsassertions_test

import (
	"github.com/armakuni/go-dns-assertions"
	"testing"
)

func TestFetchDNSRecords(t *testing.T) {
	fetcher := dnsassertions.NewTestFetcher(t)

	lookup := fetcher.FetchDNSRecords("armakuni.com", "8.8.8.8")
	lookup.AssertHasARecord("198.49.23.144")

	lookupWww := fetcher.FetchDNSRecords("www.armakuni.com", "8.8.8.8")
	lookupWww.AssertHasCNAMERecord("ext-cust.squarespace.com.")
}
