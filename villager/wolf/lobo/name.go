package lobo

import "golang.org/x/text/language"

const (
	nameAmericanEnglish      string = "Lobo"
	nameCanadianFrench       string = "Lobo"
	nameDutch                string = "Lobo"
	nameFrench               string = "Lobo"
	nameGerman               string = "Lupo"
	nameItalian              string = "Lupen"
	nameJapanese             string = "ブンジロウ"
	nameLatinAmericanSpanish string = "Lupo"
	nameKorean               string = "늑태"
	nameRussian              string = "Лобо"
	nameSpanish              string = "Lupo"
	nameSimplifiedChinese    string = "潘奇隆"
	nameTraditionalChinese   string = "潘奇隆"
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
