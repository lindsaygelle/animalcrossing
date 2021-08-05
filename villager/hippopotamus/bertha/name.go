package bertha

import "golang.org/x/text/language"

const (
	nameAmericanEnglish      string = "Bertha"
	nameCanadianFrench       string = "Bertha"
	nameDutch                string = "Bertha"
	nameFrench               string = "Bertha"
	nameGerman               string = "Berta"
	nameItalian              string = "Umberta"
	nameJapanese             string = "あんこ"
	nameLatinAmericanSpanish string = "Berta"
	nameKorean               string = "베티"
	nameRussian              string = "Берта"
	nameSpanish              string = "Berta"
	nameSimplifiedChinese    string = "豆沙"
	nameTraditionalChinese   string = "豆沙"
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
