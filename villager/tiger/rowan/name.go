package rowan

import "golang.org/x/text/language"

const (
	nameAmericanEnglish      string = "Rowan"
	nameCanadianFrench       string = "Marito"
	nameDutch                string = "Rowan"
	nameFrench               string = "Marito"
	nameGerman               string = "Gisbert"
	nameItalian              string = "Ernesto"
	nameJapanese             string = "ゴメス"
	nameLatinAmericanSpanish string = "Miguelón"
	nameKorean               string = "고메스"
	nameRussian              string = "Роуэн"
	nameSpanish              string = "Miguelón"
	nameSimplifiedChinese    string = "戈麦斯"
	nameTraditionalChinese   string = "戈麥斯"
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
