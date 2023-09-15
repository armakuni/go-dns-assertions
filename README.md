# DNS Assertions for Golang Tests

This is a Golang library to assert that DNS records exist by performing lookups on a given nameserver.

## Purpose

Adding tests that make requests to real DNS servers is probably not something you want to do often when building an application.
This library's primary intended use is for write infrastructure tests.
If you are using Golang to test you infrastructure as code (when using [Terratest](https://terratest.gruntwork.io/) for example) then this could be a useful tool.

## Usage

Install this library using the usual `god mod` command:

```shell
go mod get github.com/armakuni/go-dns-assertions
```

You can now write tests using Go test that look like this:

```go
func TestFetchDNSRecords(t *testing.T) {
	client := dnsassertions.NewTestClient(t)

	result := client.FetchDNSRecords("mysite.com", "8.8.8.8")
	result.AssertHasARecord("1.2.3.4")

	resultWww := client.FetchDNSRecords("www.mysite.com", "8.8.8.8")
	resultWww.AssertHasCNAMERecord("mysite.com.")
}
```

## Documentation

API docs can be found at [https://pkg.go.dev/github.com/armakuni/go-dns-assertions](https://pkg.go.dev/github.com/armakuni/go-dns-assertions).
