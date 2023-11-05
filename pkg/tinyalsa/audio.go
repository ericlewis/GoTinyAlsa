package tinyalsa

import (
	"GoTinyAlsa/internal/tinyapi"
	"GoTinyAlsa/pkg/pcm"
)

const PCM_IN = tinyapi.PCM_IN
const PCM_FORMAT_S16_LE = tinyapi.PCM_FORMAT_S16_LE
const PCM_FORMAT_S24_LE = tinyapi.PCM_FORMAT_S24_LE
const PCM_FORMAT_S32_LE = tinyapi.PCM_FORMAT_S32_LE
const ErrorTolerance = 10 // defines how many error frames are allowed to be read without stopping reading the next ones

func (d *AlsaDevice) GetAudioStream(config pcm.Config, audioData chan []byte) error {
	pcmDevice, err := tinyapi.PcmOpen(d.Card, d.Device, PCM_IN, config)
	// TODO Check for hw: params error -> Recommend supported config
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
		default:
			// Chan got closed, we need to close too
			break FrameReader
		}
	}

	return nil
}
