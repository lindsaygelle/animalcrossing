package yuka

import "golang.org/x/text/language"

const (
	nameAmericanEnglish      string = "Yuka"
	nameCanadianFrench       string = "Calypso"
	nameDutch                string = "Yuka"
	nameFrench               string = "Calypso"
	nameGerman               string = "Ute"
	nameItalian              string = "Yuka"
	nameJapanese             string = "ユーカリ"
	nameLatinAmericanSpanish string = "Yuka"
	nameKorean               string = "유카리"
	nameRussian              string = "Юка"
	nameSpanish              string = "Yuka"
	nameSimplifiedChinese    string = "尤嘉莉"
	nameTraditionalChinese   string = "尤嘉莉"
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
