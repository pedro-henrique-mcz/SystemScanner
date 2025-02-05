package main

import (
	"SystemScanner/helper"
	"SystemScanner/model"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"os/exec"
	"strings"
)

const JSON_PATH = "../json/computers.json"

func biuldPc() model.Computer {
	system := model.GetNewSystem()
	network := model.GetNewNetwork()
	hardware := model.GetNewHardware()
	computer := model.GetNewComputer(system, network, hardware)

	return *computer

}

func main() {

	computer := biuldPc()

	file, err := os.OpenFile(JSON_PATH, os.O_RDWR, 0644)
	helper.Check("Erro ao abrir o arquivo json.", err)

	defer file.Close()

	computersJson, err := io.ReadAll(file)
	helper.Check("Não foi possível ler o arquivo json.", err)

	computers := make([]model.Computer, 0)

	//first use case
	if string(computersJson) != "" {

		fmt.Println("entrou")
		json.Unmarshal(computersJson, &computers)

		err = os.Truncate(JSON_PATH, 0)
		helper.Check("erro ao apagar lista no json", err)

		_, err = file.Seek(0, 0)
		helper.Check("Erro ao reposicionar o ponteiro do arquivo", err)

		cmd := exec.Command("powershell", "-Command", "(Get-NetAdapter | Where-Object {$_.Status -eq 'Up'} | Select-Object -ExpandProperty MacAddress -First 1)")
		output, err := cmd.Output()
		helper.Check("Erro ao recuperar mac", err)

		stringOutput := string(output)

		found := false

		for key, computer := range computers {
			computerListMac := computer.Network.MACAddress

			if strings.TrimSpace(stringOutput) == computerListMac {
				computers[key] = computer
				found = true // Marca que o computador foi encontrado
				break        // Sai do loop, mas não da função
			}
		}

		// Se o computador não foi encontrado, adicione-o
		if !found {
			computers = append(computers, computer)
		}

		// Grava o JSON atualizado no arquivo
		jsonFile, err := json.MarshalIndent(computers, "", " ")
		helper.Check("erro ao transformar em um json", err)

		err = os.Truncate(JSON_PATH, 0)
		helper.Check("Erro ao truncar o arquivo antes de escrever", err)

		_, err = file.Write(jsonFile)
		helper.Check("erro ao escrever lista no json", err)

	} else {

		computers = append(computers, computer)

		jsonFile, err := json.MarshalIndent(computers, "", " ")
		helper.Check("erro ao transformar em um json", err)

		_, err = file.Write(jsonFile)
		helper.Check("erro ao escrever lista no json", err)

	}

}
