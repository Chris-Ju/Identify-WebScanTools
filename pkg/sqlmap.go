package pkg

import "strings"

func sqlmap(req Request, tool *string) bool {
	*tool = "Sqlmap"
	if strings.Index(req.URL, "sqlmap") != -1 {
		return true
	}
	if strings.Index(req.Body, "sqlmap") != -1 {
		return true
	}
	if strings.Index(req.Header["User-Agent"], "sqlmap") != -1 {
		return true
	}
	*tool = ""
	return false
}
