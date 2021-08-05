package candi

import "golang.org/x/text/language"

const (
	nameAmericanEnglish      string = "Candi"
	nameCanadianFrench       string = "Sucrette"
	nameDutch                string = "Candi"
	nameFrench               string = "Sucrette"
	nameGerman               string = "Renate"
	nameItalian              string = "Mella"
	nameJapanese             string = "かんゆ"
	nameLatinAmericanSpanish string = "Chuchi"
	nameKorean               string = "사탕"
	nameRussian              string = "Кэнди"
	nameSpanish              string = "Chuchi"
	nameSimplifiedChinese    string = "甘油"
	nameTraditionalChinese   string = "甘油"
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
