package jitters

import "golang.org/x/text/language"

const (
	nameAmericanEnglish      string = "Jitters"
	nameCanadianFrench       string = "Gilbert"
	nameDutch                string = "Jitters"
	nameFrench               string = "Gilbert"
	nameGerman               string = "Alex"
	nameItalian              string = "Brivido"
	nameJapanese             string = "ジーニョ"
	nameLatinAmericanSpanish string = "Camelio"
	nameKorean               string = "딩요"
	nameRussian              string = "Джиттерс"
	nameSpanish              string = "Camelio"
	nameSimplifiedChinese    string = "纪诺"
	nameTraditionalChinese   string = "紀諾"
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
