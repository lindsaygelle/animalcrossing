package truffles

import "golang.org/x/text/language"

const (
	nameAmericanEnglish      string = "Truffles"
	nameCanadianFrench       string = "Trufa"
	nameDutch                string = "Truffles"
	nameFrench               string = "Trufa"
	nameGerman               string = "Luzie"
	nameItalian              string = "Grufetta"
	nameJapanese             string = "トンコ"
	nameLatinAmericanSpanish string = "Trufas"
	nameKorean               string = "탱고"
	nameRussian              string = "Трафлс"
	nameSpanish              string = "Trufas"
	nameSimplifiedChinese    string = "嘟嘟"
	nameTraditionalChinese   string = "嘟嘟"
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
