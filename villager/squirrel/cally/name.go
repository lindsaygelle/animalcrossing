package cally

import "golang.org/x/text/language"

const (
	nameAmericanEnglish      string = "Cally"
	nameCanadianFrench       string = "Célia"
	nameDutch                string = "Cally"
	nameFrench               string = "Célia"
	nameGerman               string = "Hörnchen"
	nameItalian              string = "Rosa"
	nameJapanese             string = "パセリ"
	nameLatinAmericanSpanish string = "Almendra"
	nameKorean               string = "파슬리"
	nameRussian              string = "Кэлли"
	nameSpanish              string = "Almendra"
	nameSimplifiedChinese    string = "小芹"
	nameTraditionalChinese   string = "小芹"
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
