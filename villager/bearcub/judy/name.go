package judy

import "golang.org/x/text/language"

const (
	nameAmericanEnglish      string = "Judy"
	nameCanadianFrench       string = "Laura"
	nameDutch                string = "Judy"
	nameFrench               string = "Laura"
	nameGerman               string = "Misuzu"
	nameItalian              string = "Misuzu"
	nameJapanese             string = "みすず"
	nameLatinAmericanSpanish string = "Rosezna"
	nameKorean               string = "미애"
	nameRussian              string = "Джуди"
	nameSpanish              string = "Rosezna"
	nameSimplifiedChinese    string = "美玲"
	nameTraditionalChinese   string = "美玲"
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
