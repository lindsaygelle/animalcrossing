package megan

import "golang.org/x/text/language"

const (
	nameAmericanEnglish      string = "Megan"
	nameCanadianFrench       string = "Candy"
	nameDutch                string = "Megan"
	nameFrench               string = "Candy"
	nameGerman               string = "Dagmar"
	nameItalian              string = "Dolcinia"
	nameJapanese             string = "キャンディ"
	nameLatinAmericanSpanish string = "Carmela"
	nameKorean               string = "캔디"
	nameRussian              string = "Меган"
	nameSpanish              string = "Carmela"
	nameSimplifiedChinese    string = "糖果"
	nameTraditionalChinese   string = "糖果"
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
