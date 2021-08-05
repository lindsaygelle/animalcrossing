package boyd

import "golang.org/x/text/language"

const (
	nameAmericanEnglish      string = "Boyd"
	nameCanadianFrench       string = "Primo"
	nameDutch                string = "Boyd"
	nameFrench               string = "Primo"
	nameGerman               string = "Boyd"
	nameItalian              string = "Brando"
	nameJapanese             string = "ボイド"
	nameLatinAmericanSpanish string = "Bunga"
	nameKorean               string = "보이드"
	nameRussian              string = "Бойд"
	nameSpanish              string = "Bunga"
	nameSimplifiedChinese    string = "空空"
	nameTraditionalChinese   string = "空空"
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
