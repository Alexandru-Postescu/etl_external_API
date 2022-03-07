package model

// Representation of the output
type Output struct {
	Index        string   `json:"index"`
	Records      []Person `json:"records"`
	TotalRecords int      `json:"total_records"`
}

// TODO: you could just create a constructor for this and make it clearer
// example: func NewOutput(index string, records []Person, totalrecords int) Output
func (o *Output) Init(index string, records []Person, totalrecords int) {
	o.Index = index
	o.Records = records
	o.TotalRecords = totalrecords
}
