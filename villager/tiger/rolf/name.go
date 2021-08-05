package rolf

import "golang.org/x/text/language"

const (
	nameAmericanEnglish      string = "Rolf"
	nameCanadianFrench       string = "Ralf"
	nameDutch                string = "Rolf"
	nameFrench               string = "Ralf"
	nameGerman               string = "Boris"
	nameItalian              string = "Rolf"
	nameJapanese             string = "チョモラン"
	nameLatinAmericanSpanish string = "Albino"
	nameKorean               string = "호랭이"
	nameRussian              string = "Рольф"
	nameSpanish              string = "Albino"
	nameSimplifiedChinese    string = "朱穆朗"
	nameTraditionalChinese   string = "朱穆朗"
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
