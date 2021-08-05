package sparro

import "golang.org/x/text/language"

const (
	nameAmericanEnglish      string = "Sparro"
	nameCanadianFrench       string = "Darius"
	nameDutch                string = "Sparro"
	nameFrench               string = "Darius"
	nameGerman               string = "Nestor"
	nameItalian              string = "Piumardo"
	nameJapanese             string = "ちゅんのすけ"
	nameLatinAmericanSpanish string = "Jaime"
	nameKorean               string = "춘섭"
	nameRussian              string = "Спарро"
	nameSpanish              string = "Jaime"
	nameSimplifiedChinese    string = "周之翔"
	nameTraditionalChinese   string = "周之翔"
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
