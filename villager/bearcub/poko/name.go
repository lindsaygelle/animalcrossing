package poko

import "golang.org/x/text/language"

const (
	nameAmericanEnglish      string = "Poko"
	nameCanadianFrench       string = ""
	nameDutch                string = ""
	nameFrench               string = ""
	nameGerman               string = ""
	nameItalian              string = ""
	nameJapanese             string = "へへ"
	nameLatinAmericanSpanish string = ""
	nameKorean               string = ""
	nameRussian              string = ""
	nameSpanish              string = ""
	nameSimplifiedChinese    string = ""
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
