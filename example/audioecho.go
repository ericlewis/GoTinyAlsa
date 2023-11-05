package main

import (
	"bytes"
	"fmt"
	"github.com/Binozo/GoTinyAlsa/pkg/pcm"
	"github.com/Binozo/GoTinyAlsa/pkg/tinyalsa"
	"time"
)

func main() {
	device := tinyalsa.NewDevice(4, 0, pcm.Config{
		Channels:    1,
		SampleRate:  48000,
		PeriodSize:  2048,
		PeriodCount: 20,
		Format:      tinyalsa.PCM_FORMAT_S16_LE,
	})
	audioChan := make(chan []byte)
	go func() {
		err := device.GetAudioStream(device.DeviceConfig, audioChan)
		if err != nil {
			panic(err)
		}
	}()
	seconds := 5
	fmt.Println(fmt.Sprintf("Listening for %d seconds", seconds))
	listenStart := time.Now()
	rawAudioData := new(bytes.Buffer)
	for {
		if time.Now().Sub(listenStart).Seconds() >= float64(seconds) {
			break
		}
		audioData := <-audioChan
		rawAudioData.Write(audioData)
	}
	close(audioChan)

	sendDevice := tinyalsa.NewDevice(4, 0, pcm.Config{
		Channels:    2,
		SampleRate:  24000,
		PeriodSize:  1024,
		PeriodCount: 10,
		Format:      tinyalsa.PCM_FORMAT_S16_LE,
	})
	err := sendDevice.SendAudioStream(rawAudioData.Bytes())
	if err != nil {
		fmt.Println(err)
	}
}
