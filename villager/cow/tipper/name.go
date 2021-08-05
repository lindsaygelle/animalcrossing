package tipper

import "golang.org/x/text/language"

const (
	nameAmericanEnglish      string = "Tipper"
	nameCanadianFrench       string = "Valé"
	nameDutch                string = "Tipper"
	nameFrench               string = "Valé"
	nameGerman               string = "Angela"
	nameItalian              string = "Tipper"
	nameJapanese             string = "まきば"
	nameLatinAmericanSpanish string = "Pinta"
	nameKorean               string = "마틸다"
	nameRussian              string = "Типпер"
	nameSpanish              string = "Pinta"
	nameSimplifiedChinese    string = "牧牧"
	nameTraditionalChinese   string = "牧牧"
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
