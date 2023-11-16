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