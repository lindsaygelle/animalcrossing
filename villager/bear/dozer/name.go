package dozer

import "golang.org/x/text/language"

const (
	nameAmericanEnglish      string = "Dozer"
	nameCanadianFrench       string = ""
	nameDutch                string = ""
	nameFrench               string = "zzzzzz"
	nameGerman               string = "schnarch"
	nameItalian              string = "zzzzzz"
	nameJapanese             string = "でアリ"
	nameLatinAmericanSpanish string = ""
	nameKorean               string = ""
	nameRussian              string = ""
	nameSpanish              string = "zzzzzz"
	nameSimplifiedChinese    string = "咕"
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
