package main

import (
	"bytes"
	"fmt"
	"github.com/Binozo/GoTinyAlsa/pkg/pcm"
	"github.com/Binozo/GoTinyAlsa/pkg/tinyalsa"
	"github.com/youpy/go-wav"
	"os"
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
	fmt.Println("Stopping...")
	close(audioChan)

	wavBuf := new(bytes.Buffer)
	writer := wav.NewWriter(wavBuf, 0, uint16(device.DeviceConfig.Channels), uint32(device.DeviceConfig.SampleRate), device.DeviceConfig.Format.BitsPerSample())
	writer.Write(rawAudioData.Bytes())
	os.WriteFile("myRecording.wav", wavBuf.Bytes(), os.ModePerm)
}
