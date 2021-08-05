package clyde

import "golang.org/x/text/language"

const (
	nameAmericanEnglish      string = "Clyde"
	nameCanadianFrench       string = "Dorian"
	nameDutch                string = "Clyde"
	nameFrench               string = "Dorian"
	nameGerman               string = "Tommi"
	nameItalian              string = "Lallo"
	nameJapanese             string = "デースケ"
	nameLatinAmericanSpanish string = "Moe"
	nameKorean               string = "마철이"
	nameRussian              string = "Клайд"
	nameSpanish              string = "Moe"
	nameSimplifiedChinese    string = "笛笛"
	nameTraditionalChinese   string = "笛笛"
)

var (
	name = map[language.Tag]string{
		language.AmericanEnglish:      nameAmericanEnglish,
		language.CanadianFrench:       nameCanadianFrench,
		language.Dutch:                nameDutch,
		language.French:               nameFrench,
		language.German:               nameGerman,
		language.Italian:              nameItalian,
		language.Japanese:             nameJapanese,
		language.Korean:               nameKorean,
		language.LatinAmericanSpanish: nameLatinAmericanSpanish,
		language.Russian:              nameRussian,
		language.Spanish:              nameSpanish,
		language.SimplifiedChinese:    nameSimplifiedChinese,
		language.TraditionalChinese:   nameTraditionalChinese}
)
