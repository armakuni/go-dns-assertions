package fetcher

type LookupResult struct {
	FQDN    string
	Records []Record
}

type Fetcher interface {
	FetchDNSRecords(fqdn string, dnsServer string) (*LookupResult, error)
}
