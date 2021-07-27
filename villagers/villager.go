package villagers

type Villager interface {
	Appearances()
	Astrology()
	Animal() string
	Birthday()
	Code() string
	Gender() string
	Icon()
	Id() string
	Languages()
	Personality() string
	Special() bool
	Species() string
}
