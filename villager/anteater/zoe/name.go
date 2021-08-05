package zoe

import "golang.org/x/text/language"

const (
	nameAmericanEnglish      string = "Zoe"
	nameCanadianFrench       string = ""
	nameDutch                string = ""
	nameFrench               string = "fouf fouf"
	nameGerman               string = "schnieef"
	nameItalian              string = "oiiinfff"
	nameJapanese             string = "だモン"
	nameLatinAmericanSpanish string = ""
	nameKorean               string = ""
	nameRussian              string = ""
	nameSpanish              string = "fufuf"
	nameSimplifiedChinese    string = "没错"
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
