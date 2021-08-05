package pippy

import "golang.org/x/text/language"

const (
	nameAmericanEnglish      string = "Pippy"
	nameCanadianFrench       string = "Nadia"
	nameDutch                string = "Pippy"
	nameFrench               string = "Nadia"
	nameGerman               string = "Lotta"
	nameItalian              string = "Pippi"
	nameJapanese             string = "ロッタ"
	nameLatinAmericanSpanish string = "Méloni"
	nameKorean               string = "로타"
	nameRussian              string = "Пиппи"
	nameSpanish              string = "Méloni"
	nameSimplifiedChinese    string = "罗塔"
	nameTraditionalChinese   string = "羅塔"
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
