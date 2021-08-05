package fuchsia

import "golang.org/x/text/language"

const (
	nameAmericanEnglish      string = "Fuchsia"
	nameCanadianFrench       string = "Rosanne"
	nameDutch                string = "Fuchsia"
	nameFrench               string = "Rosanne"
	nameGerman               string = "Selina"
	nameItalian              string = "Fuxia"
	nameJapanese             string = "ジェシカ"
	nameLatinAmericanSpanish string = "Rosalina"
	nameKorean               string = "제시카"
	nameRussian              string = "Фуксия"
	nameSpanish              string = "Rosalina"
	nameSimplifiedChinese    string = "洁西卡"
	nameTraditionalChinese   string = "潔西卡"
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
