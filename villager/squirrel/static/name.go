package static

import "golang.org/x/text/language"

const (
	nameAmericanEnglish      string = "Static"
	nameCanadianFrench       string = "Électro"
	nameDutch                string = "Static"
	nameFrench               string = "Électro"
	nameGerman               string = "Rudolf"
	nameItalian              string = "Protone"
	nameJapanese             string = "スパーク"
	nameLatinAmericanSpanish string = "Protón"
	nameKorean               string = "스파크"
	nameRussian              string = "Статик"
	nameSpanish              string = "Protón"
	nameSimplifiedChinese    string = "闪电"
	nameTraditionalChinese   string = "閃電"
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
