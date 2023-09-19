package dnsassertions

import (
	"fmt"
	"github.com/armakuni/go-dns-assertions/dnsclient"
	"strings"
)

// AssertHasARecord fails the test when there are no A records with a matching IP address.
func (result *ResultWithErrorTrigger) AssertHasARecord(expectedIpv4Addr string) {
	aRecords := getRecordsByType(
		result.Result,
		"A",
		func(record dnsclient.Record) *dnsclient.A { return record.(*dnsclient.A) },
	)
	assertHasRecord(
		result,
		"A",
		aRecords,
		func(record *dnsclient.A) string { return record.Ipv4Addr },
		expectedIpv4Addr,
	)
}

// AssertHasCNAMERecord fails the test when there are no CNAME records with a matching target.
func (result *ResultWithErrorTrigger) AssertHasCNAMERecord(expectedTarget string) {
	cnameRecords := getRecordsByType(
		result.Result,
		"CNAME",
		func(record dnsclient.Record) *dnsclient.CNAME { return record.(*dnsclient.CNAME) },
	)
	assertHasRecord(
		result,
		"CNAME",
		cnameRecords,
		func(record *dnsclient.CNAME) string { return record.Target },
		expectedTarget,
	)
}

// AssertHasTXTRecord fails the test when there are no TXT records with a matching target.
func (result *ResultWithErrorTrigger) AssertHasTXTRecord(expectedTxt string) {
	cnameRecords := getRecordsByType(
		result.Result,
		"TXT",
		func(record dnsclient.Record) *dnsclient.TXT { return record.(*dnsclient.TXT) },
	)
	assertHasRecord(
		result,
		"TXT",
		cnameRecords,
		func(record *dnsclient.TXT) string { return record.Txt },
		expectedTxt,
	)
}

func assertHasRecord[RecordType dnsclient.Record](
	result *ResultWithErrorTrigger,
	recordType string,
	records []RecordType,
	getValue func(record RecordType) string,
	expectedValue string,
) {
	if err := checkMatchingRecord(result.Result, recordType, records, getValue, expectedValue); err != nil {
		result.ErrorTrigger.Errorf(err.Error())
	}
}

func checkMatchingRecord[RecordType dnsclient.Record](
	result *dnsclient.Result,
	recordType string,
	records []RecordType,
	getValue func(record RecordType) string,
	expectedValue string,
) error {
	if len(records) <= 0 {
		return fmt.Errorf("DNS assertion failed: no " + recordType + " records found")
	}

	for _, record := range records {
		if getValue(record) == expectedValue {
			return nil
		}
	}

	return fmt.Errorf(
		"DNS asserting failed: No "+recordType+" record with value %s found for %s.\nRecords Found:\n%s",
		expectedValue,
		result.FQDN,
		displayRecords(result),
	)
}

func getRecordsByType[RecordType dnsclient.Record](
	result *dnsclient.Result,
	recordType string,

	cast func(record dnsclient.Record) RecordType,
) []RecordType {
	var records []RecordType
	for _, record := range result.Records {
		if record.Type() == recordType {
			records = append(records, cast(record))
		}
	}
	return records
}

func displayRecords(result *dnsclient.Result) string {
	return "\t" + strings.Join(result.GetRawRecords(), "\n\t") + "\n"
}
