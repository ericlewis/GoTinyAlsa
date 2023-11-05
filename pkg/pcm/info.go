package pcm

type Info struct {
	Access         uint
	Format0        uint
	Format1        uint
	FormatNames    []string
	Subformat      uint
	RateMin        uint
	RateMax        uint
	ChannelsMin    uint
	ChannelsMax    uint
	SampleBitsMin  uint
	SampleBitsMax  uint
	PeriodSizeMin  uint
	PeriodSizeMax  uint
	PeriodCountMin uint
	PeriodCountMax uint
}
