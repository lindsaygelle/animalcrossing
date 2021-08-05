package patty

import "golang.org/x/text/language"

const (
	nameAmericanEnglish      string = "Patty"
	nameCanadianFrench       string = "Margaux"
	nameDutch                string = "Patty"
	nameFrench               string = "Margaux"
	nameGerman               string = "Patricia"
	nameItalian              string = "Patty"
	nameJapanese             string = "カルピ"
	nameLatinAmericanSpanish string = "Vacarena"
	nameKorean               string = "밀크"
	nameRussian              string = "Пэтти"
	nameSpanish              string = "Vacarena"
	nameSimplifiedChinese    string = "阿排"
	nameTraditionalChinese   string = "阿排"
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
