package pkg

import "strings"

func webreaver(req Request, tool *string) bool {
	*tool = "WebReaver"
	if strings.Index(req.Header["User-Agent"], "WebReaver") != -1 {
		return true
	}
	*tool = ""
	return false
}
