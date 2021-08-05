package midge

import "golang.org/x/text/language"

const (
	nameAmericanEnglish      string = "Midge"
	nameCanadianFrench       string = "Michèle"
	nameDutch                string = "Midge"
	nameFrench               string = "Michèle"
	nameGerman               string = "Anna"
	nameItalian              string = "Minù"
	nameJapanese             string = "うずまき"
	nameLatinAmericanSpanish string = "Paloma"
	nameKorean               string = "핑글이"
	nameRussian              string = "Мидж"
	nameSpanish              string = "Paloma"
	nameSimplifiedChinese    string = "小酒窝"
	nameTraditionalChinese   string = "小酒窩"
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
