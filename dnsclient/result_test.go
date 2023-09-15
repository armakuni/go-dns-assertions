package dnsclient

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestResult_GetRawRecords(t *testing.T) {
	result := Result{
		FQDN: "example.com.",
		Records: []Record{
			A{Common: &Common{Raw: "RAW A RECORD"}, Ipv4Addr: "1.2.3.4"},
			CNAME{Common: &Common{Raw: "RAW CNAME RECORD"}, Target: "web.host.com."},
		},
	}

	assert.Equal(t, []string{"RAW A RECORD", "RAW CNAME RECORD"}, result.GetRawRecords())
}
