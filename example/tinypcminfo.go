package main

import (
	"fmt"
	"github.com/Binozo/GoTinyAlsa/pkg/pcm"
	"github.com/Binozo/GoTinyAlsa/pkg/tinyalsa"
)

func main() {
	device := tinyalsa.NewDevice(4, 0, pcm.Config{})
	deviceInfo := device.GetInfo()

	for i, info := range []pcm.Info{deviceInfo.In, deviceInfo.Out} {
		if i == 0 {
			fmt.Print("Input:")
			if !deviceInfo.HasInput() {
				fmt.Println(" Not available")
				continue
			} else {
				fmt.Println("")
			}
		} else {
			fmt.Print("Output:")
			if !deviceInfo.HasOutput() {
				fmt.Println(" Not available")
				continue
			} else {
				fmt.Println("")
			}
		}

		fmt.Printf("Access: %#08x\n", info.Access)
		fmt.Printf("Format[0]: %#08x\n", info.Format0)
		fmt.Printf("Format[1]: %#08x\n", info.Format1)
		fmt.Printf("Format: %s\n", info.FormatNames)
		fmt.Printf("Subformat: %#08x\n", info.Subformat)
		fmt.Printf("Rate: min=%dHz max=%dHz\n", info.RateMin, info.RateMax)
		fmt.Printf("Channels: min=%d max=%d\n", info.ChannelsMin, info.ChannelsMax)
		fmt.Printf("Sample bits: min=%d max=%d\n", info.SampleBitsMin, info.SampleBitsMax)
		fmt.Printf("Period size: min=%d max=%d\n", info.PeriodSizeMin, info.PeriodSizeMax)
		fmt.Printf("Period count: min=%d max=%d\n", info.PeriodCountMin, info.PeriodCountMax)
		fmt.Println()
	}
}
