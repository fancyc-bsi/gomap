package plugins

func init() {
	hostVulnScanPlugin := &BasePlugin{
		Name:        "HostVulnScan",
		Command:     "nmap",
		Args:        []string{"--script", "vuln", "-T4", "-sV", "{ip}"},
		IsHostBased: true,
	}
	RegisterHostPlugin(hostVulnScanPlugin)
}
