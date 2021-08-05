package julia

import "golang.org/x/text/language"

const (
	nameAmericanEnglish      string = "Julia"
	nameCanadianFrench       string = "Julie"
	nameDutch                string = "Julia"
	nameFrench               string = "Julie"
	nameGerman               string = "Julia"
	nameItalian              string = "Giulia"
	nameJapanese             string = "ジュリア"
	nameLatinAmericanSpanish string = "Julia"
	nameKorean               string = "줄리아"
	nameRussian              string = "Джулия"
	nameSpanish              string = "Julia"
	nameSimplifiedChinese    string = "朱莉亚"
	nameTraditionalChinese   string = "朱莉亞"
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
