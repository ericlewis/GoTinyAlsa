package main

import (
	"fmt"
	"github.com/Binozo/GoTinyAlsa/pkg/tinyalsa"
)

func main() {
	cardNr := 4
	deviceNr := 0
	config := tinyalsa.BestDeviceConfig(cardNr, deviceNr, tinyalsa.PCM_IN)

	fmt.Println("Calculated best device config:")
	fmt.Println("Period count: ", config.PeriodCount)
	fmt.Println("Period size:  ", config.PeriodSize)
	fmt.Println("Channels:     ", config.Channels)
	fmt.Println("Sample rate:  ", config.SampleRate)
}
