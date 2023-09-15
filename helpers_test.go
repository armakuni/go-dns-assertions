package dnsassertions_test

import (
	"fmt"
	"github.com/armakuni/go-dns-assertions"
	"github.com/armakuni/go-dns-assertions/fetcher"
	"github.com/stretchr/testify/assert"
	"testing"
)

type testErrorable struct {
	t            *testing.T
	errorMessage *string
}

func (t *testErrorable) AssertRaisedError(expected string) {
	assert.Equal(t.t, expected, *t.errorMessage)
}

func (t *testErrorable) Errorf(format string, args ...any) {
	message := fmt.Sprintf(format, args...)
	t.errorMessage = &message
}

func (t *testErrorable) AssertNoErrorRaised() {
	if t.errorMessage != nil {
		t.t.Errorf("Unexpected error message \"%s\" was raised", *t.errorMessage)
	}
}

func createExampleResponse(t *testing.T, records []fetcher.Record) (*dnsassertions.LookupResultWithErrorable, *testErrorable) {
	errorable := testErrorable{t: t}

	result := dnsassertions.LookupResultWithErrorable{
		LookupResult: &fetcher.LookupResult{
			FQDN:    "example.com.",
			Records: records,
		},
		Errorable: &errorable,
	}

	return &result, &errorable
}

func createExampleARecord(ipv4Addr string) *fetcher.A {
	return &fetcher.A{Base: &fetcher.Base{Raw: "A\t" + ipv4Addr}, Ipv4Addr: ipv4Addr}
}

func createExampleCNAMERecord(target string) *fetcher.CNAME {
	return &fetcher.CNAME{Base: &fetcher.Base{Raw: "CNAME\t" + target}, Target: target}
}
