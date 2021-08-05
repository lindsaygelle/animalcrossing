package ricky

import "golang.org/x/text/language"

const (
	nameAmericanEnglish      string = "Ricky"
	nameCanadianFrench       string = "Rocky"
	nameDutch                string = "Ricky"
	nameFrench               string = "Rocky"
	nameGerman               string = "Ronny"
	nameItalian              string = "Sberlo"
	nameJapanese             string = "カジロウ"
	nameLatinAmericanSpanish string = "Altramuz"
	nameKorean               string = "갈가리"
	nameRussian              string = "Рики"
	nameSpanish              string = "Altramuz"
	nameSimplifiedChinese    string = "嘉治"
	nameTraditionalChinese   string = "嘉治"
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
