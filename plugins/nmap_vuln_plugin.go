package plugins

func init() {
	hostVulnScanPlugin := &BasePlugin{
		Name:        "HostVulnScan",
		Command:     "nmap",
		Args:        []string{"--script", "vuln", "{ip}"},
		IsHostBased: true,
	}
	RegisterHostPlugin(hostVulnScanPlugin)
}
