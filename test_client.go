// Package dnsassertions provides some tools to perform assertions on DNS lookups.
// You can use this for writing infrastructure tests.
package dnsassertions

import (
	"github.com/armakuni/go-dns-assertions/dnsclient"
	"github.com/armakuni/go-dns-assertions/miekgdnsclient"
)

// NewTestClient creates an instance of TestClient using the default DNSClient implementation.
// The current default implementation uses the github.com/miekg/dns library.
func NewTestClient(errorTrigger ErrorTrigger) TestClient {
	dnsClient := miekgdnsclient.New()
	return TestClient{Client: dnsClient, ErrorTrigger: errorTrigger}
}

type TestClient struct {
	Client       dnsclient.DNSClient
	ErrorTrigger ErrorTrigger
}

// FetchDNSRecords fetches all records of all types and returns a value which can have assertion methods called on them.
func (client TestClient) FetchDNSRecords(fqdn string, dnsServer string) *ResultWithErrorTrigger {
	result, err := client.Client.LookupAllRecords(fqdn, dnsServer)

	if err != nil {
		client.ErrorTrigger.Errorf(err.Error())
		return nil
	}

	return &ResultWithErrorTrigger{Result: result, ErrorTrigger: client.ErrorTrigger}
}
