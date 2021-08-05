package nosegay

import "golang.org/x/text/language"

const (
	nameAmericanEnglish      string = "Nosegay"
	nameCanadianFrench       string = ""
	nameDutch                string = ""
	nameFrench               string = "hoooonk"
	nameGerman               string = "schnauf"
	nameItalian              string = "hoooonk"
	nameJapanese             string = "でアリ"
	nameLatinAmericanSpanish string = ""
	nameKorean               string = ""
	nameRussian              string = ""
	nameSpanish              string = "achís"
	nameSimplifiedChinese    string = "确实"
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
