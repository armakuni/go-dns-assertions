package fetcher

type LookupResult struct {
	FQDN    string
	Records []Record
}

func (lookup LookupResult) GetRawRecords() []string {
	var result []string
	for _, record := range lookup.Records {
		result = append(result, record.String())
	}
	return result
}
