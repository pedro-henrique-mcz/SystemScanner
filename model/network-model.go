package model

import (
	command "SystemScanner/cmd"
)

type Network struct {
	IPAddress      string `json:"ip_address"`
	IPv6Address    string `json:"ipv6_address"`
	MACAddress     string `json:"mac_address"`
	Gateway        string `json:"gateway"`
	SubnetMask     string `json:"subnet_mask"`
	DNS            string `json:"dns"`
	ConnectionType string `json:"connection_type"`
	Protocol       string `json:"protocol"`
}

func GetNewNetwork() Network {

	networkData := command.GetNetworkData()

	return Network{
		IPAddress:      networkData["ip_address"],
		IPv6Address:    networkData["ipv6_address"],
		MACAddress:     networkData["mac_address"],
		Gateway:        networkData["gateway"],
		SubnetMask:     networkData["subnet_mask"],
		DNS:            networkData["dns"],
		ConnectionType: networkData["connection_type"],
		Protocol:       networkData["protocol"],
	}

}
