package deena

import "golang.org/x/text/language"

const (
	nameAmericanEnglish      string = "Deena"
	nameCanadianFrench       string = "Mina"
	nameDutch                string = "Deena"
	nameFrench               string = "Mina"
	nameGerman               string = "Olivia"
	nameItalian              string = "Dina"
	nameJapanese             string = "まりも"
	nameLatinAmericanSpanish string = "Martita"
	nameKorean               string = "마리모"
	nameRussian              string = "Дина"
	nameSpanish              string = "Martita"
	nameSimplifiedChinese    string = "球藻"
	nameTraditionalChinese   string = "毬藻"
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
