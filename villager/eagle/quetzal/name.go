package quetzal

import "golang.org/x/text/language"

const (
	nameAmericanEnglish      string = "Quetzal"
	nameCanadianFrench       string = ""
	nameDutch                string = ""
	nameFrench               string = "couêêêtzal"
	nameGerman               string = "KRIKRI"
	nameItalian              string = "OCCHIO"
	nameJapanese             string = "ゲロッパ"
	nameLatinAmericanSpanish string = ""
	nameKorean               string = ""
	nameRussian              string = ""
	nameSpanish              string = "cuiii"
	nameSimplifiedChinese    string = "喳喳"
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
