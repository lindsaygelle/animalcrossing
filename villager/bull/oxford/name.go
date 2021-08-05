package oxford

import "golang.org/x/text/language"

const (
	nameAmericanEnglish      string = "Oxford"
	nameCanadianFrench       string = ""
	nameDutch                string = ""
	nameFrench               string = "la vache"
	nameGerman               string = "bulli, hä"
	nameItalian              string = "muuh"
	nameJapanese             string = "でんがな"
	nameLatinAmericanSpanish string = ""
	nameKorean               string = ""
	nameRussian              string = ""
	nameSpanish              string = "torito"
	nameSimplifiedChinese    string = "可不"
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
