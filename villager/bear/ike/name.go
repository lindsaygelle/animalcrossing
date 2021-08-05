package ike

import "golang.org/x/text/language"

const (
	nameAmericanEnglish      string = "Ike"
	nameCanadianFrench       string = "Isaac"
	nameDutch                string = "Ike"
	nameFrench               string = "Isaac"
	nameGerman               string = "Ike"
	nameItalian              string = "Ortensio"
	nameJapanese             string = "ダイク"
	nameLatinAmericanSpanish string = "Astillas"
	nameKorean               string = "대공"
	nameRussian              string = "Айк"
	nameSpanish              string = "Astillas"
	nameSimplifiedChinese    string = "大功"
	nameTraditionalChinese   string = "大功"
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
