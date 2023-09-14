package fetcher

import "testing"

type Record interface {
	Type() string
}

type Base struct {
	Raw string
	T   *testing.T
}

type A struct {
	*Base
	Ipv4Addr string
}

func (_ A) Type() string {
	return "A"
}

type CNAME struct {
	*Base
	Target string
}

func (_ CNAME) Type() string {
	return "CNAME"
}
