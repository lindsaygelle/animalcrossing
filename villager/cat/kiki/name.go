package kiki

import "golang.org/x/text/language"

const (
	nameAmericanEnglish      string = "Kiki"
	nameCanadianFrench       string = "Kiki"
	nameDutch                string = "Kiki"
	nameFrench               string = "Kiki"
	nameGerman               string = "Kiki"
	nameItalian              string = "Kiki"
	nameJapanese             string = "キャビア"
	nameLatinAmericanSpanish string = "Ágata"
	nameKorean               string = "캐비어"
	nameRussian              string = "Кики"
	nameSpanish              string = "Ágata"
	nameSimplifiedChinese    string = "余子酱"
	nameTraditionalChinese   string = "余子醬"
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
