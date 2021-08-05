package ava

import "golang.org/x/text/language"

const (
	nameAmericanEnglish      string = "Ava"
	nameCanadianFrench       string = "Eva"
	nameDutch                string = "Ava"
	nameFrench               string = "Eva"
	nameGerman               string = "Gisela"
	nameItalian              string = "Ava"
	nameJapanese             string = "ドミグラ"
	nameLatinAmericanSpanish string = "Ava"
	nameKorean               string = "에바"
	nameRussian              string = "Ава"
	nameSpanish              string = "Ava"
	nameSimplifiedChinese    string = "陶米咕"
	nameTraditionalChinese   string = "陶米咕"
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
