package dnsclient

type DNSClient interface {
	LookupAllRecords(fqdn string, dnsServer string) (*Result, error)
}
