package tybalt

import "golang.org/x/text/language"

const (
	nameAmericanEnglish      string = "Tybalt"
	nameCanadianFrench       string = "Jeff"
	nameDutch                string = "Tybalt"
	nameFrench               string = "Jeff"
	nameGerman               string = "Arne"
	nameItalian              string = "Ubaldo"
	nameJapanese             string = "ハリマオ"
	nameLatinAmericanSpanish string = "Teobaldo"
	nameKorean               string = "티볼트"
	nameRussian              string = "Тибальт"
	nameSpanish              string = "Teobaldo"
	nameSimplifiedChinese    string = "谷丰"
	nameTraditionalChinese   string = "谷豐"
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
