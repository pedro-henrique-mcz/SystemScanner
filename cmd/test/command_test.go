package command

import (
	"SystemScanner/model"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"os/exec"
	"strings"
	"testing"
)

func check(mensage string, e error) {
	if e != nil {
		fmt.Println(mensage)
		panic(e)
	}
}

func TestCommand(t *testing.T) {

	computers := make([]model.Computer, 0)

	file, err := os.OpenFile("../../json/computers.json", os.O_RDONLY, 0644)
	check("Erro ao abrir arquivo", err)

	defer file.Close()

	computersJson, err := io.ReadAll(file)
	check("Erro ao ler aquivo", err)

	json.Unmarshal(computersJson, &computers)

	//check if the computer alredy exists in the file

	cmd := exec.Command("powershell", "-Command", "(Get-NetAdapter | Where-Object {$_.Status -eq 'Up'} | Select-Object -ExpandProperty MacAddress -First 1)")
	output, err := cmd.Output()
	stringOutput := string(output)
	check("Erro ao recuperar mac", err)

	system := model.GetNewSystem()
	network := model.GetNewNetwork()
	hardware := model.GetNewHardware()
	Computer := model.GetNewComputer(system, network, hardware)

	for key, computer := range computers {

		computerListMac := computer.Network.MACAddress

		if strings.TrimSpace(stringOutput) == computerListMac {

			computers[key] = *Computer
			fmt.Println(len(computers))
			return
		}

	}

	computers = append(computers, *Computer)
	fmt.Println(len(computers))

}
