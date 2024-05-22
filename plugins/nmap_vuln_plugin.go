package plugins

import (
	"fmt"
	"os/exec"
)

type HostVulnScanPlugin struct{}

func (p *HostVulnScanPlugin) Run(ip, _ string) (string, error) {
	// fmt.Printf("\033[1;34m[+] Running HostVulnScan tasks on %s\033[0m\n", ip)
	cmd := exec.Command("nmap", "--script", "vuln", ip)
	output, err := cmd.CombinedOutput()
	if err != nil {
		return "", fmt.Errorf("\033[1;31m[!] nmap vuln script failed: %v\033[0m", err)
	}
	return string(output), nil
}

func (p *HostVulnScanPlugin) Name() string {
	return "HostVulnScan"
}

func init() {
	RegisterHostPlugin("hostvulnscan", &HostVulnScanPlugin{})
}
