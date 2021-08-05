package scoot

import "golang.org/x/text/language"

const (
	nameAmericanEnglish      string = "Scoot"
	nameCanadianFrench       string = "Scooter"
	nameDutch                string = "Scoot"
	nameFrench               string = "Scooter"
	nameGerman               string = "Helmut"
	nameItalian              string = "Guizzo"
	nameJapanese             string = "マモル"
	nameLatinAmericanSpanish string = "Chema"
	nameKorean               string = "지키미"
	nameRussian              string = "Скут"
	nameSpanish              string = "Chema"
	nameSimplifiedChinese    string = "阿守"
	nameTraditionalChinese   string = "阿守"
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
