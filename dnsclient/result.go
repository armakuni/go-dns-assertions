package dnsclient

type Result struct {
	FQDN    string
	Records []Record
}

func (result Result) GetRawRecords() []string {
	var rawRecords []string
	for _, record := range result.Records {
		rawRecords = append(rawRecords, record.String())
	}
	return rawRecords
}
