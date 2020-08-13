package pkg

import "strings"

func awvs(req Request, tool *string) bool {
	*tool = "AWVS"
	if strings.Index(req.URL, "acunetix") != -1 {
		return true
	}
	if strings.Index(req.Body, "acunetix") != -1 {
		return true
	}
	for key, value := range req.Header {
		if strings.Index(strings.ToLower(key), "acunetix") != -1 ||
			strings.Index(value, "acunetix") != -1 {
			return true
		}
	}
	*tool = ""
	return false
}
