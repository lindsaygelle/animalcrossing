package celia

import "golang.org/x/text/language"

const (
	nameAmericanEnglish      string = "Celia"
	nameCanadianFrench       string = "Garance"
	nameDutch                string = "Celia"
	nameFrench               string = "Garance"
	nameGerman               string = "Lora"
	nameItalian              string = "Celia"
	nameJapanese             string = "ティファニー"
	nameLatinAmericanSpanish string = "Jazmín"
	nameKorean               string = "티파니"
	nameRussian              string = "Селия"
	nameSpanish              string = "Jazmín"
	nameSimplifiedChinese    string = "蒂芬妮"
	nameTraditionalChinese   string = "蒂芬妮"
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
