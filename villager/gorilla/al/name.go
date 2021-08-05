package al

import "golang.org/x/text/language"

const (
	nameAmericanEnglish      string = "Al"
	nameCanadianFrench       string = "Gustave"
	nameDutch                string = "Al"
	nameFrench               string = "Gustave"
	nameGerman               string = "Kokong"
	nameItalian              string = "Gregorio"
	nameJapanese             string = "たもつ"
	nameLatinAmericanSpanish string = "Álex"
	nameKorean               string = "우락"
	nameRussian              string = "Эл"
	nameSpanish              string = "Álex"
	nameSimplifiedChinese    string = "阿保"
	nameTraditionalChinese   string = "阿保"
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
