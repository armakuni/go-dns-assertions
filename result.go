package dnsassertions

import "github.com/armakuni/go-dns-assertions/dnsclient"

type ErrorTrigger interface {
	Errorf(format string, args ...any)
}

type ResultWithErrorTrigger struct {
	*dnsclient.Result
	ErrorTrigger ErrorTrigger
}
