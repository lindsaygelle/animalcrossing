package miranda

import "golang.org/x/text/language"

const (
	nameAmericanEnglish      string = "Miranda"
	nameCanadianFrench       string = "Maëllis"
	nameDutch                string = "Miranda"
	nameFrench               string = "Maëllis"
	nameGerman               string = "Tanya"
	nameItalian              string = "Miranda"
	nameJapanese             string = "ミランダ"
	nameLatinAmericanSpanish string = "Miranda"
	nameKorean               string = "미란다"
	nameRussian              string = "Миранда"
	nameSpanish              string = "Miranda"
	nameSimplifiedChinese    string = "米兰达"
	nameTraditionalChinese   string = "米蘭達"
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
