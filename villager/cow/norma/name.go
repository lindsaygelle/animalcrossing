package norma

import "golang.org/x/text/language"

const (
	nameAmericanEnglish      string = "Norma"
	nameCanadianFrench       string = "Norma"
	nameDutch                string = "Norma"
	nameFrench               string = "Norma"
	nameGerman               string = "Nelly"
	nameItalian              string = "Norma"
	nameJapanese             string = "いさこ"
	nameLatinAmericanSpanish string = "Norma"
	nameKorean               string = "미자"
	nameRussian              string = "Норма"
	nameSpanish              string = "Norma"
	nameSimplifiedChinese    string = "晨曦"
	nameTraditionalChinese   string = "晨曦"
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
