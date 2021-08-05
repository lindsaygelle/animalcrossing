package sherb

import "golang.org/x/text/language"

const (
	nameAmericanEnglish      string = "Sherb"
	nameCanadianFrench       string = "Capri"
	nameDutch                string = "Sherb"
	nameFrench               string = "Capri"
	nameGerman               string = "Morpheus"
	nameItalian              string = "Capraldo"
	nameJapanese             string = "レム"
	nameLatinAmericanSpanish string = "Morfeo"
	nameKorean               string = "래미"
	nameRussian              string = "Шерб"
	nameSpanish              string = "Morfeo"
	nameSimplifiedChinese    string = "雷姆"
	nameTraditionalChinese   string = "雷姆"
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
