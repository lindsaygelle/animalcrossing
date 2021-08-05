package kitt

import "golang.org/x/text/language"

const (
	nameAmericanEnglish      string = "Kitt"
	nameCanadianFrench       string = "Poquette"
	nameDutch                string = "Kitt"
	nameFrench               string = "Poquette"
	nameGerman               string = "Kerstin"
	nameItalian              string = "Kitt"
	nameJapanese             string = "アップリケ"
	nameLatinAmericanSpanish string = "Antípoda"
	nameKorean               string = "애플리케"
	nameRussian              string = "Китт"
	nameSpanish              string = "Antípoda"
	nameSimplifiedChinese    string = "缝缝"
	nameTraditionalChinese   string = "縫縫"
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
