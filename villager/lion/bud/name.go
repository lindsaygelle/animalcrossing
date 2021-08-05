package bud

import "golang.org/x/text/language"

const (
	nameAmericanEnglish      string = "Bud"
	nameCanadianFrench       string = "Léonard"
	nameDutch                string = "Bud"
	nameFrench               string = "Léonard"
	nameGerman               string = "Dieter"
	nameItalian              string = "Leo"
	nameJapanese             string = "グラさん"
	nameLatinAmericanSpanish string = "Surfleo"
	nameKorean               string = "선글"
	nameRussian              string = "Бад"
	nameSpanish              string = "Surfleo"
	nameSimplifiedChinese    string = "莫敬"
	nameTraditionalChinese   string = "莫敬"
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
