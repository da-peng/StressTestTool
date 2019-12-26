package base

//APIInfo  include of method url datatye and headers
type APIInfo struct {
	Method   string
	URL      string
	DataType string
	Headers  map[string]string
}
