package nmap

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
)

func RunNmapScan(target, outputDir string) (string, error) {
	xmlOutputFile := filepath.Join(outputDir, fmt.Sprintf("%s_nmap.xml", target))
	nmapOutputFile := filepath.Join(outputDir, fmt.Sprintf("%s_nmap.nmap", target))
	cmd := exec.Command("nmap", "-sC", "-p-", "-T4", "-sV", "-oX", xmlOutputFile, "-oN", nmapOutputFile, target)
	err := cmd.Run()
	if err != nil {
		return "", err
	}

	output, err := os.ReadFile(xmlOutputFile)
	if err != nil {
		return "", err
	}
	return string(output), nil
}
