package muffy

import "golang.org/x/text/language"

const (
	nameAmericanEnglish      string = "Muffy"
	nameCanadianFrench       string = "Charlène"
	nameDutch                string = "Muffy"
	nameFrench               string = "Charlène"
	nameGerman               string = "Marion"
	nameItalian              string = "Morena"
	nameJapanese             string = "フリル"
	nameLatinAmericanSpanish string = "Muaré"
	nameKorean               string = "프릴"
	nameRussian              string = "Маффи"
	nameSpanish              string = "Muaré"
	nameSimplifiedChinese    string = "彭澎"
	nameTraditionalChinese   string = "彭澎"
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
