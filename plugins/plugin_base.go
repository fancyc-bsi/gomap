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
	Protocol    string
	Timeout     time.Duration
	UseSudo     bool
}

func (p *BasePlugin) Run(ip, port string) (string, error) {
	cmdArgs := make([]string, len(p.Args))
	copy(cmdArgs, p.Args)

	protocol := "http"
	if p.Protocol != "" {
		protocol = p.Protocol
	}
	url := fmt.Sprintf("%s://%s:%s", protocol, ip, port)
	// fmt.Printf("\033[1;34m[+] URL: %s\033[0m\n", url) // Debug print

	for i, arg := range cmdArgs {
		cmdArgs[i] = strings.ReplaceAll(arg, "{url}", url)
		cmdArgs[i] = strings.ReplaceAll(cmdArgs[i], "{ip}", ip)
		cmdArgs[i] = strings.ReplaceAll(cmdArgs[i], "{port}", port)
	}

	// fmt.Printf("\033[1;34m[+] Running %s tasks on %s\033[0m\n", p.Name, ip)
	// fmt.Printf("\033[1;34m[+] Command: %s %s\033[0m\n", p.Command, strings.Join(cmdArgs, " "))

	timeout := p.Timeout
	if timeout == 0 {
		timeout = 10 * time.Minute
	}
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	if p.UseSudo {
		cmdArgs = append([]string{p.Command}, cmdArgs...)
		cmdArgs = append([]string{"sudo"}, cmdArgs...)
		cmd := exec.CommandContext(ctx, "sudo", cmdArgs...)
		output, err := cmd.CombinedOutput()
		if ctx.Err() == context.DeadlineExceeded {
			return "", fmt.Errorf("\033[1;31m[!] %s timed out\033[0m", p.Name)
		}
		if err != nil {
			return "", fmt.Errorf("\033[1;31m[!] %s failed: %v\nOutput:\n%s\033[0m", p.Name, err, output)
		}
		return string(output), nil
	} else {
		cmd := exec.CommandContext(ctx, p.Command, cmdArgs...)
		output, err := cmd.CombinedOutput()
		if ctx.Err() == context.DeadlineExceeded {
			return "", fmt.Errorf("\033[1;31m[!] %s timed out\033[0m", p.Name)
		}
		if err != nil {
			return "", fmt.Errorf("\033[1;31m[!] %s failed: %v\nOutput:\n%s\033[0m", p.Name, err, output)
		}
		return string(output), nil
	}
}

func (p *BasePlugin) GetName() string {
	return p.Name
}
