package pkg

// Identify identify webscan tools
func Identify(req Request) string {
	tool := ""
	if awvs(req, &tool) || netsparker(req, &tool) || appscan(req, &tool) || webinspect(req, &tool) ||
		rsas(req, &tool) || nessus(req, &tool) || webreaver(req, &tool) || sqlmap(req, &tool) {
	}
	return tool
}
