package tinyalsa

import (
	"GoTinyAlsa/internal/tinyapi"
	"GoTinyAlsa/pkg/pcm"
)

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

func (d *AlsaDevice) GetInfo() DeviceInfo {
	inInfo, outInfo := tinyapi.GetParams(d.Card, d.Device)
	return DeviceInfo{
		In:  inInfo,
		Out: outInfo,
	}
}
