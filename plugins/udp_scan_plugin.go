package plugins

import "time"

func init() {
	udpScanPlugin := &BasePlugin{
		Name:        "UDPScan",
		Command:     "nmap",
		Args:        []string{"-sU", "-T4", "--top-ports", "100", "{ip}"},
		IsHostBased: true,
		UseSudo:     true,
		Timeout:     1 * time.Minute,
	}
	RegisterHostPlugin(udpScanPlugin)
}
