package tad

import "golang.org/x/text/language"

const (
	nameAmericanEnglish      string = "Tad"
	nameCanadianFrench       string = "Rénato"
	nameDutch                string = "Tad"
	nameFrench               string = "Rénato"
	nameGerman               string = "Paul"
	nameItalian              string = "Achille"
	nameJapanese             string = "タンボ"
	nameLatinAmericanSpanish string = "Saltonio"
	nameKorean               string = "텀보"
	nameRussian              string = "Тэд"
	nameSpanish              string = "Saltonio"
	nameSimplifiedChinese    string = "阿田"
	nameTraditionalChinese   string = "阿田"
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
