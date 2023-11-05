package pcm

import "github.com/Binozo/GoTinyAlsa/internal/tinypcm"

type Config struct {
	Channels         int
	SampleRate       int
	PeriodSize       int
	PeriodCount      int
	Format           tinypcm.Format
	StartThreshold   int
	StopThreshold    int
	SilenceThreshold int
}
