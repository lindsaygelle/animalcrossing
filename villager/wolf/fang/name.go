package fang

import "golang.org/x/text/language"

const (
	nameAmericanEnglish      string = "Fang"
	nameCanadianFrench       string = "Pierrot"
	nameDutch                string = "Fang"
	nameFrench               string = "Pierrot"
	nameGerman               string = "Grimm"
	nameItalian              string = "Zanna"
	nameJapanese             string = "シベリア"
	nameLatinAmericanSpanish string = "Colmillo"
	nameKorean               string = "시베리아"
	nameRussian              string = "Фанг"
	nameSpanish              string = "Colmillo"
	nameSimplifiedChinese    string = "史培亚"
	nameTraditionalChinese   string = "史培亞"
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
