package plugins

func init() {
	udpScanPlugin := &BasePlugin{
		Name:        "UDPScan",
		Command:     "nmap",
		Args:        []string{"-sU", "-T4", "--top-ports", "100", "{ip}"},
		IsHostBased: true,
	}
	RegisterHostPlugin(udpScanPlugin)
}
