package fetcher

type Record interface {
	Type() string
	String() string
}

type Base struct {
	Raw string
}

func (base Base) String() string {
	return base.Raw
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
