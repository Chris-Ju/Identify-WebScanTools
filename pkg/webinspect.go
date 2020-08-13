package pkg

import "strings"

func webinspect(req Request, tool *string) bool {
	*tool = "Webinspect"
	if strings.Index(req.URL, "HP404") != -1 {
		return true
	}
	if strings.Index(req.Body, "Webinspect") != -1 {
		return true
	}
	if strings.Index(req.Header["User-Agent"], "HP ASC") != -1 {
		return true
	}
	if strings.Index(req.Header["Cookie"], "webinspect") != -1 {
		return true
	}
	if strings.Index(req.Header["Cookie"], "CustomCookie") != -1 {
		return true
	}
	if _, ok := req.Header["X-WIPP"]; ok {
		return true
	}
	if _, ok := req.Header["X-Request-Memo"]; ok {
		return true
	}
	if _, ok := req.Header["X-Scan-Memo"]; ok {
		return true
	}
	if _, ok := req.Header["X-RequestManager-Memo"]; ok {
		return true
	}
	*tool = ""
	return false
}
