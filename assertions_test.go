package dnsassertions_test

import (
	"github.com/armakuni/go-dns-assertions/dnsclient"
	"testing"
)

func TestAssertHasARecordDoesNotErrorWhenAMatchingRecordExists(t *testing.T) {
	result, errorTrigger := createExampleResponse(t, []dnsclient.Record{
		createExampleARecord("4.3.2.1"),
		createExampleARecord("1.2.3.4"),
	})

	result.AssertHasARecord("1.2.3.4")

	errorTrigger.AssertNoErrorRaised()
}

func TestAssertHasARecordRaisesAnErrorWhenNoARecordsExist(t *testing.T) {
	result, errorTrigger := createExampleResponse(t, []dnsclient.Record{})

	result.AssertHasARecord("1.2.3.4")

	errorTrigger.AssertRaisedError("DNS assertion failed: no A records found")
}

func TestAssertHasARecordRaisesAnErrorWhenNoRecordWithMatchingIPAddressExists(t *testing.T) {
	result, errorTrigger := createExampleResponse(t, []dnsclient.Record{
		createExampleARecord("4.3.2.1"),
	})

	result.AssertHasARecord("1.2.3.4")

	errorTrigger.AssertRaisedError(
		"DNS asserting failed: No A record with value 1.2.3.4 found for example.com..\n" +
			"Records Found:\n" +
			"\tA\t4.3.2.1\n",
	)
}

func TestAssertHasCNAMERecordDoesNotErrorWhenAMatchingRecordExists(t *testing.T) {
	result, errorTrigger := createExampleResponse(t, []dnsclient.Record{
		createExampleCNAMERecord("target.example.com."),
	})

	result.AssertHasCNAMERecord("target.example.com.")

	errorTrigger.AssertNoErrorRaised()
}

func TestAssertHasCNAMERecordRaisesAnErrorWhenNoARecordsExist(t *testing.T) {
	result, errorTrigger := createExampleResponse(t, []dnsclient.Record{})

	result.AssertHasCNAMERecord("target.example.com.")

	errorTrigger.AssertRaisedError("DNS assertion failed: no CNAME records found")
}

func TestAssertHasCNAMERecordRaisesAnErrorWhenNoRecordWithMatchingIPAddressExists(t *testing.T) {
	result, errorTrigger := createExampleResponse(t, []dnsclient.Record{
		createExampleCNAMERecord("target1.example.com."),
	})

	result.AssertHasCNAMERecord("target2.example.com.")

	errorTrigger.AssertRaisedError(
		"DNS asserting failed: No CNAME record with value target2.example.com. found for example.com..\n" +
			"Records Found:\n" +
			"\tCNAME\ttarget1.example.com.\n",
	)
}

func TestAssertHasTXTRecordDoesNotErrorWhenAMatchingRecordExists(t *testing.T) {
	result, errorTrigger := createExampleResponse(t, []dnsclient.Record{
		createExampleTXTRecord("target.example.com."),
	})

	result.AssertHasTXTRecord("target.example.com.")

	errorTrigger.AssertNoErrorRaised()
}

func TestAssertHasTXTRecordRaisesAnErrorWhenNoARecordsExist(t *testing.T) {
	result, errorTrigger := createExampleResponse(t, []dnsclient.Record{})

	result.AssertHasTXTRecord("target.example.com.")

	errorTrigger.AssertRaisedError("DNS assertion failed: no TXT records found")
}

func TestAssertHasTXTRecordRaisesAnErrorWhenNoRecordWithMatchingIPAddressExists(t *testing.T) {
	result, errorTrigger := createExampleResponse(t, []dnsclient.Record{
		createExampleTXTRecord("text-value-one"),
	})

	result.AssertHasTXTRecord("text-value-two")

	errorTrigger.AssertRaisedError(
		"DNS asserting failed: No TXT record with value text-value-two found for example.com..\n" +
			"Records Found:\n" +
			"\tTXT\ttext-value-one\n",
	)
}
