package puck

import "golang.org/x/text/language"

const (
	nameAmericanEnglish      string = "Puck"
	nameCanadianFrench       string = "Hervé"
	nameDutch                string = "Puck"
	nameFrench               string = "Hervé"
	nameGerman               string = "Puck"
	nameItalian              string = "Glacido"
	nameJapanese             string = "ホッケー"
	nameLatinAmericanSpanish string = "Fredo"
	nameKorean               string = "하키"
	nameRussian              string = "Пак"
	nameSpanish              string = "Fredo"
	nameSimplifiedChinese    string = "哈奇"
	nameTraditionalChinese   string = "哈奇"
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
