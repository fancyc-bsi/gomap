package plugins

func init() {
	ftpAnonymousPlugin := &BasePlugin{
		Name:        "FTPAnonymous",
		Command:     "nmap",
		Args:        []string{"-p", "{port}", "--script", "ftp-anon", "{ip}"},
		IsHostBased: false,
	}
	RegisterPlugin(ftpAnonymousPlugin, []string{"21"}, []string{"ftp"})
}
