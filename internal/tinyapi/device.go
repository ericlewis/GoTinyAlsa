package tinyapi

// #include <stdio.h>
// #include <stdlib.h>
// #include <tinyalsa/asoundlib.h>
import "C"
import (
	"errors"
	"fmt"
	"github.com/Binozo/GoTinyAlsa/pkg/pcm"
	"unsafe"
)

type PcmDevice struct {
	pcmDevice *C.struct_pcm
	Config    pcm.Config
}

func (d *PcmDevice) FrameBytesSize() int {
	return int(C.pcm_frames_to_bytes(d.pcmDevice, C.pcm_get_buffer_size(d.pcmDevice)))
}

func (d *PcmDevice) BytesPerFrame() int {
	return int(C.pcm_frames_to_bytes(d.pcmDevice, 1))
}

func (d *PcmDevice) BitsPerSample() uint16 {
	return d.Config.Format.BitsPerSample()
}

func (d *PcmDevice) ReadFrames(buffer []byte, size int) error {
	framesRead := C.uint(C.pcm_read(d.pcmDevice, unsafe.Pointer(&buffer[0]), C.uint(size)))
	if framesRead != 0 {
		// Error occurred
		return errors.New(fmt.Sprintf("couldn't read frames:%d", int(framesRead)))
	}
	return nil
}

func (d *PcmDevice) WriteFrames(buffer []byte, size int) error {
	framesWritten := C.uint(C.pcm_write(d.pcmDevice, unsafe.Pointer(&buffer[0]), C.uint(size)))
	if framesWritten != 0 {
		// Error occurred
		return errors.New(fmt.Sprintf("couldn't write frames:%d", int(framesWritten)))
	}
	return nil
}

func (d *PcmDevice) Close() {
	C.pcm_close(d.pcmDevice)
}
