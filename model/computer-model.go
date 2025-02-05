package model

type Computer struct {
	System   System
	Network  Network
	Hardware Hardware
}

func GetNewComputer(system System, network Network, hardware Hardware) *Computer {
	computer := &Computer{
		System:   system,
		Network:  network,
		Hardware: hardware,
	}

	return computer
}
