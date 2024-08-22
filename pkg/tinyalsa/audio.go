package tinyalsa

import (
	"bufio"
	"bytes"
	"errors"
	"io"
	"time"

	"github.com/ericlewis/GoTinyAlsa/internal/tinyapi"
	"github.com/ericlewis/GoTinyAlsa/pkg/pcm"
)

const PCM_IN = tinyapi.PCM_IN
const PCM_OUT = tinyapi.PCM_OUT
const PCM_FORMAT_S16_LE = tinyapi.PCM_FORMAT_S16_LE
const PCM_FORMAT_S16_BE = tinyapi.PCM_FORMAT_S16_BE
const PCM_FORMAT_S24_LE = tinyapi.PCM_FORMAT_S24_LE
const PCM_FORMAT_S24_BE = tinyapi.PCM_FORMAT_S24_BE
const PCM_FORMAT_S24_3BE = tinyapi.PCM_FORMAT_S24_3BE
const PCM_FORMAT_S24_3LE = tinyapi.PCM_FORMAT_S24_3LE
const PCM_FORMAT_S32_LE = tinyapi.PCM_FORMAT_S32_LE
const PCM_FORMAT_S32_BE = tinyapi.PCM_FORMAT_S32_BE
const ErrorTolerance = 10 // defines how many error frames are allowed to be read without stopping reading the next ones

// GetAudioStream Listens to the input of the defined device
func (d *AlsaDevice) GetAudioStream(config pcm.Config, audioData chan []byte) error {
	pcmDevice, err := tinyapi.PcmOpen(d.Card, d.Device, PCM_IN, config)
	if err != nil {
		return err
	}
	defer pcmDevice.Close()
	size := pcmDevice.FrameBytesSize()
	buffer := make([]byte, size)
	if err != nil {
		return err
	}

	errorCount := 0
	writeTimeout := time.Second * 5
FrameReader:
	for {
		err := pcmDevice.ReadFrames(buffer, size)
		if err != nil {
			if errorCount > ErrorTolerance {
				return err
			}
			errorCount += 1
			continue
		}
		defer func() {
			if r := recover(); r != nil {
				// Channel is probably closed!
			}
		}()
		select {
		case audioData <- buffer:
			// Successfully sent audio data back to the api
			continue
		case <-time.After(writeTimeout):
			// Chan got closed, we need to close too
			break FrameReader
		}
	}

	return nil
}

// SendAudioStream plays the given audio data (usually wav) to the device
func (d *AlsaDevice) SendAudioStream(audioData []byte) error {
	pcmDevice, err := tinyapi.PcmOpen(d.Card, d.Device, PCM_OUT, d.DeviceConfig)
	if err != nil {
		return err
	}
	defer pcmDevice.Close()
	size := pcmDevice.FrameBytesSize()
	buffer := make([]byte, size)

	fullBytes := new(bytes.Buffer)
	fullBytes.Write(audioData)
	reader := bufio.NewReader(fullBytes)

	for {
		n, err := reader.Read(buffer)
		if err != nil {
			return err
		}
		err = pcmDevice.WriteFrames(buffer, n)
		if err != nil {
			if errors.Is(err, io.EOF) {
				return nil // Completed
			}
			return err
		}
	}
}
