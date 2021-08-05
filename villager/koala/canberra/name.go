package canberra

import "golang.org/x/text/language"

const (
	nameAmericanEnglish      string = "Canberra"
	nameCanadianFrench       string = "Kolala"
	nameDutch                string = "Canberra"
	nameFrench               string = "Kolala"
	nameGerman               string = "Caroline"
	nameItalian              string = "Canberra"
	nameJapanese             string = "キャンベラ"
	nameLatinAmericanSpanish string = "Canberra"
	nameKorean               string = "캔버라"
	nameRussian              string = "Канберра"
	nameSpanish              string = "Canberra"
	nameSimplifiedChinese    string = "简培菈"
	nameTraditionalChinese   string = "簡培菈"
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
