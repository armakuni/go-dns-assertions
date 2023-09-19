package dnsclient

type Record interface {
	Type() string
	String() string
}

type Common struct {
	Raw string
}

func (base Common) String() string {
	return base.Raw
}

type A struct {
	*Common
	Ipv4Addr string
}

func (_ A) Type() string {
	return "A"
}

type CNAME struct {
	*Common
	Target string
}

func (_ CNAME) Type() string {
	return "CNAME"
}

type TXT struct {
	*Common
	Txt string
}

func (_ TXT) Type() string {
	return "TXT"
}
