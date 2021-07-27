package villagers

import "time"

type Villager interface {
	Appearances()
	// Astrology is the villagers astrological star sign.
	Astrology() string
	// AstrologyIcon is the villagers astrological start sign emoji.
	AstrologyIcon() string
	// Animal is the villagers animal type.
	Animal() string
	// Birthday is the villagers birthday calendar day.
	Birthday() uint8
	// BirthdayMonth is the villagers birthday calendar month.
	BirthdayMonth() time.Month
	// Code is the villagers alpha-numeric identifier.
	Code() string
	// Gender is the villagers gender.
	Gender() string
	Icon()
	// Id is the villagers identifier.
	Id() string
	Languages()
	// Personality is the villagers personality type.
	Personality() string
	// Special is whether the villager is a special or normal character.
	Special() bool
	// Species is the villageres species type.
	Species() string
}
