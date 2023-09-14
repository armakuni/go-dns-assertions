package fetcher_test

import (
	"github.com/armakuni/go-dns-assertions/fetcher"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAImplementsRecord(t *testing.T) {
	var record fetcher.Record = &fetcher.A{
		Base: &fetcher.Base{
			Raw: "RAW RECORD TEXT",
			T:   t,
		},
		Ipv4Addr: "1.2.3.4",
	}
	assert.Equal(t, record.Type(), "A")
}

func TestCnameImplementsRecord(t *testing.T) {
	var record fetcher.Record = &fetcher.CNAME{
		Base: &fetcher.Base{
			Raw: "RAW RECORD TEXT",
			T:   t,
		},
		Target: "something.example.com.",
	}
	assert.Equal(t, record.Type(), "CNAME")
}
