package gigi

import "golang.org/x/text/language"

const (
	nameAmericanEnglish      string = "Gigi"
	nameCanadianFrench       string = "Gloria"
	nameDutch                string = "Gigi"
	nameFrench               string = "Gloria"
	nameGerman               string = "Violetta"
	nameItalian              string = "Raganà"
	nameJapanese             string = "リンダ"
	nameLatinAmericanSpanish string = "Cleo"
	nameKorean               string = "린다"
	nameRussian              string = "Джиджи"
	nameSpanish              string = "Cleo"
	nameSimplifiedChinese    string = "林妲"
	nameTraditionalChinese   string = "林妲"
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
