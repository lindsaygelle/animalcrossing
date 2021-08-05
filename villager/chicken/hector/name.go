package hector

import "golang.org/x/text/language"

const (
	nameAmericanEnglish      string = "Hector"
	nameCanadianFrench       string = ""
	nameDutch                string = ""
	nameFrench               string = "mon coco"
	nameGerman               string = "kikiriii"
	nameItalian              string = "chirichì"
	nameJapanese             string = "なのだ"
	nameLatinAmericanSpanish string = ""
	nameKorean               string = ""
	nameRussian              string = ""
	nameSpanish              string = "yuu-ju"
	nameSimplifiedChinese    string = "叮叮"
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
