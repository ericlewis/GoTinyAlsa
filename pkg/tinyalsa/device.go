package tinyalsa

import (
	"github.com/Binozo/GoTinyAlsa/internal/tinyapi"
	"github.com/Binozo/GoTinyAlsa/internal/tinypcm"
	"github.com/Binozo/GoTinyAlsa/pkg/pcm"
)

type AlsaDevice struct {
	Card         int
	Device       int
	DeviceConfig pcm.Config
}

// NewDevice defines a new device you want to interact with
func NewDevice(cardNr int, deviceNr int, deviceConfig pcm.Config) AlsaDevice {
	return AlsaDevice{
		Card:         cardNr,
		Device:       deviceNr,
		DeviceConfig: deviceConfig,
	}
}

// BestDeviceConfig return you the best device config to use
func BestDeviceConfig(cardNr int, deviceNr int, format int) pcm.Config {
	device := AlsaDevice{
		Card:         cardNr,
		Device:       deviceNr,
		DeviceConfig: pcm.Config{},
	}
	info := device.GetInfo()
	infoData := info.In
	if format == PCM_IN {
		infoData = info.In
	} else {
		infoData = info.Out
	}
	return pcm.Config{
		Channels:    int(infoData.ChannelsMax),
		SampleRate:  int(infoData.RateMax),
		PeriodSize:  int(infoData.PeriodSizeMax),
		PeriodCount: int(infoData.PeriodCountMax),
		Format:      tinypcm.Format(format),
	}
}

// GetInfo returns information data about the given device's input & output
func (d *AlsaDevice) GetInfo() DeviceInfo {
	inInfo, outInfo := tinyapi.GetParams(d.Card, d.Device)
	return DeviceInfo{
		In:  inInfo,
		Out: outInfo,
	}
}
