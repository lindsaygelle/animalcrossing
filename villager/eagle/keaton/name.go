package keaton

import "golang.org/x/text/language"

const (
	nameAmericanEnglish      string = "Keaton"
	nameCanadianFrench       string = "Enzo"
	nameDutch                string = "Keaton"
	nameFrench               string = "Enzo"
	nameGerman               string = "Kai"
	nameItalian              string = "Aquilio"
	nameJapanese             string = "フランク"
	nameLatinAmericanSpanish string = "Lucho"
	nameKorean               string = "프랭크"
	nameRussian              string = "Китон"
	nameSpanish              string = "Lucho"
	nameSimplifiedChinese    string = "法兰克"
	nameTraditionalChinese   string = "法蘭克"
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
