package tinyalsa

import "GoTinyAlsa/pkg/pcm"

type DeviceInfo struct {
	Out pcm.Info
	In  pcm.Info
}

func (i *DeviceInfo) HasOutput() bool {
	return i.Out.Access != 0 && i.Out.RateMin != 0
}

func (i *DeviceInfo) HasInput() bool {
	return i.In.Access != 0 && i.In.RateMin != 0
}
