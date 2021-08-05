package gala

import "golang.org/x/text/language"

const (
	nameAmericanEnglish      string = "Gala"
	nameCanadianFrench       string = "Camille"
	nameDutch                string = "Gala"
	nameFrench               string = "Camille"
	nameGerman               string = "Oinka"
	nameItalian              string = "Lisetta"
	nameJapanese             string = "ためこ"
	nameLatinAmericanSpanish string = "Marita"
	nameKorean               string = "꽃지"
	nameRussian              string = "Гала"
	nameSpanish              string = "Marita"
	nameSimplifiedChinese    string = "小芽"
	nameTraditionalChinese   string = "小芽"
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
