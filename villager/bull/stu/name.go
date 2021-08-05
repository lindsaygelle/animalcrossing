package stu

import "golang.org/x/text/language"

const (
	nameAmericanEnglish      string = "Stu"
	nameCanadianFrench       string = "Beubeu"
	nameDutch                string = "Stu"
	nameFrench               string = "Beubeu"
	nameGerman               string = "Carlos"
	nameItalian              string = "Carlos"
	nameJapanese             string = "モーリス"
	nameLatinAmericanSpanish string = "Vitorino"
	nameKorean               string = "모리스"
	nameRussian              string = "Стью"
	nameSpanish              string = "Vitorino"
	nameSimplifiedChinese    string = "毛乐时"
	nameTraditionalChinese   string = "毛樂時"
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
