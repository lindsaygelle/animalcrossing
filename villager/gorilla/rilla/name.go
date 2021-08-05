package rilla

import "golang.org/x/text/language"

const (
	nameAmericanEnglish      string = "Rilla"
	nameCanadianFrench       string = "Rilla"
	nameDutch                string = "Rilla"
	nameFrench               string = "Rilla"
	nameGerman               string = "Rilla"
	nameItalian              string = "Rilla"
	nameJapanese             string = "リラ"
	nameLatinAmericanSpanish string = "Rila"
	nameKorean               string = "릴라"
	nameRussian              string = "Рилла"
	nameSpanish              string = "Rila"
	nameSimplifiedChinese    string = "莉拉"
	nameTraditionalChinese   string = "莉拉"
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
