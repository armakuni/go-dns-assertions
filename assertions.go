package dnsassertions

import "github.com/armakuni/go-dns-assertions/fetcher"

// AssertHasARecord fails the test when there are no A records with a matching IP address.
func (lookup *LookupResultWithErrorable) AssertHasARecord(expectedIpv4Addr string) {
	var aRecords []*fetcher.A
	for _, record := range lookup.Records {
		if record.Type() == "A" {
			aRecords = append(aRecords, record.(*fetcher.A))
		}
	}

	if len(aRecords) <= 0 {
		lookup.Errorable.Errorf("DNS assertion failed: no A records found")
		return
	}

	found := false
	recordsString := ""

	for _, aRecord := range aRecords {
		recordsString = recordsString + "\t" + aRecord.Raw + "\n"

		if aRecord.Ipv4Addr == expectedIpv4Addr {
			found = true
		}
	}

	if !found {
		lookup.Errorable.Errorf(
			"DNS asserting failed: No A record with value %s found for %s.\nRecords Found:\n%s",
			expectedIpv4Addr,
			lookup.FQDN,
			recordsString,
		)
	}
}

// AssertHasCNAMERecord fails the test when there are no CNAME records with a matching target.
func (lookup *LookupResultWithErrorable) AssertHasCNAMERecord(expectedTarget string) {
	var cnameRecords []*fetcher.CNAME
	for _, record := range lookup.Records {
		if record.Type() == "CNAME" {
			cnameRecords = append(cnameRecords, record.(*fetcher.CNAME))
		}
	}

	if len(cnameRecords) <= 0 {
		lookup.Errorable.Errorf("DNS assertion failed: no CNAME records found")
		return
	}

	found := false
	recordsString := ""

	for _, cnameRecord := range cnameRecords {
		recordsString = recordsString + "\t" + cnameRecord.Raw + "\n"

		if cnameRecord.Target == expectedTarget {
			found = true
		}
	}

	if !found {
		lookup.Errorable.Errorf(
			"DNS asserting failed: No CNAME record with value %s found for %s.\nRecords Found:\n%s",
			expectedTarget,
			lookup.FQDN,
			recordsString,
		)
	}
}
