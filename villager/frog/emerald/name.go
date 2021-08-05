package emerald

import "golang.org/x/text/language"

const (
	nameAmericanEnglish      string = "Emerald"
	nameCanadianFrench       string = ""
	nameDutch                string = ""
	nameFrench               string = "sproiiing"
	nameGerman               string = "boioioing"
	nameItalian              string = "boing"
	nameJapanese             string = "だケロ"
	nameLatinAmericanSpanish string = ""
	nameKorean               string = ""
	nameRussian              string = ""
	nameSpanish              string = "esproing"
	nameSimplifiedChinese    string = "咕咕"
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
