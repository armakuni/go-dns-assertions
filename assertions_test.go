package dnsassertions_test

import (
	"github.com/armakuni/go-dns-assertions/fetcher"
	"testing"
)

func TestAssertHasARecordDoesNotErrorWhenAMatchingRecordExists(t *testing.T) {
	result, errorable := createExampleResponse(t, []fetcher.Record{
		createExampleARecord("4.3.2.1"),
		createExampleARecord("1.2.3.4"),
	})

	result.AssertHasARecord("1.2.3.4")

	errorable.AssertNoErrorRaised()
}

func TestAssertHasARecordRaisesAnErrorWhenNoARecordsExist(t *testing.T) {
	result, errorable := createExampleResponse(t, []fetcher.Record{})

	result.AssertHasARecord("1.2.3.4")

	errorable.AssertRaisedError("DNS assertion failed: no A records found")
}

func TestAssertHasARecordRaisesAnErrorWhenNoRecordWithMatchingIPAddressExists(t *testing.T) {
	result, errorable := createExampleResponse(t, []fetcher.Record{
		createExampleARecord("4.3.2.1"),
	})

	result.AssertHasARecord("1.2.3.4")

	errorable.AssertRaisedError(
		"DNS asserting failed: No A record with value 1.2.3.4 found for example.com..\n" +
			"Records Found:\n" +
			"\tA\t4.3.2.1\n",
	)
}

func TestAssertHasCNAMERecordDoesNotErrorWhenAMatchingRecordExists(t *testing.T) {
	result, errorable := createExampleResponse(t, []fetcher.Record{
		createExampleCNAMERecord("target.example.com."),
	})

	result.AssertHasCNAMERecord("target.example.com.")

	errorable.AssertNoErrorRaised()
}

func TestAssertHasCNAMERecordRaisesAnErrorWhenNoARecordsExist(t *testing.T) {
	result, errorable := createExampleResponse(t, []fetcher.Record{})

	result.AssertHasCNAMERecord("target.example.com.")

	errorable.AssertRaisedError("DNS assertion failed: no CNAME records found")
}

func TestAssertHasCNAMERecordRaisesAnErrorWhenNoRecordWithMatchingIPAddressExists(t *testing.T) {
	result, errorable := createExampleResponse(t, []fetcher.Record{
		createExampleCNAMERecord("target1.example.com."),
	})

	result.AssertHasCNAMERecord("target2.example.com.")

	errorable.AssertRaisedError(
		"DNS asserting failed: No CNAME record with value target2.example.com. found for example.com..\n" +
			"Records Found:\n" +
			"\tCNAME\ttarget1.example.com.\n",
	)
}
