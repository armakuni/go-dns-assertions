package miekgdnsclient

import (
	"fmt"
	"github.com/armakuni/go-dns-assertions/dnsclient"
	"github.com/miekg/dns"
	"net"
)

func New() dnsclient.DNSClient {
	return &miekgDnsClient{}
}

type miekgDnsClient struct{}

func (miekgDnsClient) LookupAllRecords(fqdn string, dnsServer string) (*dnsclient.Result, error) {
	client := new(dns.Client)
	serverAddress := net.JoinHostPort(dnsServer, "53")

	records, err := fetchARecords(fqdn, client, serverAddress)
	result := &dnsclient.Result{
		FQDN:    fqdn,
		Records: records,
	}
	return result, err
}

func fetchARecords(fqdn string, client *dns.Client, serverAddress string) ([]dnsclient.Record, error) {
	response, err := performLookup(fqdn, dns.TypeA, client, serverAddress)
	if err != nil {
		return nil, err
	}

	var results []dnsclient.Record

	for _, answer := range response.Answer {
		if record := recordFromDnsRR(answer); record != nil {
			results = append(results, record)
		}
	}

	return results, nil
}

func recordFromDnsRR(answer dns.RR) dnsclient.Record {
	base := &dnsclient.Common{
		Raw: answer.String(),
	}

	switch answer.Header().Rrtype {
	case dns.TypeCNAME:
		return &dnsclient.CNAME{
			Common: base,
			Target: answer.(*dns.CNAME).Target,
		}
	case dns.TypeA:
		return &dnsclient.A{
			Common:   base,
			Ipv4Addr: answer.(*dns.A).A.String(),
		}
	}

	return nil
}

func performLookup(fqdn string, recordType uint16, client *dns.Client, serverAddress string) (*dns.Msg, error) {
	query := new(dns.Msg)
	query.SetQuestion(dns.Fqdn(fqdn), recordType)
	response, _, err := client.Exchange(query, serverAddress)

	if err != nil {
		return nil, fmt.Errorf("DNS %s record lookup failed: %s", dns.TypeToString[recordType], err.Error())
	}

	return response, nil
}
