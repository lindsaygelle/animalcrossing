package cece

import "golang.org/x/text/language"

const (
	nameAmericanEnglish      string = "Cece"
	nameCanadianFrench       string = "Kala"
	nameDutch                string = ""
	nameFrench               string = "Kala"
	nameGerman               string = "A-Chan"
	nameItalian              string = "Cecilia"
	nameJapanese             string = "なぎさ"
	nameLatinAmericanSpanish string = "Arduna"
	nameKorean               string = "나기사"
	nameRussian              string = ""
	nameSpanish              string = "Arduna"
	nameSimplifiedChinese    string = ""
	nameTraditionalChinese   string = ""
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
