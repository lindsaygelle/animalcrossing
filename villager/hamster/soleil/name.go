package soleil

import "golang.org/x/text/language"

const (
	nameAmericanEnglish      string = "Soleil"
	nameCanadianFrench       string = "Stella"
	nameDutch                string = "Soleil"
	nameFrench               string = "Stella"
	nameGerman               string = "Samira"
	nameItalian              string = "Cettina"
	nameJapanese             string = "シャンティ"
	nameLatinAmericanSpanish string = "Soraya"
	nameKorean               string = "샨티"
	nameRussian              string = "Солей"
	nameSpanish              string = "Soraya"
	nameSimplifiedChinese    string = "安安"
	nameTraditionalChinese   string = "安安"
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
