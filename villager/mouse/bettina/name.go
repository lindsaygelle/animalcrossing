package bettina

import "golang.org/x/text/language"

const (
	nameAmericanEnglish      string = "Bettina"
	nameCanadianFrench       string = "Sabrina"
	nameDutch                string = "Bettina"
	nameFrench               string = "Sabrina"
	nameGerman               string = "Schoki"
	nameItalian              string = "Rattina"
	nameJapanese             string = "マルコ"
	nameLatinAmericanSpanish string = "Ottina"
	nameKorean               string = "마르카"
	nameRussian              string = "Беттина"
	nameSpanish              string = "Ottina"
	nameSimplifiedChinese    string = "丸子"
	nameTraditionalChinese   string = "丸子"
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
