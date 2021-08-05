package lyman

import "golang.org/x/text/language"

const (
	nameAmericanEnglish      string = "Lyman"
	nameCanadianFrench       string = "Kalyptus"
	nameDutch                string = "Lyman"
	nameFrench               string = "Kalyptus"
	nameGerman               string = "Pepe"
	nameItalian              string = "Nicola"
	nameJapanese             string = "オズモンド"
	nameLatinAmericanSpanish string = "Chipi"
	nameKorean               string = "오즈먼드"
	nameRussian              string = "Лайман"
	nameSpanish              string = "Chipi"
	nameSimplifiedChinese    string = "敖志明"
	nameTraditionalChinese   string = "敖志明"
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
