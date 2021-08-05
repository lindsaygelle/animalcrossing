package yodel

import "golang.org/x/text/language"

const (
	nameAmericanEnglish      string = "Yodel"
	nameCanadianFrench       string = ""
	nameDutch                string = ""
	nameFrench               string = "yolahitoou"
	nameGerman               string = "jodeli"
	nameItalian              string = "ullallà"
	nameJapanese             string = "ヨホホー"
	nameLatinAmericanSpanish string = ""
	nameKorean               string = ""
	nameRussian              string = ""
	nameSpanish              string = "joya"
	nameSimplifiedChinese    string = ""
	nameTraditionalChinese   string = ""
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
