package curt

import "golang.org/x/text/language"

const (
	nameAmericanEnglish      string = "Curt"
	nameCanadianFrench       string = "Curt"
	nameDutch                string = "Curt"
	nameFrench               string = "Curt"
	nameGerman               string = "Kurt"
	nameItalian              string = "Curt"
	nameJapanese             string = "ガンテツ"
	nameLatinAmericanSpanish string = "Gorbaché"
	nameKorean               string = "뚝심"
	nameRussian              string = "Керт"
	nameSpanish              string = "Gorbaché"
	nameSimplifiedChinese    string = "铁熊"
	nameTraditionalChinese   string = "鐵熊"
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
