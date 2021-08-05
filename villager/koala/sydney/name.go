package sydney

import "golang.org/x/text/language"

const (
	nameAmericanEnglish      string = "Sydney"
	nameCanadianFrench       string = "Koaline"
	nameDutch                string = "Sydney"
	nameFrench               string = "Koaline"
	nameGerman               string = "Silke"
	nameItalian              string = "Sidney"
	nameJapanese             string = "シドニー"
	nameLatinAmericanSpanish string = "Sidney"
	nameKorean               string = "시드니"
	nameRussian              string = "Сидни"
	nameSpanish              string = "Sidney"
	nameSimplifiedChinese    string = "思颖"
	nameTraditionalChinese   string = "思穎"
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
