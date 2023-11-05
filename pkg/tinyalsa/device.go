package tinyalsa

import "GoTinyAlsa/pkg/pcm"

type AlsaDevice struct {
	Card         int
	Device       int
	DeviceConfig pcm.Config
}

func NewDevice(cardNr int, deviceNr int, deviceConfig pcm.Config) AlsaDevice {
	return AlsaDevice{
		Card:         cardNr,
		Device:       deviceNr,
		DeviceConfig: deviceConfig,
	}
}
