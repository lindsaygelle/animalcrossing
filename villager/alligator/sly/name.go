package sly

import "golang.org/x/text/language"

const (
	nameAmericanEnglish      string = "Sly"
	nameCanadianFrench       string = "Chuck"
	nameDutch                string = "Sly"
	nameFrench               string = "Chuck"
	nameGerman               string = "Steve"
	nameItalian              string = "Driberto"
	nameJapanese             string = "ハイド"
	nameLatinAmericanSpanish string = "Estallón"
	nameKorean               string = "하이드"
	nameRussian              string = "Слай"
	nameSpanish              string = "Estallón"
	nameSimplifiedChinese    string = "海德"
	nameTraditionalChinese   string = "海德"
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
