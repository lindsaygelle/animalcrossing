package egbert

import "golang.org/x/text/language"

const (
	nameAmericanEnglish      string = "Egbert"
	nameCanadianFrench       string = "Herbert"
	nameDutch                string = "Egbert"
	nameFrench               string = "Herbert"
	nameGerman               string = "Waldemar"
	nameItalian              string = "Chicco"
	nameJapanese             string = "しもやけ"
	nameLatinAmericanSpanish string = "Norberto"
	nameKorean               string = "김희"
	nameRussian              string = "Эгберт"
	nameSpanish              string = "Norberto"
	nameSimplifiedChinese    string = "黄金鸡"
	nameTraditionalChinese   string = "黃金雞"
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
