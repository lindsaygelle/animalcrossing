package tiffany

import "golang.org/x/text/language"

const (
	nameAmericanEnglish      string = "Tiffany"
	nameCanadianFrench       string = "Tiphaine"
	nameDutch                string = "Tiffany"
	nameFrench               string = "Tiphaine"
	nameGerman               string = "Michelle"
	nameItalian              string = "Stefania"
	nameJapanese             string = "バズレー"
	nameLatinAmericanSpanish string = "Tiffany"
	nameKorean               string = "바슬레"
	nameRussian              string = "Тиффани"
	nameSpanish              string = "Tiffany"
	nameSimplifiedChinese    string = "大姐头"
	nameTraditionalChinese   string = "大姐頭"
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
