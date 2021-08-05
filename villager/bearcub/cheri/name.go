package cheri

import "golang.org/x/text/language"

const (
	nameAmericanEnglish      string = "Cheri"
	nameCanadianFrench       string = "Rosalie"
	nameDutch                string = "Cheri"
	nameFrench               string = "Rosalie"
	nameGerman               string = "Claudia"
	nameItalian              string = "Bonbon"
	nameJapanese             string = "アセロラ"
	nameLatinAmericanSpanish string = "Cerecita"
	nameKorean               string = "아세로라"
	nameRussian              string = "Шери"
	nameSpanish              string = "Cerecita"
	nameSimplifiedChinese    string = "樱桃"
	nameTraditionalChinese   string = "櫻桃"
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
