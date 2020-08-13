package pkg

import "strings"

func netsparker(req Request, tool *string) bool {
	*tool = "Netsparker"
	if strings.Index(strings.ToLower(req.URL), "netsparker") != -1 {
		return true
	}
	if strings.Index(req.Body, "netsparker") != -1 {
		return true
	}
	for _, value := range req.Header {
		if strings.Index(strings.ToLower(value), "netsparker") != -1 {
			return true
		}
	}
	*tool = ""
	return false
}
