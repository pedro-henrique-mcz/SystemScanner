package model

import command "SystemScanner/cmd"

type System struct {
	Name         string `json:"os_name"`
	Version      string `json:"os_version"`
	License      string `json:"os_license"`
	Architecture string `json:"architecture"`
	Host         string `json:"host_name"`
}

func GetNewSystem() System {

	systemInfo := command.GetSystemInfo()

	return System{
		Name:         systemInfo["Name"],
		Version:      systemInfo["Version"],
		License:      systemInfo["License"],
		Architecture: systemInfo["Architecture"],
		Host:         systemInfo["Host"],
	}
}
