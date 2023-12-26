# GoTinyAlsa 🚀

Go bindings for [TinyAlsa](https://github.com/tinyalsa/tinyalsa)

Features:
- ⚡ Easy yet powerful API
- 📱 Android support
- 🤓☝️ well documented

## Installation
```
$ go get -u github.com/Binozo/GoTinyAlsa
```
## Troubleshooting
#### Linking
If you get the following error: `error while loading shared libraries: libtinyalsa.so.2: cannot open shared object file: No such file or directory`
and you built tinyalsa from source, make sure you set `LD_LIBRARY_PATH=/usr/local/lib` as an environment variable.

#### Deprecation warnings
Those warnings
```shell
# github.com/Binozo/GoTinyAlsa/internal/tinyapi
cgo-gcc-prolog: In function ‘_cgo_ee639364c86c_Cfunc_pcm_read’:
cgo-gcc-prolog:152:2: warning: ‘pcm_read’ is deprecated [-Wdeprecated-declarations]
In file included from /usr/local/include/tinyalsa/asoundlib.h:33,
from ../internal/tinyapi/device.go:5:
```
are normal and to ensure compatibility with Android, we need to use those deprecated functions.
Take a look [here](https://github.com/tinyalsa/tinyalsa/blob/google-origin/include/tinyalsa/asoundlib.h) for available APIs for the Android platform.

## Quickstart
> [!NOTE]
> Take a look at the `examples` folder for example usage.

### Create your device
```go
package main

import "github.com/Binozo/GoTinyAlsa"

func main() {
	// You need to set deviceCard and deviceNr yourself
	// To list output devices, execute 'aplay -L'
	// To list input devices, execute 'arecord -L'
	deviceCard := 4
	deviceNr := 0
	device := tinyalsa.NewDevice(deviceCard, deviceNr, pcm.Config{
		Channels:    1,
		SampleRate:  48000,
		PeriodSize:  2048,
		PeriodCount: 20,
		Format:      tinyalsa.PCM_FORMAT_S16_LE,
	})
}
```

### Wait until it's ready
```go
    // Wait until our microphone (PCM_IN) input is ready with the maximum duration of 2 seconds
	err := device.WaitUntilReady(tinyalsa.PCM_IN, time.Second*2)
	if err != nil {
		fmt.Println("An error occurred while waiting until device is ready:", err)
	}
```