package pkg

import "strings"

func appscan(req Request, tool *string) bool {
	*tool = "Appscan"
	if strings.Index(req.URL, "Appscan") != -1 {
		return true
	}
	if strings.Index(req.Body, "Appscan") != -1 {
		return true
	}
	for _, value := range req.Header {
		if strings.Index(value, "Appscan") != -1 {
			return true
		}
	}
	*tool = ""
	return false
}
