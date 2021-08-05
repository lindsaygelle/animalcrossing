package spork

import "golang.org/x/text/language"

const (
	nameAmericanEnglish      string = "Spork"
	nameCanadianFrench       string = "Justin"
	nameDutch                string = "Crackle"
	nameFrench               string = "Justin"
	nameGerman               string = "Schwarte"
	nameItalian              string = "Cicciolo"
	nameJapanese             string = "ポーク"
	nameLatinAmericanSpanish string = "Espork"
	nameKorean               string = "포크"
	nameRussian              string = "Крэкл"
	nameSpanish              string = "Espork"
	nameSimplifiedChinese    string = "肉肉"
	nameTraditionalChinese   string = "肉肉"
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
