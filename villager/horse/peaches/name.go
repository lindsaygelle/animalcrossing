package peaches

import "golang.org/x/text/language"

const (
	nameAmericanEnglish      string = "Peaches"
	nameCanadianFrench       string = "Prune"
	nameDutch                string = "Peaches"
	nameFrench               string = "Prune"
	nameGerman               string = "Claire"
	nameItalian              string = "Ronzina"
	nameJapanese             string = "ドサコ"
	nameLatinAmericanSpanish string = "Perla"
	nameKorean               string = "말자"
	nameRussian              string = "Пичес"
	nameSpanish              string = "Perla"
	nameSimplifiedChinese    string = "小薰"
	nameTraditionalChinese   string = "小薰"
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
