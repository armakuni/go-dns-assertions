package dnsassertions_test

import (
	"github.com/armakuni/go-dns-assertions"
	"testing"
)

func TestFetchDNSRecords(t *testing.T) {
	client := dnsassertions.NewTestClient(t)

	result := client.FetchDNSRecords("armakuni.com", "8.8.8.8")
	result.AssertHasARecord("198.49.23.144")

	resultWww := client.FetchDNSRecords("www.armakuni.com", "8.8.8.8")
	resultWww.AssertHasCNAMERecord("ext-cust.squarespace.com.")
}
