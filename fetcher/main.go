package fetcher

type Fetcher interface {
	FetchDNSRecords(fqdn string, dnsServer string) (*LookupResult, error)
}
