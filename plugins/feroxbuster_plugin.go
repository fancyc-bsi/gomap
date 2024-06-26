package plugins

import "time"

func init() {
	feroxbusterPlugin := &BasePlugin{
		Name:        "Feroxbuster",
		Command:     "feroxbuster",
		Args:        []string{"--silent", "-w", "/usr/share/seclists/Discovery/Web-Content/directory-list-2.3-medium.txt", "-C," "404", "--url", "{url}"},
		IsHostBased: false,
		Timeout:     5 * time.Minute,
	}
	RegisterPlugin(feroxbusterPlugin, []string{}, []string{"http", "https"})
}
