package stitches

import "golang.org/x/text/language"

const (
	nameAmericanEnglish      string = "Stitches"
	nameCanadianFrench       string = "Miro"
	nameDutch                string = "Stitches"
	nameFrench               string = "Miro"
	nameGerman               string = "Berry"
	nameItalian              string = "Toppetta"
	nameJapanese             string = "パッチ"
	nameLatinAmericanSpanish string = "Parches"
	nameKorean               string = "패치"
	nameRussian              string = "Стичес"
	nameSpanish              string = "Parches"
	nameSimplifiedChinese    string = "玩具熊"
	nameTraditionalChinese   string = "玩具熊"
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
