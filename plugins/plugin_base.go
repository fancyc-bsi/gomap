package plugins

import (
	"fmt"
	"os/exec"
	"strings"
)

type BasePlugin struct {
	Name        string
	Command     string
	Args        []string
	IsHostBased bool
	Protocol    string // Add protocol field
}

func (p *BasePlugin) Run(ip, port string) (string, error) {
	cmdArgs := make([]string, len(p.Args))
	copy(cmdArgs, p.Args)

	// Default to http if Protocol is not set
	protocol := "http"
	if p.Protocol != "" {
		protocol = p.Protocol
	}
	url := fmt.Sprintf("%s://%s:%s", protocol, ip, port)
	for i, arg := range cmdArgs {
		cmdArgs[i] = strings.ReplaceAll(arg, "{url}", url)
		cmdArgs[i] = strings.ReplaceAll(cmdArgs[i], "{ip}", ip)
		cmdArgs[i] = strings.ReplaceAll(cmdArgs[i], "{port}", port)
	}

	// fmt.Printf("\033[1;34m[+] Running %s tasks on %s\033[0m\n", p.Name, ip)
	// fmt.Printf("\033[1;34m[+] Command: %s %s\033[0m\n", p.Command, strings.Join(cmdArgs, " "))
	// fmt.Printf("\033[1;34m[+] Command Args: %v\033[0m\n", cmdArgs)

	cmd := exec.Command(p.Command, cmdArgs...)
	output, err := cmd.CombinedOutput()
	if err != nil {
		return "", fmt.Errorf("\033[1;31m[!] %s failed: %v\033[0m", p.Name, err)
	}
	return string(output), nil
}

func (p *BasePlugin) GetName() string {
	return p.Name
}
