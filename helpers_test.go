package dnsassertions_test

import (
	"fmt"
	"github.com/armakuni/go-dns-assertions"
	"github.com/armakuni/go-dns-assertions/dnsclient"
	"github.com/stretchr/testify/assert"
	"testing"
)

type testErrorTrigger struct {
	t            *testing.T
	errorMessage *string
}

func (t *testErrorTrigger) AssertRaisedError(expected string) {
	assert.Equal(t.t, expected, *t.errorMessage)
}

func (t *testErrorTrigger) Errorf(format string, args ...any) {
	message := fmt.Sprintf(format, args...)
	t.errorMessage = &message
}

func (t *testErrorTrigger) AssertNoErrorRaised() {
	if t.errorMessage != nil {
		t.t.Errorf("Unexpected error message \"%s\" was raised", *t.errorMessage)
	}
}

func createExampleResponse(t *testing.T, records []dnsclient.Record) (*dnsassertions.ResultWithErrorTrigger, *testErrorTrigger) {
	errorTrigger := testErrorTrigger{t: t}

	result := dnsassertions.ResultWithErrorTrigger{
		Result: &dnsclient.Result{
			FQDN:    "example.com.",
			Records: records,
		},
		ErrorTrigger: &errorTrigger,
	}

	return &result, &errorTrigger
}

func createExampleARecord(ipv4Addr string) *dnsclient.A {
	return &dnsclient.A{Common: &dnsclient.Common{Raw: "A\t" + ipv4Addr}, Ipv4Addr: ipv4Addr}
}

func createExampleCNAMERecord(target string) *dnsclient.CNAME {
	return &dnsclient.CNAME{Common: &dnsclient.Common{Raw: "CNAME\t" + target}, Target: target}
}
