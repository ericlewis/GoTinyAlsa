package tinyalsa

import (
	"time"

	"github.com/ericlewis/GoTinyAlsa/internal/tinyapi"
)

type AudioSession struct {
	pcmDevice tinyapi.PcmDevice
}

// NewAudioSession creates a new Audio session to stream audio data.
// Do not forget to call `Close()`
func (d *AlsaDevice) NewAudioSession() (AudioSession, error) {
	pcmDevice, err := tinyapi.PcmOpen(d.Card, d.Device, PCM_OUT, d.DeviceConfig)
	if err != nil {
		return AudioSession{}, err
	}
	return AudioSession{
		pcmDevice: pcmDevice,
	}, nil
}

// Close the open pcm device and free it for other use
func (a *AudioSession) Close() {
	a.pcmDevice.Close()
}

func (a *AudioSession) Stop() {
	a.pcmDevice.Stop()
}

// BufferSize returns the buffer size expected by the device
func (a *AudioSession) BufferSize() int {
	return a.pcmDevice.FrameBytesSize()
}

// BytesPerFrame returns the number of bytes per frame
func (a *AudioSession) BytesPerFrame() int {
	return a.pcmDevice.BytesPerFrame()
}

// BitsPerSample returns the number of bits per sample
func (a *AudioSession) BitsPerSample() uint16 {
	return a.pcmDevice.BitsPerSample()
}

// IsReady checks if the device is ready
func (a *AudioSession) IsReady() bool {
	return a.pcmDevice.IsReady()
}

// WaitUntilReady waits until the device is ready or timeout occurs
func (a *AudioSession) WaitUntilReady(timeout time.Duration) error {
	return a.pcmDevice.WaitUntilReady(timeout)
}

// Pump the audio data into the device.
func (a *AudioSession) Pump(data []byte) error {
	return a.pcmDevice.WriteFrames(data, len(data))
}

// Read reads audio data from the device
func (a *AudioSession) Read(buffer []byte) error {
	return a.pcmDevice.ReadFrames(buffer, len(buffer))
}

// GetError returns the current error state of the device
func (a *AudioSession) GetError() error {
	return a.pcmDevice.GetError()
}
