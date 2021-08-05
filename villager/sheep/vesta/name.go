package vesta

import "golang.org/x/text/language"

const (
	nameAmericanEnglish      string = "Vesta"
	nameCanadianFrench       string = "Hélaine"
	nameDutch                string = "Vesta"
	nameFrench               string = "Hélaine"
	nameGerman               string = "Dolly"
	nameItalian              string = "Lanella"
	nameJapanese             string = "メリヤス"
	nameLatinAmericanSpanish string = "Vesta"
	nameKorean               string = "메리어스"
	nameRussian              string = "Веста"
	nameSpanish              string = "Vesta"
	nameSimplifiedChinese    string = "绵绵"
	nameTraditionalChinese   string = "綿綿"
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
