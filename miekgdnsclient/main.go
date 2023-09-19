package miekgdnsclient

import (
	"fmt"
	"github.com/armakuni/go-dns-assertions/dnsclient"
	"github.com/miekg/dns"
	"net"
	"strings"
)

func New() dnsclient.DNSClient {
	return &miekgDnsClient{}
}

type miekgDnsClient struct{}

func (miekgDnsClient) LookupAllRecords(fqdn string, dnsServer string) (*dnsclient.Result, error) {
	client := new(dns.Client)
	serverAddress := net.JoinHostPort(dnsServer, "53")

	aRecords, err := fetchARecords(fqdn, client, serverAddress)
	txtRecords, err := fetchTxtRecords(fqdn, client, serverAddress)

	result := &dnsclient.Result{
		FQDN:    fqdn,
		Records: append(aRecords, txtRecords...),
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

func fetchTxtRecords(fqdn string, client *dns.Client, serverAddress string) ([]dnsclient.Record, error) {
	response, err := performLookup(fqdn, dns.TypeTXT, client, serverAddress)
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
	case dns.TypeTXT:
		return &dnsclient.TXT{
			Common: base,
			Txt:    strings.Join(answer.(*dns.TXT).Txt, "\n"),
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
