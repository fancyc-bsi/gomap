package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"sync"

	"github.com/fancyc-bsi/gomap/internal/nmap"
	"github.com/fancyc-bsi/gomap/internal/reporting"
	_ "github.com/fancyc-bsi/gomap/plugins" // Import to register plugins
)

func main() {
	if len(os.Args) < 2 {
		log.Fatalf("\033[1;31m[!] Usage: %s <target or target file>\033[0m", os.Args[0])
	}
	targetInput := os.Args[1]

	outputDir := "_output"
	err := os.MkdirAll(outputDir, os.ModePerm)
	if err != nil {
		log.Fatalf("\033[1;31m[!] Failed to create output directory: %v\033[0m", err)
	}

	var targets []string
	file, err := os.Open(targetInput)
	if err == nil {
		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			targets = append(targets, scanner.Text())
		}
		file.Close()
	} else {
		targets = append(targets, targetInput)
	}

	var wg sync.WaitGroup
	for _, target := range targets {
		wg.Add(1)
		go func(target string) {
			defer wg.Done()
			fmt.Printf("\033[1;34m[+] Scanning target: %s\033[0m\n", target)
			nmapOutput, err := nmap.RunNmapScan(target, outputDir)
			if err != nil {
				log.Printf("\033[1;31m[!] Failed to run nmap scan on target %s: %v\033[0m", target, err)
				return
			}

			parsedResult, err := nmap.ParseNmapResults(nmapOutput)
			if err != nil {
				log.Printf("\033[1;31m[!] Failed to parse nmap results for target %s: %v\033[0m", target, err)
				return
			}

			err = nmap.PerformTasksBasedOnResults(parsedResult, outputDir, target)
			if err != nil {
				log.Printf("\033[1;31m[!] Failed to perform tasks for target %s: %v\033[0m", target, err)
			} else {
				fmt.Printf("\033[1;32m[+] Tasks completed successfully for target %s\033[0m\n", target)
			}
		}(target)
	}
	wg.Wait()

	// Render the Markdown files into a comprehensive HTML report
	reportFile := filepath.Join(outputDir, "report.html")
	err = reporting.RenderReport(outputDir, reportFile)
	if err != nil {
		log.Fatalf("\033[1;31m[!] Failed to render report: %v\033[0m", err)
	}
	fmt.Printf("\033[1;32m[+] Report generated successfully: %s\033[0m\n", reportFile)
}
