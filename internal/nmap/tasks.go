package nmap

import (
	"fmt"
	"path/filepath"
	"sync"

	"github.com/fancyc-bsi/gomap/internal/utils"
	"github.com/fancyc-bsi/gomap/plugins"
)

func PerformTasksBasedOnResults(nmapResult *NmapRun, outputDir, target string) error {
	var wg sync.WaitGroup
	for _, host := range nmapResult.Host {
		ip := host.Address.Addr
		if host.Status.State == "up" {
			// Run host-based plugins
			for _, plugin := range plugins.GetHostPlugins() {
				wg.Add(1)
				go func(plugin plugins.HostPlugin, ip string) {
					defer wg.Done()
					outputFile := filepath.Join(outputDir, fmt.Sprintf("%s_host_%s.txt", target, plugin.GetName()))
					fmt.Printf("\033[1;34m[+] Running plugin %s on host %s\033[0m\n", plugin.GetName(), ip)
					output, err := plugin.Run(ip, "")
					if err != nil {
						fmt.Printf("\033[1;31m[!] Failed to run plugin %s: %v\033[0m\n", plugin.GetName(), err)
						return
					}
					err = utils.WriteToFile(outputFile, output)
					if err != nil {
						fmt.Printf("\033[1;31m[!] Failed to write output for plugin %s: %v\033[0m\n", plugin.GetName(), err)
					}
				}(plugin, ip)
			}
		}
		// Run port-based plugins
		for _, port := range host.Ports.Port {
			if port.State.State == "open" {
				wg.Add(1)
				go func(ip, portId, service string) {
					defer wg.Done()
					err := handleOpenPort(ip, portId, service, outputDir, target)
					if err != nil {
						fmt.Printf("\033[1;31m[!] Failed to handle port %s: %v\033[0m\n", portId, err)
					}
				}(ip, port.PortId, port.Service.Name)
			}
		}
	}
	wg.Wait()
	return nil
}

func handleOpenPort(ip, port, service, outputDir, target string) error {
	plugins, err := plugins.GetPluginByPort(port, service)
	if err != nil {
		return fmt.Errorf("failed to get plugins for port %s: %v", port, err)
	}
	if len(plugins) == 0 {
		fmt.Printf("\033[1;33m[!] No tasks defined for port %s with service %s\033[0m\n", port, service)
		return nil
	}
	for _, plugin := range plugins {
		outputFile := filepath.Join(outputDir, fmt.Sprintf("%s_%s_%s.txt", target, service, port))
		fmt.Printf("\033[1;34m[+] Running plugin %s on %s:%s\033[0m\n", plugin.GetName(), ip, port)
		output, err := plugin.Run(ip, port)
		if err != nil {
			return fmt.Errorf("\033[1;31m[!] Failed to run plugin %s: %v\033[0m", plugin.GetName(), err)
		}
		err = utils.WriteToFile(outputFile, output)
		if err != nil {
			return fmt.Errorf("\033[1;31m[!] Failed to write output for port %s: %v\033[0m", port, err)
		}
	}
	return nil
}
