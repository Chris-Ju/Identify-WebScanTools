package pkg

import "strings"

func nessus(req Request, tool *string) bool {
	*tool = "Nessus"
	if strings.Index(req.URL, "nessus") != -1 {
		return true
	}
	if strings.Index(req.Body, "nessus") != -1 {
		return true
	}
	for _, value := range req.Header {
		if strings.Index(value, "nessus") != -1 {
			return true
		}
	}
	*tool = ""
	return false
}
