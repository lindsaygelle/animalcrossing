package cyrano

import "golang.org/x/text/language"

const (
	nameAmericanEnglish      string = "Cyrano"
	nameCanadianFrench       string = "Cyrano"
	nameDutch                string = "Cyrano"
	nameFrench               string = "Cyrano"
	nameGerman               string = "Theo"
	nameItalian              string = "Cirano"
	nameJapanese             string = "さくらじま"
	nameLatinAmericanSpanish string = "Cirano"
	nameKorean               string = "사지마"
	nameRussian              string = "Сирано"
	nameSpanish              string = "Cirano"
	nameSimplifiedChinese    string = "阳明"
	nameTraditionalChinese   string = "陽明"
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
