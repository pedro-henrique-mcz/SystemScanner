package model

import command "SystemScanner/cmd"

type Hardware struct {
	CPU             string `json:"cpu"`
	RamStorage      string `json:"ram_storage"`
	RamType         string `json:"ram_type"`
	StorageType     string `json:"storage_type"`
	StorageCapacity string `json:"storage_capacity"`
	MotherBoard     string `json:"mother_board"`
}

func GetNewHardware() Hardware {

	data := command.GetHardwareData()

	return Hardware{
		CPU:             data["cpu"],
		RamStorage:      data["ram_storage"],
		RamType:         data["ram_type"],
		StorageType:     data["storage_type"],
		StorageCapacity: data["storage_capacity"],
		MotherBoard:     data["mother_board"],
	}

}
