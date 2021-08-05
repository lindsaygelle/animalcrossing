package carmen

import "golang.org/x/text/language"

const (
	nameAmericanEnglish      string = "Carmen"
	nameCanadianFrench       string = "Zoé"
	nameDutch                string = "Carmen"
	nameFrench               string = "Zoé"
	nameGerman               string = "Hilda"
	nameItalian              string = "Conny"
	nameJapanese             string = "チョコ"
	nameLatinAmericanSpanish string = "Linda"
	nameKorean               string = "초코"
	nameRussian              string = "Кармен"
	nameSpanish              string = "Linda"
	nameSimplifiedChinese    string = "巧蔻"
	nameTraditionalChinese   string = "巧蔻"
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
