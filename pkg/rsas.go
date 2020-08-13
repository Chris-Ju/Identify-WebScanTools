package pkg

import "strings"

func rsas(req Request, tool *string) bool {
	*tool = "Rsas"
	if strings.Index(req.URL, "nsfocus") != -1 {
		return true
	}
	if strings.Index(req.Header["User-Agent"], "Rsas") != -1 {
		return true
	}
	*tool = ""
	return false
}
