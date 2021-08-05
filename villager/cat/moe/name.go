package moe

import "golang.org/x/text/language"

const (
	nameAmericanEnglish      string = "Moe"
	nameCanadianFrench       string = "Momo"
	nameDutch                string = "Moe"
	nameFrench               string = "Momo"
	nameGerman               string = "Tristan"
	nameItalian              string = "Moe"
	nameJapanese             string = "ジンペイ"
	nameLatinAmericanSpanish string = "Tristán"
	nameKorean               string = "진상"
	nameRussian              string = "Моэ"
	nameSpanish              string = "Tristán"
	nameSimplifiedChinese    string = "仁平"
	nameTraditionalChinese   string = "仁平"
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
