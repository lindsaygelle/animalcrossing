package erik

import "golang.org/x/text/language"

const (
	nameAmericanEnglish      string = "Erik"
	nameCanadianFrench       string = "Abraham"
	nameDutch                string = "Erik"
	nameFrench               string = "Abraham"
	nameGerman               string = "Erik"
	nameItalian              string = "Cervasio"
	nameJapanese             string = "チャック"
	nameLatinAmericanSpanish string = "Cervasio"
	nameKorean               string = "자끄"
	nameRussian              string = "Эрик"
	nameSpanish              string = "Cervasio"
	nameSimplifiedChinese    string = "查克"
	nameTraditionalChinese   string = "查克"
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
