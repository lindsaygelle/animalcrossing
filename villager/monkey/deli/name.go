package deli

import "golang.org/x/text/language"

const (
	nameAmericanEnglish      string = "Deli"
	nameCanadianFrench       string = "Magogo"
	nameDutch                string = "Deli"
	nameFrench               string = "Magogo"
	nameGerman               string = "Anton"
	nameItalian              string = "Bobo"
	nameJapanese             string = "デリー"
	nameLatinAmericanSpanish string = "Cana"
	nameKorean               string = "델리"
	nameRussian              string = "Дели"
	nameSpanish              string = "Cana"
	nameSimplifiedChinese    string = "德里"
	nameTraditionalChinese   string = "德里"
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
