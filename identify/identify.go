package identify

import (
	"errors"
	"strings"

	"github.com/Chris-Ju/Identify-WebScanTools/pkg"
)

// Run 传入 Request Bytes，返回 Web 扫描工具
func Run(req []byte) (string, error) {
	r, err := unpacking(req)
	if err != nil {
		return "", err
	}
	return identify(r), nil
}

func unpacking(req []byte) (r pkg.Request, err error) {

	err = errors.New("Unpacking failed")
	header := make(map[string]string)

	buf := strings.Split(string(req), "\r\n\r\n")
	if len(buf) < 2 {
		r.Body = ""
	} else {
		r.Body = strings.Join(buf[1:], "\r\n\r\n")
	}

	headers := strings.Split(buf[0], "\r\n")

	requestLines := strings.Split(headers[0], " ")

	if len(requestLines) < 3 {
		return
	}
	r.Method = requestLines[0]
	r.URL = requestLines[1]
	r.Protocol = requestLines[2]

	count := len(headers)
	for i := 1; i < count; i++ {
		temp := strings.Split(headers[i], ":")
		if len(temp) < 2 {
			return
		}
		header[temp[0]] = strings.Join(temp[1:], ":")
	}
	r.Header = header
	err = nil
	return
}

func identify(req pkg.Request) string {
	return pkg.Identify(req)
}
