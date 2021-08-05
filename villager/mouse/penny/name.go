package penny

import "golang.org/x/text/language"

const (
	nameAmericanEnglish      string = "Penny"
	nameCanadianFrench       string = "oh pétard!"
	nameDutch                string = ""
	nameFrench               string = "oh pétard!"
	nameGerman               string = "piiquie"
	nameItalian              string = "Unknown"
	nameJapanese             string = "なのさ"
	nameLatinAmericanSpanish string = "carambita"
	nameKorean               string = ""
	nameRussian              string = ""
	nameSpanish              string = "carambita"
	nameSimplifiedChinese    string = "Unknown"
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
