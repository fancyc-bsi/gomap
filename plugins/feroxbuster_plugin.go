package plugins

func init() {
	feroxbusterPlugin := &BasePlugin{
		Name:        "Feroxbuster",
		Command:     "feroxbuster",
		Args:        []string{"--silent", "-x", "txt,php,html", "-w", "/usr/share/seclists/Discovery/Web-Content/directory-list-2.3-medium.txt", "--url", "{url}"},
		IsHostBased: false,
		Protocol:    "http", // Default to HTTP, can be changed to HTTPS if needed
	}
	RegisterPlugin(feroxbusterPlugin, []string{}, []string{"http", "https"})
}
