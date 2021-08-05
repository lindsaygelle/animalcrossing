package bitty

import "golang.org/x/text/language"

const (
	nameAmericanEnglish      string = "Bitty"
	nameCanadianFrench       string = "Potama"
	nameDutch                string = "Bitty"
	nameFrench               string = "Potama"
	nameGerman               string = "Biggi"
	nameItalian              string = "Rosetta"
	nameJapanese             string = "エーミー"
	nameLatinAmericanSpanish string = "Ofelia"
	nameKorean               string = "비티"
	nameRussian              string = "Битти"
	nameSpanish              string = "Ofelia"
	nameSimplifiedChinese    string = "艾咪"
	nameTraditionalChinese   string = "艾咪"
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
