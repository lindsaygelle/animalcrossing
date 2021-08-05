package raddle

import "golang.org/x/text/language"

const (
	nameAmericanEnglish      string = "Raddle"
	nameCanadianFrench       string = "Fabien"
	nameDutch                string = "Raddle"
	nameFrench               string = "Fabien"
	nameGerman               string = "Sani"
	nameItalian              string = "Massimo"
	nameJapanese             string = "カックン"
	nameLatinAmericanSpanish string = "Radiolo"
	nameKorean               string = "개군"
	nameRussian              string = "Рэддл"
	nameSpanish              string = "Radiolo"
	nameSimplifiedChinese    string = "嘉俊"
	nameTraditionalChinese   string = "嘉俊"
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
