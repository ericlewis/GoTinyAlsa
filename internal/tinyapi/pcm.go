package tinyapi

// #cgo LDFLAGS: -L/usr/local/lib -ldl -ltinyalsa
// #include <stdio.h>
// #include <stdlib.h>
// #include <tinyalsa/asoundlib.h>
import "C"
import (
	"GoTinyAlsa/pkg/pcm"
	"errors"
)

const PCM_IN = C.PCM_IN

const (
	PCM_FORMAT_INVALID = iota - 1
	/** Signed 16-bit, little endian */
	PCM_FORMAT_S16_LE
	/** Signed, 32-bit, little endian */
	PCM_FORMAT_S32_LE
	/** Signed, 8-bit */
	PCM_FORMAT_S8
	/** Signed, 24-bit (32-bit in memory), little endian */
	PCM_FORMAT_S24_LE
	/** Signed, 24-bit, little endian */
	PCM_FORMAT_S24_3LE

	/* End of compatibility section. */

	/** Signed, 16-bit, big endian */
	PCM_FORMAT_S16_BE
	/** Signed, 24-bit (32-bit in memory), big endian */
	PCM_FORMAT_S24_BE
	/** Signed, 24-bit, big endian */
	PCM_FORMAT_S24_3BE
	/** Signed, 32-bit, big endian */
	PCM_FORMAT_S32_BE
	/** 32-bit float, little endian */
	PCM_FORMAT_FLOAT_LE
	/** 32-bit float, big endian */
	PCM_FORMAT_FLOAT_BE
	/** Max of the enumeration list, not an actual format. */
	PCM_FORMAT_MAX
)

func PcmOpen(cardNr int, deviceNr int, openFlags int, config pcm.Config) (PcmDevice, error) {
	var internalConfig *C.struct_pcm_config = new(C.struct_pcm_config)
	internalConfig.channels = C.uint(config.Channels)
	internalConfig.rate = C.uint(config.SampleRate)
	internalConfig.period_size = C.uint(config.PeriodSize)
	internalConfig.period_count = C.uint(config.PeriodCount)

	var format C.enum_pcm_format = C.enum_pcm_format(config.Format)
	internalConfig.format = format

	internalConfig.start_threshold = C.ulong(config.StartThreshold)
	internalConfig.stop_threshold = C.ulong(config.StopThreshold)
	internalConfig.silence_threshold = C.ulong(config.SilenceThreshold)

	// Open device
	var pcmDevice *C.struct_pcm = C.pcm_open(C.uint(cardNr), C.uint(deviceNr), C.uint(openFlags), internalConfig)

	// Check if device is ready
	if C.pcm_is_ready(pcmDevice) == 0 {
		p := C.pcm_get_error(pcmDevice)
		s := C.GoString(p)
		return PcmDevice{}, errors.New(s)
	}
	return PcmDevice{
		pcmDevice: pcmDevice,
		Config:    config,
	}, nil
}
