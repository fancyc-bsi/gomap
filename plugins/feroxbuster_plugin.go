package plugins

func init() {
	feroxbusterPlugin := &BasePlugin{
		Name:        "Feroxbuster",
		Command:     "feroxbuster",
		Args:        []string{"--silent", "--smart", "--auto-bail", "-x", "txt,php,html", "--url", "{url}"},
		IsHostBased: false,
		Protocol:    "http", // Default to HTTP, can be changed to HTTPS if needed
	}
	RegisterPlugin(feroxbusterPlugin, []string{}, []string{"http", "https"})
}
