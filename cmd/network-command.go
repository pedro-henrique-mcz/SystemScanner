package command

import (
	"fmt"
	"log"
	"os/exec"
	"strings"
)

func GetNetworkData() map[string]string {

	output := make(map[string]string)

	commands := map[string]string{
		"ip_address":      "(Get-NetIPAddress | Where-Object { $_.AddressFamily -eq 'IPv4' -and $_.IPAddress -notlike '169.*' } | Select-Object -ExpandProperty IPAddress -First 1)",
		"ipv6_address":    "(Get-NetIPAddress | Where-Object { $_.AddressFamily -eq 'IPv6' } | Select-Object -ExpandProperty IPAddress -First 1)",
		"mac_address":     "(Get-NetAdapter | Where-Object {$_.Status -eq 'Up'} | Select-Object -ExpandProperty MacAddress -First 1)",
		"gateway":         "(Get-NetIPAddress | Where-Object { $_.AddressFamily -eq 'IPv4' -and $_.DefaultGateway -ne $null } | Select-Object -ExpandProperty DefaultGateway -First 1)",
		"subnet_mask":     "(Get-NetIPAddress | Where-Object { $_.AddressFamily -eq 'IPv4' } | Select-Object -ExpandProperty PrefixLength -First 1)",
		"dns":             "(Get-DnsClientServerAddress -AddressFamily IPv4 | Select-Object -ExpandProperty ServerAddresses) -join ','",
		"connection_type": "(Get-NetAdapter | Where-Object {$_.Status -eq \"Up\"} | Select-Object -ExpandProperty Name -First 1)",
		"protocol":        "\"TCP/IP\"",
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
