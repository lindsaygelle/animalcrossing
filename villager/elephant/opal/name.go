package opal

import "golang.org/x/text/language"

const (
	nameAmericanEnglish      string = "Opal"
	nameCanadianFrench       string = "Opaline"
	nameDutch                string = "Opal"
	nameFrench               string = "Opaline"
	nameGerman               string = "Olga"
	nameItalian              string = "Opal"
	nameJapanese             string = "オパール"
	nameLatinAmericanSpanish string = "Ópalo"
	nameKorean               string = "오팔"
	nameRussian              string = "Опал"
	nameSpanish              string = "Ópalo"
	nameSimplifiedChinese    string = "露露"
	nameTraditionalChinese   string = "露露"
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
