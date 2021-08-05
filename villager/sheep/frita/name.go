package frita

import "golang.org/x/text/language"

const (
	nameAmericanEnglish      string = "Frita"
	nameCanadianFrench       string = "Clarabêl"
	nameDutch                string = "Frita"
	nameFrench               string = "Clarabêl"
	nameGerman               string = "Martina"
	nameItalian              string = "Landa"
	nameJapanese             string = "ウェンディ"
	nameLatinAmericanSpanish string = "Ovinia"
	nameKorean               string = "웬디"
	nameRussian              string = "Фрита"
	nameSpanish              string = "Ovinia"
	nameSimplifiedChinese    string = "温蒂"
	nameTraditionalChinese   string = "溫蒂"
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
