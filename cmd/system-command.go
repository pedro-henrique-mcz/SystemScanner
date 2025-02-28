package command

import (
	"fmt"
	"os/exec"
	"strings"
)

func GetSystemInfo() map[string]string {

	outputList := make(map[string]string)

	commands := map[string]string{
		"Name":         "(Get-WmiObject Win32_OperatingSystem).Caption",
		"Version":      "(Get-WmiObject Win32_OperatingSystem).Version",
		"License":      "@(Get-CimInstance -Query 'SELECT LicenseStatus FROM SoftwareLicensingProduct WHERE PartialProductKey IS NOT NULL').LicenseStatus -contains 1",
		"Architecture": "(Get-WmiObject Win32_OperatingSystem).OSArchitecture",
		"Host":         "$env:COMPUTERNAME",
	}

	for key, value := range commands {

		cmd := exec.Command("powershell", "-Command", value)
		output, err := cmd.CombinedOutput()

		if err != nil {
			fmt.Println(err)
		}

		outputList[key] = strings.TrimSpace(string(output))

	}

	return outputList

}
