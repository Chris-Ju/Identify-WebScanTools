package pkg

// Request Http 请求
type Request struct {
	Method   string
	URL      string
	Protocol string
	Header   map[string]string
	Body     string
}
