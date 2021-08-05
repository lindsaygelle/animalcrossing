package robin

import "golang.org/x/text/language"

const (
	nameAmericanEnglish      string = "Robin"
	nameCanadianFrench       string = "Robie"
	nameDutch                string = "Robin"
	nameFrench               string = "Robie"
	nameGerman               string = "Jule"
	nameItalian              string = "Rossana"
	nameJapanese             string = "パーチク"
	nameLatinAmericanSpanish string = "Aria"
	nameKorean               string = "파틱"
	nameRussian              string = "Робин"
	nameSpanish              string = "Aria"
	nameSimplifiedChinese    string = "喳喳"
	nameTraditionalChinese   string = "喳喳"
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
