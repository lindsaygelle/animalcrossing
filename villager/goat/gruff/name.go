package gruff

import "golang.org/x/text/language"

const (
	nameAmericanEnglish      string = "Gruff"
	nameCanadianFrench       string = "Grognon"
	nameDutch                string = "Gruff"
	nameFrench               string = "Grognon"
	nameGerman               string = "Gregor"
	nameItalian              string = "Capriolé"
	nameJapanese             string = "ビリー"
	nameLatinAmericanSpanish string = "Sátiro"
	nameKorean               string = "빌리"
	nameRussian              string = "Графф"
	nameSpanish              string = "Sátiro"
	nameSimplifiedChinese    string = "比利"
	nameTraditionalChinese   string = "比利"
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
