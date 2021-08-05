package flo

import "golang.org/x/text/language"

const (
	nameAmericanEnglish      string = "Flo"
	nameCanadianFrench       string = "Nora"
	nameDutch                string = "Flo"
	nameFrench               string = "Nora"
	nameGerman               string = "Susanne"
	nameItalian              string = "Nives"
	nameJapanese             string = "レイラ"
	nameLatinAmericanSpanish string = "Nieves"
	nameKorean               string = "레이라"
	nameRussian              string = "Фло"
	nameSpanish              string = "Nieves"
	nameSimplifiedChinese    string = "蕾拉"
	nameTraditionalChinese   string = "蕾拉"
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
