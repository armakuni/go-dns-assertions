package dnsclient_test

import (
	"github.com/armakuni/go-dns-assertions/dnsclient"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAImplementsRecord(t *testing.T) {
	var record dnsclient.Record = &dnsclient.A{
		Common:   &dnsclient.Common{Raw: "RAW RECORD TEXT"},
		Ipv4Addr: "1.2.3.4",
	}
	assert.Equal(t, "A", record.Type())
	assert.Equal(t, "RAW RECORD TEXT", record.String())
}

func TestCNAMEImplementsRecord(t *testing.T) {
	var record dnsclient.Record = &dnsclient.CNAME{
		Common: &dnsclient.Common{Raw: "RAW RECORD TEXT"},
		Target: "something.example.com.",
	}
	assert.Equal(t, "CNAME", record.Type())
	assert.Equal(t, "RAW RECORD TEXT", record.String())
}
