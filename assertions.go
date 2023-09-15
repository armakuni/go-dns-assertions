package dnsassertions

import (
	"fmt"
	"github.com/armakuni/go-dns-assertions/fetcher"
	"strings"
)

// AssertHasARecord fails the test when there are no A records with a matching IP address.
func (lookup *LookupResultWithErrorable) AssertHasARecord(expectedIpv4Addr string) {
	aRecords := getRecordsByType(
		lookup.LookupResult,
		"A",
		func(record fetcher.Record) *fetcher.A { return record.(*fetcher.A) },
	)
	assertHasRecord(
		lookup,
		"A",
		aRecords,
		func(record *fetcher.A) string { return record.Ipv4Addr },
		expectedIpv4Addr,
	)
}

// AssertHasCNAMERecord fails the test when there are no CNAME records with a matching target.
func (lookup *LookupResultWithErrorable) AssertHasCNAMERecord(expectedTarget string) {
	cnameRecords := getRecordsByType(
		lookup.LookupResult,
		"CNAME",
		func(record fetcher.Record) *fetcher.CNAME { return record.(*fetcher.CNAME) },
	)
	assertHasRecord(
		lookup,
		"CNAME",
		cnameRecords,
		func(record *fetcher.CNAME) string { return record.Target },
		expectedTarget,
	)
}

func assertHasRecord[RecordType fetcher.Record](
	lookup *LookupResultWithErrorable,
	recordType string,
	records []RecordType,
	getValue func(record RecordType) string,
	expectedValue string,
) {
	if err := checkMatchingRecord(lookup.LookupResult, recordType, records, getValue, expectedValue); err != nil {
		lookup.Errorable.Errorf(err.Error())
	}
}

func checkMatchingRecord[RecordType fetcher.Record](
	lookup *fetcher.LookupResult,
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
		lookup.FQDN,
		displayRecords(lookup),
	)
}

func getRecordsByType[RecordType fetcher.Record](
	lookup *fetcher.LookupResult,
	recordType string,

	cast func(record fetcher.Record) RecordType,
) []RecordType {
	var records []RecordType
	for _, record := range lookup.Records {
		if record.Type() == recordType {
			records = append(records, cast(record))
		}
	}
	return records
}

func displayRecords(lookup *fetcher.LookupResult) string {
	return "\t" + strings.Join(lookup.GetRawRecords(), "\n\t") + "\n"
}
