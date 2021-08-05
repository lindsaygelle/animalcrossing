package goose

import "golang.org/x/text/language"

const (
	nameAmericanEnglish      string = "Goose"
	nameCanadianFrench       string = "Pouli"
	nameDutch                string = "Goose"
	nameFrench               string = "Pouli"
	nameGerman               string = "Konrad"
	nameItalian              string = "Chiricò"
	nameJapanese             string = "ケンタ"
	nameLatinAmericanSpanish string = "Gus"
	nameKorean               string = "건태"
	nameRussian              string = "Гус"
	nameSpanish              string = "Gus"
	nameSimplifiedChinese    string = "肯肯"
	nameTraditionalChinese   string = "肯肯"
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
