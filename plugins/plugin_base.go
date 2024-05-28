package plugins

import (
	"context"
	"fmt"
	"os/exec"
	"strings"
	"time"
)

type BasePlugin struct {
	Name        string
	Command     string
	Args        []string
	IsHostBased bool
	Protocol    string // Optional, for HTTP/HTTPS plugins
	Timeout     time.Duration
}

func (p *BasePlugin) Run(ip, port string) (string, error) {
	cmdArgs := make([]string, len(p.Args))
	copy(cmdArgs, p.Args)

	url := fmt.Sprintf("%s://%s:%s", p.Protocol, ip, port)
	for i, arg := range cmdArgs {
		cmdArgs[i] = strings.ReplaceAll(arg, "{url}", url)
		cmdArgs[i] = strings.ReplaceAll(arg, "{ip}", ip)
		cmdArgs[i] = strings.ReplaceAll(arg, "{port}", port)
	}

	// fmt.Printf("\033[1;34m[+] Running %s tasks on %s\033[0m\n", p.Name, ip)
	// fmt.Printf("\033[1;34m[+] Command: %s %s\033[0m\n", p.Command, strings.Join(cmdArgs, " "))

	ctx, cancel := context.WithTimeout(context.Background(), p.Timeout)
	defer cancel()

	cmd := exec.CommandContext(ctx, p.Command, cmdArgs...)
	output, err := cmd.CombinedOutput()

	if ctx.Err() == context.DeadlineExceeded {
		return "", fmt.Errorf("\033[1;31m[!] %s timed out\033[0m", p.Name)
	}

	if err != nil {
		return "", fmt.Errorf("\033[1;31m[!] %s failed: %v\033[0m", p.Name, err)
	}
	return string(output), nil
}

func (p *BasePlugin) GetName() string {
	return p.Name
}
