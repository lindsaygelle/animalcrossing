package medli

import "golang.org/x/text/language"

const (
	nameAmericanEnglish      string = "Medli"
	nameCanadianFrench       string = "Médolie"
	nameDutch                string = ""
	nameFrench               string = "Médolie"
	nameGerman               string = "Medolie"
	nameItalian              string = "Famirè"
	nameJapanese             string = "メドリ"
	nameLatinAmericanSpanish string = "Medli"
	nameKorean               string = "메들리"
	nameRussian              string = ""
	nameSpanish              string = "Medli"
	nameSimplifiedChinese    string = ""
	nameTraditionalChinese   string = ""
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
