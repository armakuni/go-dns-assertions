package dnsassertions_test

import (
	"fmt"
	"github.com/armakuni/go-dns-assertions"
	"github.com/armakuni/go-dns-assertions/dnsclient"
	"github.com/stretchr/testify/assert"
	"testing"
)

type testDnsClient struct {
	willReturn       *dnsclient.Result
	fetchedFqdn      *string
	fetchedDnsServer *string
}

func (t *testDnsClient) LookupAllRecords(fqdn string, dnsServer string) (*dnsclient.Result, error) {
	if t.willReturn == nil {
		return nil, fmt.Errorf("no return value mocked")
	}

	t.fetchedFqdn = &fqdn
	t.fetchedDnsServer = &dnsServer

	return t.willReturn, nil
}

func TestFetchDNSRecordsTriggersErrorWhenClientErrors(t *testing.T) {
	errorTrigger := testErrorTrigger{t: t}
	dnsClient := &testDnsClient{willReturn: nil}
	testClient := dnsassertions.TestClient{Client: dnsClient, ErrorTrigger: &errorTrigger}

	testClient.FetchDNSRecords("example.com", "8.8.8.8")

	errorTrigger.AssertRaisedError("no return value mocked")
}

func TestFetchDNSRecordsDoesNotRaiseAnErrorWhenLookupIsSuccessful(t *testing.T) {
	errorTrigger := testErrorTrigger{t: t}
	dnsClient := &testDnsClient{willReturn: &dnsclient.Result{}}
	testClient := dnsassertions.TestClient{Client: dnsClient, ErrorTrigger: &errorTrigger}

	testClient.FetchDNSRecords("example.com", "8.8.8.8")

	errorTrigger.AssertNoErrorRaised()
}

func TestFetchDNSRecordsReturnsTheResultWhenSuccessful(t *testing.T) {
	errorTrigger := testErrorTrigger{t: t}
	expected := dnsclient.Result{FQDN: "returned-example.com"}
	dnsClient := &testDnsClient{willReturn: &expected}
	testClient := dnsassertions.TestClient{Client: dnsClient, ErrorTrigger: &errorTrigger}

	actual := testClient.FetchDNSRecords("example.com", "8.8.8.8")

	assert.Equal(t, "example.com", *dnsClient.fetchedFqdn)
	assert.Equal(t, "8.8.8.8", *dnsClient.fetchedDnsServer)
	assert.Equal(t,
		dnsassertions.ResultWithErrorTrigger{Result: &expected, ErrorTrigger: &errorTrigger},
		*actual,
	)
}
