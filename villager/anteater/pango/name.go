package pango

import "golang.org/x/text/language"

const (
	nameAmericanEnglish      string = "Pango"
	nameCanadianFrench       string = "Mathilda"
	nameDutch                string = "Pango"
	nameFrench               string = "Mathilda"
	nameGerman               string = "Mathilda"
	nameItalian              string = "Carlotta"
	nameJapanese             string = "パトラ"
	nameLatinAmericanSpanish string = "Aspidora"
	nameKorean               string = "패트라"
	nameRussian              string = "Панго"
	nameSpanish              string = "Aspidora"
	nameSimplifiedChinese    string = "佩希"
	nameTraditionalChinese   string = "佩希"
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
