package dnsassertions

import (
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

func assertHasRecord[T fetcher.Record](
	lookup *LookupResultWithErrorable,
	recordType string,
	records []T,
	getValue func(record T) string,
	expectedValue string,
) {
	if len(records) <= 0 {
		lookup.Errorable.Errorf("DNS assertion failed: no " + recordType + " records found")
		return
	}

	for _, record := range records {
		if getValue(record) == expectedValue {
			return
		}
	}

	lookup.Errorable.Errorf(
		"DNS asserting failed: No "+recordType+" record with value %s found for %s.\nRecords Found:\n%s",
		expectedValue,
		lookup.FQDN,
		displayRecords(lookup.LookupResult),
	)
}

func getRecordsByType[T fetcher.Record](lookup *fetcher.LookupResult, recordType string, cast func(record fetcher.Record) T) []T {
	var records []T
	for _, record := range lookup.Records {
		if record.Type() == recordType {
			records = append(records, cast(record))
		}
	}
	return records
}

func displayRecords(lookup *fetcher.LookupResult) string {
	return "\t" + strings.Join(getRawRecords(lookup), "\n\t") + "\n"
}

func getRawRecords(lookup *fetcher.LookupResult) []string {
	var result []string
	for _, record := range lookup.Records {
		result = append(result, record.String())
	}
	return result
}
