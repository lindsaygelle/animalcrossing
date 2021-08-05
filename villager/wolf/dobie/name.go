package dobie

import "golang.org/x/text/language"

const (
	nameAmericanEnglish      string = "Dobie"
	nameCanadianFrench       string = "Loupiot"
	nameDutch                string = "Dobie"
	nameFrench               string = "Loupiot"
	nameGerman               string = "Sigmund"
	nameItalian              string = "Sigmund"
	nameJapanese             string = "けん"
	nameLatinAmericanSpanish string = "Germán"
	nameKorean               string = "켄"
	nameRussian              string = "Доби"
	nameSpanish              string = "Germán"
	nameSimplifiedChinese    string = "阿研"
	nameTraditionalChinese   string = "阿研"
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
