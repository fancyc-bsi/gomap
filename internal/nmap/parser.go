package nmap

import (
	"encoding/xml"
)

type NmapRun struct {
	Host []struct {
		Address struct {
			Addr string `xml:"addr,attr"`
		} `xml:"address"`
		Status struct {
			State string `xml:"state,attr"`
		} `xml:"status"`
		Ports struct {
			Port []struct {
				PortId string `xml:"portid,attr"`
				State  struct {
					State string `xml:"state,attr"`
				} `xml:"state"`
				Service struct {
					Name string `xml:"name,attr"`
				} `xml:"service"`
			} `xml:"port"`
		} `xml:"ports"`
	} `xml:"host"`
}

func ParseNmapResults(xmlData string) (*NmapRun, error) {
	var result NmapRun
	err := xml.Unmarshal([]byte(xmlData), &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}
