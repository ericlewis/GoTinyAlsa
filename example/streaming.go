package main

import (
	"github.com/Binozo/GoTinyAlsa/pkg/pcm"
	"github.com/Binozo/GoTinyAlsa/pkg/tinyalsa"
)

func main() {
	// This is our speaker output
	sendDevice := tinyalsa.NewDevice(4, 0, pcm.Config{
		Channels:    2,
		SampleRate:  24000,
		PeriodSize:  1024,
		PeriodCount: 10,
		Format:      tinyalsa.PCM_FORMAT_S16_LE,
	})
	audioSession, err := sendDevice.NewAudioSession()
	if err != nil {
		panic(err)
	}
	// Do not forget to close
	defer audioSession.Close()
	bufferSize := audioSession.BufferSize()
	for {
		// Replace `audioWavData` with your wav data
		audioWavData := make([]byte, bufferSize)
		err = audioSession.Pump(audioWavData)
		if err != nil {
			panic(err)
		}
	}
}
