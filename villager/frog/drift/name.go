package drift

import "golang.org/x/text/language"

const (
	nameAmericanEnglish      string = "Drift"
	nameCanadianFrench       string = "Gordon"
	nameDutch                string = "Drift"
	nameFrench               string = "Gordon"
	nameGerman               string = "Caspar"
	nameItalian              string = "Prospero"
	nameJapanese             string = "ドク"
	nameLatinAmericanSpanish string = "Surfín"
	nameKorean               string = "덕"
	nameRussian              string = "Дрифт"
	nameSpanish              string = "Surfín"
	nameSimplifiedChinese    string = "毒仔"
	nameTraditionalChinese   string = "毒仔"
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
