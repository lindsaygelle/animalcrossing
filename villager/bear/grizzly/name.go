package grizzly

import "golang.org/x/text/language"

const (
	nameAmericanEnglish      string = "Grizzly"
	nameCanadianFrench       string = "Grizzly"
	nameDutch                string = "Grizzly"
	nameFrench               string = "Grizzly"
	nameGerman               string = "Gerald"
	nameItalian              string = "Grizzly"
	nameJapanese             string = "ムー"
	nameLatinAmericanSpanish string = "Gruñón"
	nameKorean               string = "무뚝"
	nameRussian              string = "Гризли"
	nameSpanish              string = "Gruñón"
	nameSimplifiedChinese    string = "穆穆"
	nameTraditionalChinese   string = "穆穆"
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
