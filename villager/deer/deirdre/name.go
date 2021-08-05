package deirdre

import "golang.org/x/text/language"

const (
	nameAmericanEnglish      string = "Deirdre"
	nameCanadianFrench       string = "Bichoune"
	nameDutch                string = "Deirdre"
	nameFrench               string = "Bichoune"
	nameGerman               string = "Dina"
	nameItalian              string = "Daria"
	nameJapanese             string = "ナディア"
	nameLatinAmericanSpanish string = "Venada"
	nameKorean               string = "나디아"
	nameRussian              string = "Дейрдре"
	nameSpanish              string = "Venada"
	nameSimplifiedChinese    string = "娣雅"
	nameTraditionalChinese   string = "娣雅"
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
