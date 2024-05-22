package plugins

import (
	"fmt"
	"os/exec"
)

type FeroxbusterPlugin struct{}

func (p *FeroxbusterPlugin) Run(ip, port string) (string, error) {
	url := fmt.Sprintf("http://%s:%s", ip, port)
	// fmt.Printf("\033[1;34m[+] Running Feroxbuster tasks on %s:%s\033[0m\n", ip, port)
	cmd := exec.Command("feroxbuster", "--url", url, "--silent", "-w", "/usr/share/seclists/Discovery/Web-Content/directory-list-2.3-medium.txt")
	output, err := cmd.CombinedOutput()
	if err != nil {
		return "", fmt.Errorf("\033[1;31m[!] feroxbuster failed: %v\033[0m", err)
	}
	return string(output), nil
}

func (p *FeroxbusterPlugin) Name() string {
	return "Feroxbuster"
}

func init() {
	RegisterPlugin("feroxbuster", &FeroxbusterPlugin{}, []string{}, []string{"http"})
}
