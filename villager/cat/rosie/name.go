package rosie

import "golang.org/x/text/language"

const (
	nameAmericanEnglish      string = "Rosie"
	nameCanadianFrench       string = "Rosie"
	nameDutch                string = "Rosie"
	nameFrench               string = "Rosie"
	nameGerman               string = "Sophie"
	nameItalian              string = "Grinfia"
	nameJapanese             string = "ブーケ"
	nameLatinAmericanSpanish string = "Minina"
	nameKorean               string = "부케"
	nameRussian              string = "Рози"
	nameSpanish              string = "Minina"
	nameSimplifiedChinese    string = "彭花"
	nameTraditionalChinese   string = "彭花"
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
