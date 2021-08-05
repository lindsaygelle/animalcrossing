package angus

import "golang.org/x/text/language"

const (
	nameAmericanEnglish      string = "Angus"
	nameCanadianFrench       string = "Angus"
	nameDutch                string = "Angus"
	nameFrench               string = "Angus"
	nameGerman               string = "Angus"
	nameItalian              string = "Angus"
	nameJapanese             string = "セルバンテス"
	nameLatinAmericanSpanish string = "Aliste"
	nameKorean               string = "반데스"
	nameRussian              string = "Ангус"
	nameSpanish              string = "Aliste"
	nameSimplifiedChinese    string = "施万德"
	nameTraditionalChinese   string = "施萬德"
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
