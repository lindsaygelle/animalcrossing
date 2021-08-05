package kyle

import "golang.org/x/text/language"

const (
	nameAmericanEnglish      string = "Kyle"
	nameCanadianFrench       string = "Gary"
	nameDutch                string = "Kyle"
	nameFrench               string = "Gary"
	nameGerman               string = "Wolfgang"
	nameItalian              string = "Ululo"
	nameJapanese             string = "リカルド"
	nameLatinAmericanSpanish string = "Ataúlfo"
	nameKorean               string = "리카르도"
	nameRussian              string = "Кайл"
	nameSpanish              string = "Ataúlfo"
	nameSimplifiedChinese    string = "李可"
	nameTraditionalChinese   string = "李可"
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
