package paolo

import "golang.org/x/text/language"

const (
	nameAmericanEnglish      string = "Paolo"
	nameCanadianFrench       string = "Paolo"
	nameDutch                string = "Paolo"
	nameFrench               string = "Paolo"
	nameGerman               string = "Paolo"
	nameItalian              string = "Paolino"
	nameJapanese             string = "パオロ"
	nameLatinAmericanSpanish string = "Paolo"
	nameKorean               string = "파올로"
	nameRussian              string = "Паоло"
	nameSpanish              string = "Paolo"
	nameSimplifiedChinese    string = "保罗"
	nameTraditionalChinese   string = "保羅"
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
