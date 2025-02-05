package command

import (
	"fmt"
	"log"
	"os/exec"
	"strings"
)

func GetHardwareData() map[string]string {

	output := make(map[string]string)

	commands := map[string]string{
		"cpu":              "Get-CimInstance -ClassName Win32_Processor | Select-Object -ExpandProperty Name",
		"ram_storage":      "[math]::Round((Get-CimInstance -ClassName Win32_PhysicalMemory | Measure-Object -Property Capacity -Sum).Sum / 1GB, 2)",
		"ram_type":         "Get-CimInstance -ClassName Win32_PhysicalMemory | Select-Object -ExpandProperty MemoryType",
		"storage_type":     "Get-PhysicalDisk | Select-Object -ExpandProperty MediaType",
		"storage_capacity": "(Get-PhysicalDisk | Measure-Object -Property Size -Sum).Sum / 1GB",
		"mother_board":     "Get-CimInstance -ClassName Win32_BaseBoard | Select-Object -ExpandProperty Product",
	}

	for key, command := range commands {

		cmd := exec.Command("powershell", "-Command", command)
		commandOutput, err := cmd.Output()

		if err != nil {
			fmt.Println("Erro ao montar Network, erro no comando", key)
			log.Fatal(err)
		}

		output[key] = strings.TrimSpace(string(commandOutput))

	}

	return output

}
