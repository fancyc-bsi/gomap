package plugins

import (
	"context"
	"fmt"
	"io"
	"os"
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

	for i, arg := range cmdArgs {
		cmdArgs[i] = strings.ReplaceAll(arg, "{url}", url)
		cmdArgs[i] = strings.ReplaceAll(cmdArgs[i], "{ip}", ip)
		cmdArgs[i] = strings.ReplaceAll(cmdArgs[i], "{port}", port)
	}

	fmt.Printf("\033[1;34m[+] Running %s tasks on %s\033[0m\n", p.Name, ip)
	fmt.Printf("\033[1;34m[+] Command: %s %s\033[0m\n", p.Command, strings.Join(cmdArgs, " "))

	timeout := p.Timeout
	if timeout == 0 {
		timeout = 10 * time.Minute
	}
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	var cmd *exec.Cmd
	if p.UseSudo {
		cmdArgs = append([]string{p.Command}, cmdArgs...)
		cmdArgs = append([]string{"sudo"}, cmdArgs...)
		cmd = exec.CommandContext(ctx, "sudo", cmdArgs...)
	} else {
		cmd = exec.CommandContext(ctx, p.Command, cmdArgs...)
	}

	stdoutPipe, err := cmd.StdoutPipe()
	if err != nil {
		return "", fmt.Errorf("\033[1;31m[!] Failed to get stdout pipe for %s: %v\033[0m", p.Name, err)
	}
	stderrPipe, err := cmd.StderrPipe()
	if err != nil {
		return "", fmt.Errorf("\033[1;31m[!] Failed to get stderr pipe for %s: %v\033[0m", p.Name, err)
	}

	outputFile := fmt.Sprintf("_output/%s_%s_%s.md", ip, p.Name, port)
	outFile, err := os.Create(outputFile)
	if err != nil {
		return "", fmt.Errorf("\033[1;31m[!] Failed to create output file for %s: %v\033[0m", p.Name, err)
	}
	defer outFile.Close()

	if err := cmd.Start(); err != nil {
		return "", fmt.Errorf("\033[1;31m[!] Failed to start %s: %v\033[0m", p.Name, err)
	}

	go io.Copy(outFile, stdoutPipe)
	go io.Copy(outFile, stderrPipe)

	err = cmd.Wait()
	if ctx.Err() == context.DeadlineExceeded {
		return "", fmt.Errorf("\033[1;31m[!] %s timed out\033[0m", p.Name)
	}
	if err != nil {
		return "", fmt.Errorf("\033[1;31m[!] %s failed: %v\033[0m", p.Name, err)
	}

	return "", nil
}

func (p *BasePlugin) GetName() string {
	return p.Name
}
