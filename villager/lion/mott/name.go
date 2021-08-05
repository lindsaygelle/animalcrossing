package mott

import "golang.org/x/text/language"

const (
	nameAmericanEnglish      string = "Mott"
	nameCanadianFrench       string = "Aimé"
	nameDutch                string = "Mott"
	nameFrench               string = "Aimé"
	nameGerman               string = "Leonhard"
	nameItalian              string = "Leopold"
	nameJapanese             string = "リック"
	nameLatinAmericanSpanish string = "Jones"
	nameKorean               string = "릭"
	nameRussian              string = "Мотт"
	nameSpanish              string = "Jones"
	nameSimplifiedChinese    string = "李克"
	nameTraditionalChinese   string = "李克"
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
