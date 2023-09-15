package fetcher_test

import (
	"github.com/armakuni/go-dns-assertions/fetcher"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLookupResult_GetRawRecords(t *testing.T) {
	lookup := fetcher.LookupResult{
		FQDN: "example.com.",
		Records: []fetcher.Record{
			fetcher.A{Base: &fetcher.Base{Raw: "RAW A RECORD"}, Ipv4Addr: "1.2.3.4"},
			fetcher.CNAME{Base: &fetcher.Base{Raw: "RAW CNAME RECORD"}, Target: "web.host.com."},
		},
	}

	assert.Equal(t, []string{"RAW A RECORD", "RAW CNAME RECORD"}, lookup.GetRawRecords())
}
