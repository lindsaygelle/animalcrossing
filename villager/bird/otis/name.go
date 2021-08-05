package otis

import "golang.org/x/text/language"

const (
	nameAmericanEnglish      string = "Otis"
	nameCanadianFrench       string = ""
	nameDutch                string = ""
	nameFrench               string = "j'pense"
	nameGerman               string = "denk ich"
	nameItalian              string = "tuitcì"
	nameJapanese             string = "なのじゃ"
	nameLatinAmericanSpanish string = ""
	nameKorean               string = ""
	nameRussian              string = ""
	nameSpanish              string = "digo"
	nameSimplifiedChinese    string = "神说"
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
